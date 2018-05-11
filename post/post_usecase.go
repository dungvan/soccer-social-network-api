package post

import (
	"errors"
	"net/http"
	"strings"

	"github.com/dungvan2512/soccer-social-network/infrastructure"
	"github.com/dungvan2512/soccer-social-network/model"
	"github.com/dungvan2512/soccer-social-network/shared/base"
	"github.com/dungvan2512/soccer-social-network/shared/utils"
	"github.com/jinzhu/gorm"
)

// Usecase interface
type Usecase interface {

	//===========================================
	//====================POST===================
	//===========================================

	// Index usecase
	Index(userID uint) (IndexResponse, error)
	// GetByUserID usecase
	GetByUserID(userID uint) (IndexResponse, error)
	// Create a post
	Create(CreateRequest) (postID uint, err error)
	// Show a post
	Show(postID uint) (RespPost, error)
	// Update a post
	Update(request UpdateRequest, ctxUser model.User) (RespPost, error)
	// CountUpStar increase star
	CountUpStar(StarCountRequest) (StarCountResponse, error)
	// CountDownStar  decrease star
	CountDownStar(StarCountRequest) (StarCountResponse, error)
	// Delete a post
	Delete(postID uint, ctxUser model.User) error
	// Add images to s3
	UploadImages(request UploadImagesRequest) (UploadImagesResponse, error)

	//===========================================
	//==================COMMENT==================
	//===========================================

	CommentCreate(r CreateCommentRequest) (uint, error)
	CommentCountUpStar(request StarCountRequest) (StarCountResponse, error)
	CommentCountDownStar(request StarCountRequest) (StarCountResponse, error)
	CommentUpdate(r UpdateCommentRequest, ctxUser model.User) error
	CommentDelete(commentID uint, ctxUser model.User) error
}

type usecase struct {
	base.Usecase
	db         *gorm.DB
	repository Repository
}

func (u *usecase) Index(page uint) (IndexResponse, error) {
	if page < 1 {
		page = 1
	}
	total, posts, err := u.repository.GetAllPost(page)
	if err == gorm.ErrRecordNotFound {
		return IndexResponse{Posts: []RespPost{}}, nil
	}
	if err != nil {
		return IndexResponse{Total: total, Posts: []RespPost{}}, utils.ErrorsWrap(err, "repository.GetAllPost() error")
	}
	response := IndexResponse{
		Total: total,
		Posts: func() []RespPost {
			respUsers := make([]RespPost, 0)
			for _, post := range posts {
				respUsers = append(respUsers, RespPost{
					ID: post.ID,
					User: RespUser{
						ID:        post.UserID,
						UserName:  post.UserName,
						FirstName: post.FirstName,
						LastName:  post.LastName,
					},
					Caption:   post.Caption,
					CreatedAt: post.CreatedAt,
				})
			}
			return respUsers
		}(),
	}
	return response, err
}

func (u *usecase) GetByUserID(userID uint) (IndexResponse, error) {
	indexResp := IndexResponse{}
	postUser, result, err := u.repository.GetAllPostsByUserID(userID)
	if err != nil {
		return indexResp, utils.ErrorsWrap(err, "repository.GetAllPostsByUserID() error.")
	}
	indexResp.Posts = []RespPost{}

	bucketName := infrastructure.GetConfigString("objectstorage.bucketname")
	for _, post := range result {
		err = u.repository.GetRelatedPostImages(&post)
		if err != nil {
			utils.ErrorsWrap(err, "repository.GetRelatedPostImages() error")
		}
		data := RespPost{
			ID:      post.ID,
			User:    postUser,
			Caption: post.Caption,
			ImageURLs: func() []interface{} {
				output := []interface{}{}
				if post.Images != nil && len(post.Images) > 0 {
					for _, image := range post.Images {
						imageurl := utils.GetStorageURL(infrastructure.Storage, infrastructure.Endpoint, infrastructure.Secure, bucketName, utils.GetObjectPath(infrastructure.Storage, s3ImagePath, image.Name), infrastructure.Region)
						output = append(output, imageurl)
					}
				}
				return output
			}(),
			VideoURLs: func() []interface{} {
				output := []interface{}{}
				// for _, videoName := range post.SourceImageFileName {
				// 	videourl := utils.GetStorageURL(infrastructure.Storage, infrastructure.Endpoint, infrastructure.Secure, bucketName, utils.GetObjectPath(infrastructure.Storage, S3ImagePath, videoName), infrastructure.Region)
				// 	output = append(output, videourl)
				// }
				return output
			}(),
			CreatedAt: post.CreatedAt,
		}
		indexResp.Posts = append(indexResp.Posts, data)
	}
	indexResp.Total = uint(len(result))
	return indexResp, err
}

func (u *usecase) Create(r CreateRequest) (uint, error) {
	var isError = false
	tx := u.db.Begin()
	defer func() {
		if isError {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	post := &model.Post{UserID: r.UserID, Caption: r.Caption, StarCount: &model.StarCount{Quantity: 0}}

	if r.PlaceID != "" {

	}
	if r.Images != nil {
		post.Images = []model.Image{}
		for _, imageName := range r.Images {
			post.Images = append(post.Images, model.Image{Name: imageName})
		}
	}
	if r.Videos != nil {

	}
	err := u.repository.CreatePost(post, tx)
	if err != nil {
		isError = true
		return 0, err
	}
	hashtagsID := []uint{}
	if r.Hashtags != nil && len(r.Hashtags) != 0 {
		for key, hashtag := range r.Hashtags {
			r.Hashtags[key] = strings.ToLower(hashtag)
		}
		err := u.repository.CreateHashtags(r.Hashtags, tx)
		if err != nil {
			isError = true
			return 0, utils.ErrorsWrap(err, "repository.CreateHashtags error")
		}
		hashtagsID, err = u.repository.GetHashTagsIDByKeyWords(r.Hashtags, tx)
		if err != nil {
			isError = true
			return 0, utils.ErrorsWrap(err, "repository.GetHashTagsIDByKeyWords error")
		}
	}

	err = u.repository.CreatePostHashtags(post.ID, hashtagsID, tx)
	if err != nil {
		isError = true
		return 0, utils.ErrorsWrap(err, "repository.CreatePostHashtags error")
	}

	return post.ID, nil
}

func (u *usecase) Show(postID uint) (RespPost, error) {
	response := RespPost{}
	user, post, err := u.repository.FindPostByID(postID)
	if err == gorm.ErrRecordNotFound {
		response.TypeOfStatusCode = http.StatusNotFound
		return response, utils.ErrorsNew("the post dose not exist")
	} else if err != nil {
		return response, utils.ErrorsWrap(err, "repository.FindPostByID error")
	}
	err = u.repository.GetRelatedPostImages(post)
	if err != nil {
		utils.ErrorsWrap(err, "repository.GetRelatedPostImages() error")
	}

	bucketName := infrastructure.GetConfigString("objectstorage.bucketname")
	response.ID = post.ID
	response.Caption = post.Caption
	response.User = user
	response.CreatedAt = post.CreatedAt
	response.ImageURLs = func() []interface{} {
		output := []interface{}{}
		if post.Images != nil && len(post.Images) > 0 {
			for _, image := range post.Images {
				imageurl := utils.GetStorageURL(infrastructure.Storage, infrastructure.Endpoint, infrastructure.Secure, bucketName, utils.GetObjectPath(infrastructure.Storage, s3ImagePath, image.Name), infrastructure.Region)
				output = append(output, imageurl)
			}
		}
		return output
	}()
	response.VideoURLs = func() []interface{} {
		output := []interface{}{}
		// if post.Images != nil && len(post.Images) > 0 {
		// 	for _, image := range post.Images {
		// 		imageurl := utils.GetStorageURL(infrastructure.Storage, infrastructure.Endpoint, infrastructure.Secure, bucketName, utils.GetObjectPath(infrastructure.Storage, s3ImagePath, image.Name), infrastructure.Region)
		// 		output = append(output, imageurl)
		// 	}
		// }
		return output
	}()

	return response, nil
}

func (u *usecase) CountUpStar(request StarCountRequest) (StarCountResponse, error) {
	var post *model.Post
	response := StarCountResponse{}
	_, post, err := u.repository.FindPostByID(request.ID)
	if err == gorm.ErrRecordNotFound {
		response.TypeOfStatusCode = http.StatusNotFound
		return response, errors.New("the post does not exist")
	} else if err != nil {
		return response, utils.ErrorsWrap(err, "repository.FinPostByID error")
	}

	tx := u.db.Begin()
	postStar, err := u.repository.FindPostStar(request.UserID, request.ID)
	switch {
	case err == nil && postStar.DeletedAt == nil:
		response.TypeOfStatusCode = http.StatusBadRequest
		tx.Rollback()
		return response, errors.New("Can't tap star many time")
	case err == gorm.ErrRecordNotFound:
		_, err = u.repository.CreatePostStar(request.UserID, request.ID, tx)
		if err != nil {
			tx.Rollback()
			return response, utils.ErrorsWrap(err, "repository.CreatePostStar() error")
		}
		break
	case postStar.DeletedAt != nil:
		_, err = u.repository.RestorePostStar(request.UserID, request.ID, tx)
		if err != nil {
			tx.Rollback()
			return response, utils.ErrorsWrap(err, "repository.RestorePostStar() error")
		}
		break
	default:
		return response, utils.ErrorsWrap(err, "repository.FindPostStar() error")
	}

	_, err = u.repository.FindPostStarCount(*post)
	if err != nil {
		tx.Rollback()
		return response, utils.ErrorsWrap(err, "repository.FindPostStarCount() error")
	}
	starCount, err := u.repository.UpdatePostStarCount(upUnit, request.ID, tx)
	if err != nil {
		tx.Rollback()
		return response, utils.ErrorsWrap(err, "repository.UpdatePostStarCount() error")
	}
	response.StarCount = starCount
	tx.Commit()
	return response, err
}

func (u *usecase) CountDownStar(request StarCountRequest) (StarCountResponse, error) {
	var post *model.Post
	response := StarCountResponse{}
	_, post, err := u.repository.FindPostByID(request.ID)
	if err == gorm.ErrRecordNotFound {
		response.TypeOfStatusCode = http.StatusNotFound
		return response, errors.New("The outfit does not exist")
	} else if err != nil {
		return response, utils.ErrorsWrap(err, "repository.FindPostByID error")
	}
	postStarCount, err := u.repository.FindPostStarCount(*post)
	if err == gorm.ErrRecordNotFound || (err == nil && postStarCount.Quantity == defaultStarCount) {
		response.TypeOfStatusCode = http.StatusBadRequest
		return response, errors.New("The outfit has no stars")
	} else if err != nil {
		return response, utils.ErrorsWrap(err, "repository.FindPostStarCount error")
	}

	postStar, err := u.repository.FindPostStar(request.UserID, request.ID)
	if err == gorm.ErrRecordNotFound || postStar.DeletedAt != nil {
		response.TypeOfStatusCode = http.StatusBadRequest
		return response, errors.New("User has not tapped or untapped the star before")
	} else if err != nil {
		return response, utils.ErrorsWrap(err, "repository.FindPostStar error")
	}

	tx := u.db.Begin()
	starCount, err := u.repository.UpdatePostStarCount(downUnit, request.ID, tx)
	if err != nil {
		tx.Rollback()
		return response, utils.ErrorsWrap(err, "repository.UpdatePostStarCount error")
	}
	_, err = u.repository.DeletePostStar(request.UserID, request.ID, tx)
	if err != nil {
		tx.Rollback()
		return response, utils.ErrorsWrap(err, "repository.DeletePostStar error")
	}

	tx.Commit()
	response.StarCount = starCount
	return response, err
}

func (u *usecase) UploadImages(request UploadImagesRequest) (UploadImagesResponse, error) {
	response := UploadImagesResponse{[]string{}}
	for index, image := range request.Images {
		err := u.repository.AddImageToS3(image, s3ImagePath)
		if err != nil {
			utils.ErrorsWrap(err, "can't upload file "+string(index)+" to S3")
			continue
		}
		response.ImageNames = append(response.ImageNames, image.Name)
	}
	return response, nil
}

func (u *usecase) Update(r UpdateRequest, ctxUser model.User) (RespPost, error) {
	var err error
	user, post, err := u.repository.FindPostByID(r.ID)
	if err != nil {
		return RespPost{}, utils.ErrorsWrap(err, "repository.FindPostByID() error")
	}
	if ctxUser.Role != "s_admin" {
		if post.UserID != ctxUser.ID {
			return RespPost{}, utils.ErrorsNew("Forbbiden to update the post")
		}
	}
	post.Caption = r.Caption
	err = u.repository.UpdatePost(post)
	if err != nil {
		return RespPost{}, utils.ErrorsWrap(err, "repository.UpdatePost() error")
	}
	return RespPost{
		ID:        post.ID,
		User:      user,
		Caption:   post.Caption,
		CreatedAt: post.CreatedAt,
	}, nil
}

func (u *usecase) Delete(postID uint, ctxUser model.User) error {
	var err error
	if ctxUser.Role != "s_admin" {
		_, post, err := u.repository.FindPostByID(postID)
		if err != nil {
			return utils.ErrorsWrap(err, "repository.FindPostByID() error")
		}
		if post.UserID != ctxUser.ID {
			return utils.ErrorsNew("Forbbiden to delete the post")
		}
	}

	tx := u.db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()

	err = u.repository.DeletePost(postID, tx)
	if err != nil {
		return utils.ErrorsWrap(err, "repository.DeletePost() error")
	}
	err = u.repository.DeleteRelatedPostImages(postID, tx)
	if err != nil {
		return utils.ErrorsWrap(err, "repository.DeleteRelatePostImages() error")
	}
	return nil
}

//===========================================
//==================COMMENT==================
//===========================================

func (u *usecase) CommentCreate(r CreateCommentRequest) (uint, error) {
	comment := &model.Comment{UserID: r.UserID, PostID: r.PostID, Content: r.Content, StarCount: &model.StarCount{Quantity: 0}}
	err := u.repository.CreateComment(comment)
	if err != nil {
		return 0, err
	}
	return comment.ID, nil
}

func (u *usecase) CommentCountUpStar(request StarCountRequest) (StarCountResponse, error) {
	var comment *model.Comment
	response := StarCountResponse{}
	comment, err := u.repository.FindCommentByID(request.ID)
	if err == gorm.ErrRecordNotFound {
		response.TypeOfStatusCode = http.StatusNotFound
		return response, errors.New("the comment does not exist")
	} else if err != nil {
		return response, utils.ErrorsWrap(err, "repository.FindCommentByID error")
	}

	tx := u.db.Begin()
	commentStar, err := u.repository.FindCommentStar(request.UserID, request.ID)
	switch {
	case err == nil && commentStar.DeletedAt == nil:
		response.TypeOfStatusCode = http.StatusBadRequest
		tx.Rollback()
		return response, errors.New("Can't tap star many time")
	case err == gorm.ErrRecordNotFound:
		_, err = u.repository.CreateCommentStar(request.UserID, request.ID, tx)
		if err != nil {
			tx.Rollback()
			return response, utils.ErrorsWrap(err, "repository.CreateCommentStar() error")
		}
		break
	case commentStar.DeletedAt != nil:
		_, err = u.repository.RestoreCommentStar(request.UserID, request.ID, tx)
		if err != nil {
			tx.Rollback()
			return response, utils.ErrorsWrap(err, "repository.RestoreCommentStar() error")
		}
		break
	default:
		return response, utils.ErrorsWrap(err, "repository.FindCommentStar() error")
	}

	_, err = u.repository.FindCommentStarCount(*comment)
	if err != nil {
		tx.Rollback()
		return response, utils.ErrorsWrap(err, "repository.FindCommentStarCount() error")
	}
	starCount, err := u.repository.UpdateCommentStarCount(upUnit, request.ID, tx)
	if err != nil {
		tx.Rollback()
		return response, utils.ErrorsWrap(err, "repository.UpdateCommentStarCount() error")
	}
	response.StarCount = starCount
	tx.Commit()
	return response, err
}

func (u *usecase) CommentCountDownStar(request StarCountRequest) (StarCountResponse, error) {
	var comment *model.Comment
	response := StarCountResponse{}
	comment, err := u.repository.FindCommentByID(request.ID)
	if err == gorm.ErrRecordNotFound {
		response.TypeOfStatusCode = http.StatusNotFound
		return response, errors.New("The outfit does not exist")
	} else if err != nil {
		return response, utils.ErrorsWrap(err, "repository.FindCommentByID error")
	}
	commentStarCount, err := u.repository.FindCommentStarCount(*comment)
	if err == gorm.ErrRecordNotFound || (err == nil && commentStarCount.Quantity == defaultStarCount) {
		response.TypeOfStatusCode = http.StatusBadRequest
		return response, errors.New("The outfit has no stars")
	} else if err != nil {
		return response, utils.ErrorsWrap(err, "repository.FindCommentStarCount error")
	}

	commentStar, err := u.repository.FindCommentStar(request.UserID, request.ID)
	if err == gorm.ErrRecordNotFound || commentStar.DeletedAt != nil {
		response.TypeOfStatusCode = http.StatusBadRequest
		return response, errors.New("User has not tapped or untapped the star before")
	} else if err != nil {
		return response, utils.ErrorsWrap(err, "repository.FindCommentStar error")
	}

	tx := u.db.Begin()
	starCount, err := u.repository.UpdateCommentStarCount(downUnit, request.ID, tx)
	if err != nil {
		tx.Rollback()
		return response, utils.ErrorsWrap(err, "repository.UpdateCommentStarCount error")
	}
	_, err = u.repository.DeleteCommentStar(request.UserID, request.ID, tx)
	if err != nil {
		tx.Rollback()
		return response, utils.ErrorsWrap(err, "repository.DeleteCommentStar error")
	}

	tx.Commit()
	response.StarCount = starCount
	return response, err
}

func (u *usecase) CommentUpdate(r UpdateCommentRequest, ctxUser model.User) error {
	var err error
	comment, err := u.repository.FindCommentByID(r.ID)
	if err != nil {
		return utils.ErrorsWrap(err, "repository.FindCommentByID() error")
	}
	if ctxUser.Role != "s_admin" {
		if comment.UserID != ctxUser.ID {
			return utils.ErrorsNew("Forbbiden to update the comment")
		}
	}
	comment.Content = r.Content
	err = u.repository.UpdateComment(comment)
	if err != nil {
		return utils.ErrorsWrap(err, "repository.UpdateComment() error")
	}
	return nil
}

func (u *usecase) CommentDelete(commentID uint, ctxUser model.User) error {
	var err error
	if ctxUser.Role != "s_admin" {
		comment, err := u.repository.FindCommentByID(commentID)
		if err != nil {
			return utils.ErrorsWrap(err, "repository.FindCommentByID() error")
		}
		if comment.UserID != ctxUser.ID {
			return utils.ErrorsNew("Forbbiden to delete the comment")
		}
	}

	err = u.repository.DeleteComment(commentID)
	if err != nil {
		return utils.ErrorsWrap(err, "repository.DeleteComment() error")
	}
	return nil
}

// NewUsecase creare new instance of Usecase
func NewUsecase(bu *base.Usecase, db *gorm.DB, r Repository) Usecase {
	return &usecase{*bu, db, r}
}
