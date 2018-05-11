package post

import (
	"mime/multipart"
)

// IndexRequest struct
type IndexRequest struct {
	Page uint `form:"page" validate:"omitempty,min=1"`
}

// CreateRequest struct
type CreateRequest struct {
	Caption  string   `json:"caption" validate:"required_if=Images|required_if=Videos"`
	Images   []string `json:"image_names" validate:"omitempty,dive,gt=0,image_name"`
	Videos   []string `json:"video_names" validate:"omitempty,dive,gt=0,video_name"`
	PlaceID  string   `json:"place_id" validate:"omitempty,lt=257"`
	Tags     []string `json:"tags" validate:"omitempty,max_array_len=30,dive,gt=0"`
	Hashtags []string `json:"hashtags" validate:"omitempty,max_array_len=30,dive,gt=0,lt=100,hashtag"`
	UserID   uint     `validate:"required"`
}

// UploadImagesRequest struct
type UploadImagesRequest struct {
	Images []Image `validate:"dive"`
}

// UploadVideosRequest struct
type UploadVideosRequest struct {
}

// UpdateRequest struct
type UpdateRequest struct {
	ID      uint   `validate:"required"`
	Caption string `json:"caption" validate:"required"`
}

// StarCountRequest struct.
type StarCountRequest struct {
	ID     uint `validate:"required"`
	UserID uint `validate:"required"`
}

// Image struct
type Image struct {
	Body     multipart.File `form:"image file" validate:"required"`
	MimeType string         `form:"image type" validate:"omitempty,eq=image/bmp|eq=image/dib|eq=image/jpeg|eq=image/jp2|eq=image/png|eq=image/webp|eq=image/x-portable-anymap|eq=image/x-portable-bitmap|eq=image/x-portable-graymap|eq=image/x-portable-pixmap|eq=image/x-cmu-raster|eq=image/tiff|eq=image/gif"`
	Size     int64          `form:"image size" validate:"omitempty,gt=0,max=10485760"`
	Name     string
}

//===========================================
//==================COMMENT==================
//===========================================

// CreateCommentRequest struct
type CreateCommentRequest struct {
	Content string `json:"caption" validate:"required"`
	PostID  uint   `json:"post_id" validate:"required"`
	UserID  uint   `validate:"required"`
}

// UpdateCommentRequest struct
type UpdateCommentRequest struct {
	ID      uint   `validate:"required"`
	Content string `json:"content" validate:"required"`
}
