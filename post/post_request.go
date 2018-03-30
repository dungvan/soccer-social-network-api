package post

import "github.com/dungvan2512/socker-social-network/model"

// CreateRequest struct
type CreateRequest struct {
	Caption             string     `json:"caption" validate:"required_if=SourceImageFileName|required_if=SourceVideoFileName"`
	SourceImageFileName string     `json:"source_image_file_name" validate:"omitempty,gt=0,source_im_name"`
	SourceVideoFileName string     `json:"source_video_file_name" validate:"omitempty,gt=0,source_video_name"`
	PlaceID             string     `json:"place_id" validate:"omitempty,lt=257"`
	Hashtags            []string   `json:"hashtags" validate:"omitempty,max_array_len=30,dive,gt=0,lt=100,hashtag"`
	User                model.User `validate:"required"`
}

// UploadImage struct
type UploadImage struct {
}

// UploadVideo struct
type UploadVideo struct {
}
