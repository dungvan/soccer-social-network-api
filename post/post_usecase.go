package post

import (
	"strings"

	"github.com/dungvan2512/socker-social-network/model"
	"github.com/dungvan2512/socker-social-network/shared/base"
	"github.com/dungvan2512/socker-social-network/shared/utils"
	"github.com/jinzhu/gorm"
)

// Usecase interface
type Usecase interface {
	// Create a post
	Create(r CreateRequest) (postID uint, err error)
}

type usecase struct {
	base.Usecase
	db         *gorm.DB
	repository Repository
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
	post := model.Post{UserID: r.User.ID, Caption: r.Caption}
	if r.PlaceID != "" {

	}
	if r.SourceImageFileName != "" {

	}
	if r.SourceVideoFileName != "" {

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

// NewUsecase creare new instance of Usecase
func NewUsecase(bu *base.Usecase, db *gorm.DB, r Repository) Usecase {
	return &usecase{*bu, db, r}
}
