package model

import (
	"database/sql"
	"time"
)

type (
	Post struct {
		ID          int64
		UserId      int64
		Description string
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
	FetchedComment struct {
		ID              int64
		Comment         string
		Username        string
		User_id         string
		CreatedAt       time.Time
		ParentCommentId int64
		ProfilePicture  sql.NullString
		Path            string
	}
	PostModel struct {
		ID             int64
		Description    string
		Username       string
		User_id        string
		CreatedAt      time.Time
		ProfilePicture sql.NullString
	}
	PostMedia struct {
		Id      sql.NullString
		FileUrl sql.NullString
	}
	FetchedPosts struct {
		ID             int64
		Description    string
		Username       string
		User_id        string
		CreatedAt      string
		ProfilePicture sql.NullString
		PostMedia      []PostMedia
	}
)
