package repository

import "github.com/cetRide/api-rideyu/model"

type Repository interface {
	CreateAccount() (model.User)
}