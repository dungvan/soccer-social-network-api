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
	Type             string        `json:"type"`
	ImageURLs        interface{}   `json:"image_urls"`
	VideoURLs        interface{}   `json:"video_urls"`
	CreatedAt        time.Time     `json:"created_at"`
	StarCount        uint          `json:"star_count"`
	Comments         []RespComment `json:"comments"`
	StarFlag         bool          `json:"star_flag"`
}

// CreateCommentResponse struct
type CreateCommentResponse struct {
	ID     uint     `json:"id"`
	PostID uint     `json:"post_id"`
	User   RespUser `json:"user"`
}

// RespComment struct
type RespComment struct {
	ID        uint      `json:"id"`
	PostID    uint      `json:"post_id"`
	Content   string    `json:"content"`
	StarCount uint      `json:"star_count"`
	StarFlag  bool      `json:"star_flag"`
	User      RespUser  `json:"user"`
	CreatedAt time.Time `json:"created_at"`
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

// HashtagSearchResponse struct
type HashtagSearchResponse struct {
	Total uint                `json:"total"`
	Posts []RespPosyByHashtag `json:"posts"`
}

// RespPosyByHashtag struct
type RespPosyByHashtag struct {
	ID        uint        `json:"id"`
	User      RespUser    `json:"user"`
	Caption   string      `json:"caption"`
	Type      string      `json:"type"`
	ImageURLs interface{} `json:"image_urls"`
	CreatedAt time.Time   `json:"created_at"`
}
