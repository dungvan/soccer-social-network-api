package post

import "github.com/dungvan2512/socker-social-network/model"

// CreateRequest struct
type CreateRequest struct {
	Caption  string     `json:"caption" validate:"required_if=Images|required_if=Videos"`
	Images   []string   `json:"source_image_file_name" validate:"omitempty,dive,gt=0,image_name"`
	Videos   []string   `json:"source_video_file_name" validate:"omitempty,dive,gt=0,video_name"`
	PlaceID  string     `json:"place_id" validate:"omitempty,lt=257"`
	Tags     []string   `json:"tags" validate:"omitempty,max_array_len=30,dive,gt=0"`
	Hashtags []string   `json:"hashtags" validate:"omitempty,max_array_len=30,dive,gt=0,lt=100,hashtag"`
	User     model.User `validate:"required"`
}

// UploadImage struct
type UploadImage struct {
}

// UploadVideo struct
type UploadVideo struct {
}

// UpdateRequest struct
type UpdateRequest struct {
	PostID  uint   `validate:"required"`
	Caption string `validate:"retuired"`
}

// StarCountRequest struct.
type StarCountRequest struct {
	PostID uint `validate:"required"`
	UserID uint `validate:"required"`
}
