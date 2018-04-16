package post

import (
	"errors"
	"net/http"
	"strings"

	"github.com/dungvan2512/socker-social-network/model"
	"github.com/dungvan2512/socker-social-network/shared/base"
	"github.com/dungvan2512/socker-social-network/shared/utils"
	"github.com/jinzhu/gorm"
)

// Usecase interface
type Usecase interface {
	// Index usecase
	Index(userID uint) (IndexResponse, error)
	// Create a post
	Create(CreateRequest) (postID uint, err error)
	// Update a post
	Update(UpdateRequest) (uint, error)
	// CountUpStar increase star
	CountUpStar(StarCountRequest) (StarCountResponse, error)
	// CountDownStar  decrease star
	CountDownStar(StarCountRequest) (StarCountResponse, error)
}

type usecase struct {
	base.Usecase
	db         *gorm.DB
	repository Repository
}

func (u *usecase) Index(userID uint) (IndexResponse, error) {
	indexResp := IndexResponse{}
	result, err := u.repository.GetAllPostsByUserID(userID)
	if err != nil {
		return indexResp, utils.ErrorsWrap(err, "repository.GetAllPostsByUserID() error.")
	}
	indexResp.ResultCount = len(result)
	indexResp.Posts = []RespPost{}

	// bucketName := infrastructure.GetConfigString("objectstorage.bucketname")
	for _, post := range result {
		data := RespPost{
			ID:      post.ID,
			UserID:  userID,
			Caption: post.Caption,
			ImageURLs: func() []interface{} {
				output := []interface{}{}
				// if post.SourceImageFileName != nil && len(post.SourceImageFileName) > 0 {
				// 	for _, imageName := range post.SourceImageFileName {
				// 		imageurl := utils.GetStorageURL(infrastructure.Storage, infrastructure.Endpoint, infrastructure.Secure, bucketName, utils.GetObjectPath(infrastructure.Storage, S3ImagePath, imageName), infrastructure.Region)
				// 		output = append(output, imageurl)
				// 	}
				// }
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
	post := model.Post{UserID: r.User.ID, Caption: r.Caption, StarCount: &model.StarCount{Quantity: 0}}

	if r.PlaceID != "" {

	}
	if r.Images != nil {

	}
	if r.Videos != nil {

	}
	postResponse, err := u.repository.CreatePost(post, tx)
	if err != nil {
		isError = true
		return 0, err
	}
	hashtagIDs := []uint{}
	if r.Hashtags != nil && len(r.Hashtags) != 0 {
		for key, hashtag := range r.Hashtags {
			r.Hashtags[key] = strings.ToLower(hashtag)
		}
		err := u.repository.CreateHashtags(r.Hashtags, tx)
		if err != nil {
			isError = true
			return 0, utils.ErrorsWrap(err, "repository.CreateHashtags error")
		}
		hashtagIDs, err = u.repository.GetHashTagsIDByKeyWords(r.Hashtags, tx)
		if err != nil {
			isError = true
			return 0, utils.ErrorsWrap(err, "repository.GetHashTagsIDByKeyWords error")
		}
	}

	err = u.repository.CreatePostHashtags(postResponse.ID, hashtagIDs, tx)
	if err != nil {
		isError = true
		return 0, utils.ErrorsWrap(err, "repository.CreatePostHashtags error")
	}

	return postResponse.ID, nil
}

func (u *usecase) Update(r UpdateRequest) (uint, error) {
	return 0, nil
}

func (u *usecase) CountUpStar(request StarCountRequest) (StarCountResponse, error) {
	var post *model.Post
	response := StarCountResponse{}
	post, err := u.repository.FindPostByID(request.PostID)
	if err == gorm.ErrRecordNotFound {
		response.TypeOfStatusCode = http.StatusNotFound
		return response, errors.New("The post does not exist")
	} else if err != nil {
		return response, utils.ErrorsWrap(err, "repository.FinPostByID error")
	}

	tx := u.db.Begin()
	postStar, err := u.repository.FindPostStar(request.UserID, request.PostID)
	switch {
	case err == nil && postStar.DeletedAt == nil:
		response.TypeOfStatusCode = http.StatusBadRequest
		tx.Rollback()
		return response, errors.New("Can't tap star many time")
	case err == gorm.ErrRecordNotFound:
		_, err = u.repository.CreatePostStar(request.UserID, request.PostID, tx)
		if err != nil {
			tx.Rollback()
			return response, utils.ErrorsWrap(err, "repository.CreatePostStar() error")
		}
		break
	case postStar.DeletedAt != nil:
		_, err = u.repository.RestorePostStar(request.UserID, request.PostID, tx)
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
	starCount, err := u.repository.UpdatePostStarCount(upUnit, request.PostID, tx)
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
	post, err := u.repository.FindPostByID(request.PostID)
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

	postStar, err := u.repository.FindPostStar(request.UserID, request.PostID)
	if err == gorm.ErrRecordNotFound || postStar.DeletedAt != nil {
		response.TypeOfStatusCode = http.StatusBadRequest
		return response, errors.New("User has not tapped or untapped the star before")
	} else if err != nil {
		return response, utils.ErrorsWrap(err, "repository.FindPostStar error")
	}

	tx := u.db.Begin()
	starCount, err := u.repository.UpdatePostStarCount(downUnit, request.PostID, tx)
	if err != nil {
		tx.Rollback()
		return response, utils.ErrorsWrap(err, "repository.UpdatePostStarCount error")
	}
	_, err = u.repository.DeletePostStar(request.UserID, request.PostID, tx)
	if err != nil {
		tx.Rollback()
		return response, utils.ErrorsWrap(err, "repository.DeletePostStar error")
	}

	tx.Commit()
	response.StarCount = starCount
	return response, err
}

// NewUsecase creare new instance of Usecase
func NewUsecase(bu *base.Usecase, db *gorm.DB, r Repository) Usecase {
	return &usecase{*bu, db, r}
}
