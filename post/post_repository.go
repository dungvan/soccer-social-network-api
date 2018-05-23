package post

import (
	"errors"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/dungvan2512/soccer-social-network-api/infrastructure"
	"github.com/dungvan2512/soccer-social-network-api/model"
	"github.com/dungvan2512/soccer-social-network-api/shared/base"
	"github.com/dungvan2512/soccer-social-network-api/shared/utils"
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
)

// Repository interface
type Repository interface {

	//=======================================
	//==================POST=================
	//=======================================

	// GetAllPost return all post with pagination
	GetAllPost(userID, page uint) (total uint, posts []Post, err error)
	// GetAllPostsByUserID return all of post record
	GetAllPostsByUserID(userIDCreate, userIDCall, page uint) (uint, []Post, error)
	// CreatePost registers record to table post
	CreatePost(post *model.Post, transaction *gorm.DB) error
	// CreateHashtags is insert hashtag list into hashtag table if it does not exist.
	CreateHashtags(hashtags []string, transaction *gorm.DB) error
	// GetHashTagsIDByKeyWords get array id hashtag from hashtags request.
	GetHashTagsIDByKeyWords(hashtags []string, transaction *gorm.DB) ([]uint, error)
	// CreatePostHashtags create outfit hashtags
	CreatePostHashtags(postID uint, hashtagsID []uint, transaction *gorm.DB) error
	// FindPostStarHistory find PostStar contain record has deleted_at != null.
	FindPostStar(userID, postID uint) (*model.PostStar, error)
	// FindPostByID find a post by id
	FindPostByID(postID uint) (RespUser, *model.Post, error)
	// CreatePostStar create new PostStar
	CreatePostStar(userID, postID uint, transasction *gorm.DB) (rowsAffected int64, err error)
	// RestorePostStar update deleted_at field = nill
	RestorePostStar(userID uint, postID uint, transaction *gorm.DB) (rowsAffected int64, err error)
	// FindPostStarCount find StarCount.
	FindPostStarCount(post model.Post) (*model.StarCount, error)
	// UpdatePostStarCount update StarCount field in the star_counts table
	UpdatePostStarCount(unit int, postID uint, transaction *gorm.DB) (starCount uint, err error)
	// soft delete PostStar
	DeletePostStar(userID uint, postID uint, transaction *gorm.DB) (rowsAffected int64, err error)
	// AddImageToS3 upload image to server S3
	AddImageToS3(image Image, s3Path string) error
	// GetRelatedPostImages get images of a post
	GetRelatedPostImages(post *model.Post) error
	// UpdatePost
	UpdatePost(post *model.Post) error
	// DeletePost
	DeletePost(postID uint, transaction *gorm.DB) error
	// DeleteRelatePostImages
	DeleteRelatedPostImages(postID uint, transaction *gorm.DB) error

	//=======================================
	//================COMMENT================
	//=======================================

	// GetAllCommentsByPostID return all of comment record
	GetAllCommentsByPostID(postID uint) ([]Comment, error)
	// CreateComment registers record to table comment
	CreateComment(comment *model.Comment) (RespUser, error)
	// FindCommentStarHistory find CommentStar contain record has deleted_at != null.
	FindCommentStar(userID, commentID uint) (*model.CommentStar, error)
	// FindCommentByID find a comment by id
	FindCommentByID(commentID uint) (*model.Comment, error)
	// CreateCommentStar create new CommentStar
	CreateCommentStar(userID, commentID uint, transasction *gorm.DB) (rowsAffected int64, err error)
	// RestoreCommentStar update deleted_at field = nill
	RestoreCommentStar(userID uint, commentID uint, transaction *gorm.DB) (rowsAffected int64, err error)
	// FindCommentStarCount find StarCount.
	FindCommentStarCount(comment model.Comment) (*model.StarCount, error)
	// UpdateCommentStarCount update StarCount field in the star_counts table
	UpdateCommentStarCount(unit int, commentID uint, transaction *gorm.DB) (starCount uint, err error)
	// soft delete CommentStar
	DeleteCommentStar(userID uint, commentID uint, transaction *gorm.DB) (rowsAffected int64, err error)
	// UpdateComment
	UpdateComment(comment *model.Comment) error
	// DeleteComment
	DeleteComment(commentID uint) error
}

type repository struct {
	base.Repository
	db    *gorm.DB
	redis *redis.Conn
	s3    infrastructure.NewS3RequestFunc
}

//=======================================
//==================POST=================
//=======================================

func (r *repository) GetAllPost(userID, page uint) (uint, []Post, error) {
	var total uint
	var err error
	posts := make([]Post, 0)
	result := r.db.Model(&model.Post{}).
		Select("users.user_name, users.first_name, users.last_name, posts.id, posts.user_id, posts.caption, posts.type, posts.created_at, COALESCE (star_counts.quantity,0) AS star_count, post_stars.id, CASE WHEN post_stars.id IS NOT NULL THEN true ELSE false END AS star_flag").
		Joins(`JOIN users ON (posts.user_id = users.id AND users.deleted_at IS NULL)`).
		Joins(`JOIN star_counts ON (star_counts.owner_type = 'posts' AND star_counts.owner_id = posts.id AND star_counts.deleted_at IS NULL)`).
		Joins(`LEFT JOIN post_stars ON (posts.id = post_stars.post_id AND post_stars.user_id = ? AND post_stars.deleted_at IS NULL)`, userID)
	result.Count(&total)
	if total <= pagingLimit*(page-1) {
		return total, posts, gorm.ErrRecordNotFound
	}
	err = result.
		Offset(pagingLimit * (page - 1)).
		Limit(pagingLimit).Order("posts.created_at desc, posts.id desc").
		Scan(&posts).Error
	return total, posts, utils.ErrorsWrap(err, "can't get all posts")
}

func (r *repository) GetAllPostsByUserID(userIDCreate, userIDCall, page uint) (uint, []Post, error) {
	var total uint
	var err error
	posts := make([]Post, 0)
	result := r.db.Model(&model.Post{}).
		Select("users.user_name, users.first_name, users.last_name, posts.id, posts.user_id, posts.caption, posts.type, posts.created_at, COALESCE (star_counts.quantity,0) AS star_count, post_stars.id, CASE WHEN post_stars.id IS NOT NULL THEN true ELSE false END AS star_flag").
		Joins(`JOIN users ON (posts.user_id = users.id AND users.id = ? AND users.deleted_at IS NULL)`, userIDCreate).
		Joins(`JOIN star_counts ON (star_counts.owner_type = 'posts' AND star_counts.owner_id = posts.id AND star_counts.deleted_at IS NULL)`).
		Joins(`LEFT JOIN post_stars ON (posts.id = post_stars.post_id AND post_stars.user_id = ? AND post_stars.deleted_at IS NULL)`, userIDCall)
	result.Count(&total)
	if total <= pagingLimit*(page-1) {
		return total, posts, gorm.ErrRecordNotFound
	}
	err = result.
		Offset(pagingLimit * (page - 1)).
		Limit(pagingLimit).Order("posts.created_at desc, posts.id desc").
		Scan(&posts).Error
	return total, posts, utils.ErrorsWrap(err, "can't get all posts")
}

func (r *repository) CreatePost(p *model.Post, tx *gorm.DB) error {
	result := tx.Create(p)
	return utils.ErrorsWrap(result.Error, "can't create post")
}

func (r *repository) CreateHashtags(hashtags []string, tx *gorm.DB) error {
	if len(hashtags) == 0 {
		return nil
	}
	sql := "INSERT INTO hashtags (key_word, created_at, updated_at) VALUES "
	vals := []interface{}{}
	params := []string{}
	for _, hashtag := range hashtags {
		params = append(params, "(?, ?, ?)")
		vals = append(vals, hashtag, time.Now(), time.Now())
	}
	sql += strings.Join(params, ",")
	sql += " ON CONFLICT (key_word) DO NOTHING"
	err := tx.Exec(sql, vals...).Error
	return utils.ErrorsWrap(err, "can't insert hashtag")
}

func (r *repository) GetHashTagsIDByKeyWords(hashtags []string, tx *gorm.DB) ([]uint, error) {
	sqlSelect := "SELECT id FROM hashtags WHERE key_word IN (?)"
	var hashtagsID []uint
	var hashtagID uint
	rows, err := tx.Raw(sqlSelect, hashtags).Rows()
	if err != nil {
		return hashtagsID, utils.ErrorsWrap(err, "can't get hashtags")
	}
	defer func() {
		_ = rows.Close()
	}()
	for rows.Next() {
		err = rows.Scan(&hashtagID)
		if err != nil {
			return nil, utils.ErrorsWrap(err, "can't get hashtags")
		}
		hashtagsID = append(hashtagsID, hashtagID)
	}
	return hashtagsID, nil
}

func (r *repository) CreatePostHashtags(postID uint, hashtagsID []uint, tx *gorm.DB) error {
	if len(hashtagsID) == 0 {
		return nil
	}
	sqlStr := "INSERT INTO post_hashtags(post_id, hashtag_id, created_at, updated_at) VALUES "
	vals := []interface{}{}
	params := []string{}
	for _, hashtagID := range hashtagsID {
		params = append(params, "(?, ?, ?, ?)")
		vals = append(vals, postID, hashtagID, time.Now(), time.Now())
	}
	sqlStr += strings.Join(params, ",")
	err := tx.Exec(sqlStr, vals...).Error
	return utils.ErrorsWrap(err, "can't create data.")
}

func (r *repository) FindPostStar(userID, postID uint) (*model.PostStar, error) {
	postStar := &model.PostStar{}
	err := r.db.Unscoped().Where("post_id = ? AND user_id = ?", postID, userID).First(postStar).Error
	if err == gorm.ErrRecordNotFound {
		return postStar, err
	}
	return postStar, utils.ErrorsWrap(err, "can't find")
}

func (r *repository) FindPostByID(postID uint) (RespUser, *model.Post, error) {
	post := &model.Post{}
	user := RespUser{}
	err := r.db.Where("id = ?", postID).First(post).Error
	if err == gorm.ErrRecordNotFound {
		return user, post, err
	}
	if err != nil {
		return user, post, utils.ErrorsWrap(err, "can't find")
	}
	err = r.db.Model(&model.User{}).Where("id = ?", post.UserID).Scan(&user).Error
	return user, post, utils.ErrorsWrap(err, "can't get related post-user.")
}

func (r *repository) CreatePostStar(userID, postID uint, tx *gorm.DB) (int64, error) {
	postStar := &model.PostStar{PostID: postID, UserID: userID}
	result := tx.Create(postStar)
	return result.RowsAffected, utils.ErrorsWrap(result.Error, "can't create")
}

func (r *repository) RestorePostStar(userID uint, postID uint, tx *gorm.DB) (int64, error) {
	postStar, err := r.FindPostStar(userID, postID)
	if err != nil {
		return 0, utils.ErrorsWrap(err, "can't find")
	}
	postStar.DeletedAt = nil
	result := tx.Unscoped().Save(postStar)
	return result.RowsAffected, utils.ErrorsWrap(result.Error, "can't restore")
}

func (r *repository) FindPostStarCount(post model.Post) (*model.StarCount, error) {
	starCount := &model.StarCount{}
	err := r.db.Model(&post).Related(starCount, "StarCount").Error
	if err == gorm.ErrRecordNotFound {
		return starCount, err
	}
	return starCount, utils.ErrorsWrap(err, "can't find")
}

func (r *repository) UpdatePostStarCount(unit int, postID uint, tx *gorm.DB) (uint, error) {
	starCount := &model.StarCount{}
	err := tx.Where("owner_id = ? and owner_type = ?", postID, "posts").First(starCount).Error
	if err != nil {
		return 0, utils.ErrorsWrap(err, "can't find")
	}
	starCount.Quantity += uint(unit)
	err = tx.Save(starCount).Error
	return starCount.Quantity, utils.ErrorsWrap(err, "can't update")
}

func (r *repository) DeletePostStar(userID uint, postID uint, tx *gorm.DB) (int64, error) {
	result := tx.Where("user_id = ? AND post_id =?", userID, postID).Delete(&model.PostStar{})
	return result.RowsAffected, utils.ErrorsWrap(result.Error, "can't delete")
}

func (r *repository) AddImageToS3(image Image, s3Path string) error {
	bucketName := infrastructure.GetConfigString("objectstorage.bucketname")
	if image.Body == nil {
		return utils.ErrorsWrap(errors.New("no file detected"), "can't find uploadfile")
	}
	objectName := utils.GetObjectPath(infrastructure.Storage, s3Path, image.Name)
	params, err := r.s3().SetParam(image.Body, bucketName, objectName, image.MimeType, s3.BucketCannedACLPublicReadWrite).UploadToS3()
	r.Logger.Debug(params.String())
	if err != nil {
		return utils.ErrorsWrap(err, "can't put S3")
	}
	return nil
}

func (r *repository) GetRelatedPostImages(post *model.Post) error {
	post.Images = make([]model.Image, 0)
	result := r.db.Model(post).Related(&post.Images)
	return utils.ErrorsWrap(result.Error, "can't get posts-images relation")
}

func (r *repository) UpdatePost(post *model.Post) error {
	return utils.ErrorsWrap(r.db.Model(&model.Post{}).Update(post).Error, "can't update post")
}

func (r *repository) DeletePost(postID uint, transaction *gorm.DB) error {
	return utils.ErrorsWrap(r.db.Where("id = ?", postID).Delete(&model.Post{}).Error, "can't Delete post")
}

func (r *repository) DeleteRelatedPostImages(postID uint, transaction *gorm.DB) error {
	return utils.ErrorsWrap(r.db.Where("post_id = ?", postID).Delete(&model.Image{}).Error, "can't Delete related post-images")
}

//=======================================
//================COMMENT================
//=======================================

func (r *repository) GetAllCommentsByPostID(postID uint) ([]Comment, error) {
	comments := make([]Comment, 0)
	err := r.db.Model(&model.Comment{}).
		Select("comments.id, comments.user_id, users.user_name, users.first_name, users.last_name, comments.content, comments.created_at, COALESCE (star_counts.quantity,0) AS star_count, comment_stars.id, CASE WHEN comment_stars.id IS NOT NULL THEN true ELSE false END AS star_flag").
		Joins(`JOIN users ON (comments.user_id = users.id AND users.deleted_at IS NULL AND comments.post_id = ?)`, postID).
		Joins(`JOIN star_counts ON (star_counts.owner_type = 'comments' AND star_counts.owner_id = comments.id AND star_counts.deleted_at IS NULL)`).
		Joins(`LEFT JOIN comment_stars ON (comments.id = comment_stars.comment_id AND comment_stars.deleted_at IS NULL)`).
		Limit(100).
		Order("comments.created_at asc, comments.id asc").
		Scan(&comments).Error
	return comments, utils.ErrorsWrap(err, "can't get comment.")
}

func (r *repository) CreateComment(c *model.Comment) (RespUser, error) {
	err := r.db.Create(c).Error
	if err != nil {
		return RespUser{}, utils.ErrorsWrap(err, "can't create comment")
	}
	user := RespUser{}
	err = r.db.Model(&model.User{}).Where("id = ?", c.UserID).Scan(&user).Error
	if err != nil {
		return RespUser{}, utils.ErrorsWrap(err, "can't create comment")
	}
	return user, err
}

func (r *repository) FindCommentStar(userID, commentID uint) (*model.CommentStar, error) {
	commentStar := &model.CommentStar{}
	err := r.db.Unscoped().Where("comment_id = ? AND user_id = ?", commentID, userID).First(commentStar).Error
	if err == gorm.ErrRecordNotFound {
		return commentStar, err
	}
	return commentStar, utils.ErrorsWrap(err, "can't find")
}

func (r *repository) FindCommentByID(commentID uint) (*model.Comment, error) {
	comment := &model.Comment{}
	err := r.db.Where("id = ?", commentID).First(comment).Error
	if err == gorm.ErrRecordNotFound {
		return comment, err
	}
	return comment, utils.ErrorsWrap(err, "can't find")
}

func (r *repository) CreateCommentStar(userID, commentID uint, tx *gorm.DB) (int64, error) {
	commentStar := &model.CommentStar{CommentID: commentID, UserID: userID}
	result := tx.Create(commentStar)
	return result.RowsAffected, utils.ErrorsWrap(result.Error, "can't create")
}

func (r *repository) RestoreCommentStar(userID uint, commentID uint, tx *gorm.DB) (int64, error) {
	commentStar, err := r.FindCommentStar(userID, commentID)
	if err != nil {
		return 0, utils.ErrorsWrap(err, "can't find")
	}
	commentStar.DeletedAt = nil
	result := tx.Unscoped().Save(commentStar)
	return result.RowsAffected, utils.ErrorsWrap(result.Error, "can't restore")
}

func (r *repository) FindCommentStarCount(comment model.Comment) (*model.StarCount, error) {
	starCount := &model.StarCount{}
	err := r.db.Model(&comment).Related(starCount, "StarCount").Error
	if err == gorm.ErrRecordNotFound {
		return starCount, err
	}
	return starCount, utils.ErrorsWrap(err, "can't find")
}

func (r *repository) UpdateCommentStarCount(unit int, commentID uint, tx *gorm.DB) (uint, error) {
	starCount := &model.StarCount{}
	err := tx.Where("owner_id = ? and owner_type = ?", commentID, "comments").First(starCount).Error
	if err != nil {
		return 0, utils.ErrorsWrap(err, "can't find")
	}
	starCount.Quantity += uint(unit)
	err = tx.Save(starCount).Error
	return starCount.Quantity, utils.ErrorsWrap(err, "can't update")
}

func (r *repository) DeleteCommentStar(userID uint, commentID uint, tx *gorm.DB) (int64, error) {
	result := tx.Where("user_id = ? AND comment_id =?", userID, commentID).Delete(&model.CommentStar{})
	return result.RowsAffected, utils.ErrorsWrap(result.Error, "can't delete")
}

func (r *repository) UpdateComment(comment *model.Comment) error {
	return utils.ErrorsWrap(r.db.Model(&model.Comment{}).Update(comment).Error, "can't update comment")
}

func (r *repository) DeleteComment(commentID uint) error {
	return utils.ErrorsWrap(r.db.Where("id = ?", commentID).Delete(&model.Comment{}).Error, "can't Delete comment")
}

// NewRepository create new instance of Repository
func NewRepository(br *base.Repository, db *gorm.DB, redis *redis.Conn, s3RequestFunc infrastructure.NewS3RequestFunc) Repository {
	return &repository{*br, db, redis, s3RequestFunc}
}
