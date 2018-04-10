package post

import "time"

// CreateResponse struct
type CreateResponse struct {
	PostID uint `json:"post_id"`
}

// IndexResponse struct
type IndexResponse struct {
	ResultCount int `json:"result_count"`
	Posts       []RespPost
}

// RespPost struct
type RespPost struct {
	ID        uint        `json:"id"`
	UserID    uint        `json:"user_id"`
	Caption   string      `json:"caption"`
	ImageURLs interface{} `json:"image_url"`
	VideoURLs interface{} `json:"video_url"`
	CreatedAt time.Time   `json:"created_at"`
}

// StarCountResponse responses from CountUpStar and CountDownStar function.
// JSON responses payload structure
type StarCountResponse struct {
	TypeOfStatusCode int  `json:"-"`
	StarCount        uint `json:"star_count"`
}
