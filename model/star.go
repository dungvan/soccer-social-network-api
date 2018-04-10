package model

// Star struct
type Star struct {
	*User
	Posts []Post `gorm:"many2many:post_stars"`
}
