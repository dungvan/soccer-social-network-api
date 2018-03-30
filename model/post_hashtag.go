package model

import "github.com/jinzhu/gorm"

// PostHashtag struct
type PostHashtag struct {
	*gorm.Model
	PostID    uint `gorm:"column:post_id;not null"`
	HashtagID uint `gorm:"column:hashtag_id;not null"`
}

// TableName for PostHashtag
func (PostHashtag) TableName() string {
	return "post_hashtags"
}
