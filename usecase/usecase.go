package usecase

import (
	"github.com/cetRide/api-rideyu/infrastructure/repository"
	"github.com/cetRide/api-rideyu/model"
)

type RepoHandler struct {
	repository repository.Repository
}

//Returns new NewRepoHandler
func NewRepoHandler(repository repository.Repository) *RepoHandler {

	return &RepoHandler{
		repository: repository,
	}
}

type UseCase interface {
	RegisterUser() model.User
}

