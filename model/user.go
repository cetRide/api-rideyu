package model

import "database/sql"

type (
	User struct {
		ID        int64
		Username  string
		Email     string
		Phone     string
		Password  string
		FirstName string
		LastName  string
	}
	Follower struct {
		ID        int64
		Follower  int64
		Following int64
	}
	ListOfFollowers struct {
		FollowerId     int64
		UserId         int64
		Username       string
		ProfilePicture sql.NullString
	}
)
