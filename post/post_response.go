package post

import (
	"time"
)

// CreateResponse struct
type CreateResponse struct {
	PostID uint `json:"post_id"`
}

// IndexResponse struct
type IndexResponse struct {
	TypeOfStatusCode int        `json:"-"`
	Total            uint       `json:"total"`
	Posts            []RespPost `json:"posts"`
}

// RespPost struct
type RespPost struct {
	TypeOfStatusCode int           `json:"-"`
	ID               uint          `json:"id"`
	User             RespUser      `json:"user"`
	Caption          string        `json:"caption"`
	ImageURLs        interface{}   `json:"image_url"`
	VideoURLs        interface{}   `json:"video_url"`
	CreatedAt        time.Time     `json:"created_at"`
	StarCount        uint          `json:"star_count"`
	Comments         []RespComment `json:"comments"`
}

// CreateCommentResponse struct
type CreateCommentResponse struct {
	CommentID uint `json:"comment_id"`
}

// RespComment struct
type RespComment struct {
	ID      uint     `json:"id"`
	Content string   `json:"content"`
	User    RespUser `json:"user"`
}

// RespUser struct
type RespUser struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// StarCountResponse responses from CountUpStar and CountDownStar function.
// JSON responses payload structure
type StarCountResponse struct {
	TypeOfStatusCode int  `json:"-"`
	StarCount        uint `json:"star_count"`
}

// UploadImagesResponse struct
type UploadImagesResponse struct {
	ImageNames []string `json:"image_names"`
}
