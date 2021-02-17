package repository

import (
	"context"

	"github.com/cetRide/api-rideyu/model"
)

type Repository interface {
	CreateAccount(context.Context, *model.User) (*model.User, error)
}
