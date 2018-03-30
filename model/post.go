package model

import (
	"github.com/jinzhu/gorm"
)

// LocationID type
type LocationID uint

// Post struct
type Post struct {
	*gorm.Model
	UserID              uint        `gorm:"column:user_id"`
	Caption             string      `gorm:"column:caption"`
	SourceImageFileName string      `gorm:"column:source_image_file_name"`
	SourceVideoFileName string      `gorm:"column:source_video_file_name"`
	LocationID          *LocationID `gorm:"column:location_id"`
}
