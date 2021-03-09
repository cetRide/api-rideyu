package repository

import (
	"context"
	"database/sql"

	"github.com/cetRide/api-rideyu/model"
)

type Repository interface {
	user
	post
}

type (
	user interface {
		SaveUser(context.Context, *model.User) (*model.User, error)
		FindByEmail(context.Context, string) (*model.User, error)
		FindByPhone(context.Context, string) (*model.User, error)
		FindByUsername(context.Context, string) (*model.User, error)
		FindById(context.Context, int64) (*model.User, error)
		Follow(context.Context, int64, int64) (*model.Follower, error)
		UnFollow(ctx context.Context, follower int64) (map[string]interface{}, error)
		ViewListOfFollowing(ctx context.Context, id int64) ([]*model.ListOfFollowers, error)
		ViewListOfFollowers(ctx context.Context, id int64) ([]*model.ListOfFollowers, error)
	}

	post interface {
		SavePost(context.Context, *model.Post) (sql.Result, error)
		SaveComment(context.Context, *model.Comment) (sql.Result, error)
		ReplyComment(context.Context, *model.Comment, int64) (sql.Result, error)
	}
)
