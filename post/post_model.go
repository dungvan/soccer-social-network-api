package post

import (
	"time"

	"github.com/dungvan2512/soccer-social-network-api/model"
)

// Comment struct table
type Comment struct {
	ID        uint
	Content   string
	UserID    uint
	UserName  string
	FirstName string
	LastName  string
	StarCount uint
	StarFlag  bool
	CreatedAt time.Time
}

// Post struct table
type Post struct {
	*model.Post
	UserName  string
	FirstName string
	LastName  string
	StarCount uint
	StarFlag  bool
}
