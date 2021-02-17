package services

import (
	"context"

	"github.com/cetRide/api-rideyu/forms"
	"github.com/cetRide/api-rideyu/model"
)

func (a *RepoHandler) RegisterUser(ctx context.Context, form *forms.UserForm) (*model.User, error) {

	user := &model.User{
		FirstName: form.FirstName,
		LastName:  form.LastName,
	}

	results, err := a.repository.CreateAccount(ctx, user)

	if err != nil {
		return nil, err
	}

	return results, nil
}
