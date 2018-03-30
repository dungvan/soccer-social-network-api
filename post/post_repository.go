package post

import (
	"strings"

	"github.com/dungvan2512/socker-social-network/model"
	"github.com/dungvan2512/socker-social-network/shared/base"
	"github.com/dungvan2512/socker-social-network/shared/utils"
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
)

// Repository interface
type Repository interface {
	// CreatePost registers record to table post
	CreatePost(post model.Post, transaction *gorm.DB) (*model.Post, error)
	// CreateHashtags is insert hashtag list into hashtag table if it does not exist.
	CreateHashtags(hashtags []string, transaction *gorm.DB) error
	// GetHashTagsIDByKeyWords get array id hashtag from hashtags request.
	GetHashTagsIDByKeyWords(hashtags []string, transaction *gorm.DB) ([]uint, error)
	// CreatePostHashtags create outfit hashtags
	CreatePostHashtags(postID uint, hashtagsID []uint, transaction *gorm.DB) error
}

type repository struct {
	base.Repository
	db    *gorm.DB
	redis *redis.Conn
}

func (r *repository) CreatePost(p model.Post, tx *gorm.DB) (*model.Post, error) {
	result := tx.Create(&p)
	return &p, utils.ErrorsWrap(result.Error, "can't create post")
}

func (r *repository) CreateHashtags(hashtags []string, tx *gorm.DB) error {
	if len(hashtags) == 0 {
		return nil
	}
	sql := "INSERT INTO hashtag (key) VALUES "
	vals := []interface{}{}
	params := []string{}
	for _, hashtag := range hashtags {
		params = append(params, "(?)")
		vals = append(vals, hashtag)
	}
	sql += strings.Join(params, ",")
	sql += " ON CONFLICT (key) DO NOTHING"
	err := tx.Exec(sql, vals...).Error
	return utils.ErrorsWrap(err, "can't insert hashtag")
}

func (r *repository) GetHashTagsIDByKeyWords(hashtags []string, tx *gorm.DB) ([]uint, error) {
	sqlSelect := "SELECT id FROM hashtag WHERE key IN (?)"
	var hashtagsID []uint
	var hashtagID uint
	rows, err := tx.Raw(sqlSelect, hashtags).Rows()
	if err != nil {
		return hashtagsID, utils.ErrorsWrap(err, "can't get hashtag")
	}
	defer func() {
		_ = rows.Close()
	}()
	for rows.Next() {
		err = rows.Scan(&hashtagID)
		if err != nil {
			return nil, utils.ErrorsWrap(err, "can't get hashtag")
		}
		hashtagsID = append(hashtagsID, hashtagID)
	}
	return hashtagsID, nil
}

func (r *repository) CreatePostHashtags(oufitID uint, hashtagsID []uint, tx *gorm.DB) error {
	if len(hashtagsID) == 0 {
		return nil
	}
	sqlStr := "INSERT INTO post_hashtags(post_id, hashtag_id) VALUES "
	vals := []interface{}{}
	params := []string{}
	for _, hashtagID := range hashtagsID {
		params = append(params, "(?, ?)")
		vals = append(vals, oufitID, hashtagID)
	}
	sqlStr += strings.Join(params, ",")
	err := tx.Exec(sqlStr, vals...).Error
	return utils.ErrorsWrap(err, "can't create data.")
}

// NewRepository create new instance of Repository
func NewRepository(br *base.Repository, db *gorm.DB, redis *redis.Conn) Repository {
	return &repository{*br, db, redis}
}
