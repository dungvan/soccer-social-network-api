package model

import "github.com/jinzhu/gorm"

// PostHashtag struct
type PostHashtag struct {
	*gorm.Model
	PostID    uint
	HashtagID uint
}
