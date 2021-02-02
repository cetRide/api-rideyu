package controllers

import "github.com/cetRide/api-rideyu/usecase"

type UseCaseHandler struct {
	auseCase usecase.UseCase
}

func NewUseCaseHandler(au usecase.UseCase) *UseCaseHandler {

	return &UseCaseHandler{
		auseCase: au,
	}
}
