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
	ID             uint        `json:"id"`
	UserID         uint        `json:"user_id"`
	Caption        string      `json:"caption"`
	SourceImageURL interface{} `json:"source_image_url"`
	SourceVideoURL interface{} `json:"source_video_url"`
	CreatedAt      time.Time   `json:"created_at"`
}
