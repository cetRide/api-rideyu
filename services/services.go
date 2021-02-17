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
}
