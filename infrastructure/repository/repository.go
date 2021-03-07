package repository

import (
	"context"

	"github.com/cetRide/api-rideyu/model"
)

type Repository interface {
	user
}

type user interface {
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
