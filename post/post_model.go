package post

import "time"

// Comment struct table
type Comment struct {
	ID        uint
	Content   string
	UserID    uint
	UserName  string
	FirstName string
	LastName  string
	CreatedAt time.Time
}

// Post struct table
type Post struct {
	ID        uint
	Caption   string
	UserID    uint
	UserName  string
	FirstName string
	LastName  string
	CreatedAt time.Time
}
