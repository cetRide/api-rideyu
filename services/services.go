package services

import (
	"context"

	"github.com/cetRide/api-rideyu/forms"
	repo "github.com/cetRide/api-rideyu/infrastructure/repository"
	"github.com/cetRide/api-rideyu/model"
)

type RepoHandler struct {
	repository repo.Repository
}

//Returns new NewRepoHandler
func NewRepoHandler(repository repo.Repository) *RepoHandler {

	return &RepoHandler{
		repository: repository,
	}
}

type UseCase interface {
	RegisterUser(context.Context, *forms.UserForm) (*model.User, error)
	Login(context.Context, *forms.LoginForm) (*model.User, error)
	Follow(context.Context, *forms.FollowersForm) (*model.Follower, error)
	UnFollow(context.Context, *forms.FollowersForm) (map[string]interface{}, error)
	ViewListOfFollowing(context.Context, *forms.FollowersForm) ([]*model.ListOfFollowers, error)
	ViewListOfFollowers(context.Context, *forms.FollowersForm) ([]*model.ListOfFollowers, error)
}
