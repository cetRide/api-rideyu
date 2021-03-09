package model

import "time"

type (
	Post struct {
		ID          int64
		UserId      int64
		Description string
		Likes       int64
		Comments    int64
		Location    string
		CreatedAt   time.Time
	}
	Comment struct {
		ID        int64
		UserId    int64
		Comment   string
		PostId    int64
		Likes     int64
		CreatedAt time.Time
	}
)
