package model

import (
	"github.com/jinzhu/gorm"
)

// Post struct
type Post struct {
	*gorm.Model
	Caption             string `gorm:"column:caption"`
	SourceImageFileName string `gorm:"column:source_image_file_name"`
	SourceVideoFileName string `gorm:"column:source_video_file_name"`
	LocationID          uint   `gorm:"column:location_id"`
}
