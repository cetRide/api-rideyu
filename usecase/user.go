package usecase

import "github.com/cetRide/api-rideyu/model"

func (a *RepoHandler) RegisterUser() model.User {
	return a.repository.CreateAccount()

}
