package post

import (
	"errors"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/dungvan2512/soccer-social-network/infrastructure"
	"github.com/dungvan2512/soccer-social-network/model"
	"github.com/dungvan2512/soccer-social-network/shared/base"
	"github.com/dungvan2512/soccer-social-network/shared/utils"
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
)

// Repository interface
type Repository interface {
	// GetAllPostsByUserID return all of post record
	GetAllPostsByUserID(userID uint) ([]model.Post, error)
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
	FindPostByID(postID uint) (*model.Post, error)
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
}

type repository struct {
	base.Repository
	db    *gorm.DB
	redis *redis.Conn
	s3    infrastructure.NewS3RequestFunc
}

func (r *repository) GetAllPostsByUserID(userID uint) ([]model.Post, error) {
	posts := make([]model.Post, 0)
	err := r.db.Model(&model.Post{}).
		Select("id, caption, created_at").Where("user_id = ?", userID).
		Limit(100).
		Order("created_at desc, id desc").
		Scan(&posts).Error
	return posts, utils.ErrorsWrap(err, "can't get post.")
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

func (r *repository) FindPostByID(postID uint) (*model.Post, error) {
	post := &model.Post{}
	err := r.db.Where("id = ?", postID).First(post).Error
	if err == gorm.ErrRecordNotFound {
		return post, err
	}
	return post, utils.ErrorsWrap(err, "can't find")
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

// NewRepository create new instance of Repository
func NewRepository(br *base.Repository, db *gorm.DB, redis *redis.Conn, s3RequestFunc infrastructure.NewS3RequestFunc) Repository {
	return &repository{*br, db, redis, s3RequestFunc}
}
