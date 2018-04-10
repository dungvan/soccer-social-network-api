package model

import "github.com/jinzhu/gorm"

// Hashtag struct
type Hashtag struct {
	*gorm.Model
	KeyWord string
	Posts   []Post `gorm:"many2many:post_hashtags"`
}
