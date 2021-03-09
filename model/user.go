package model

import (
	"database/sql"
	"time"
)

type (
	User struct {
		ID        int64
		Username  string
		Email     string
		Phone     string
		Password  string
		FirstName string
		LastName  string
		CreatedAt time.Time
	}
	Follower struct {
		ID        int64
		Follower  int64
		Following int64
		CreatedAt time.Time
	}
	ListOfFollowers struct {
		FollowerId     int64
		UserId         int64
		Username       string
		ProfilePicture sql.NullString
	}
)
