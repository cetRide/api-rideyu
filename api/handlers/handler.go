package handlers

import service "github.com/cetRide/api-rideyu/services"

type UseCaseHandler struct {
	service service.UseCase
}

func NewUseCaseHandler(usecase service.UseCase) *UseCaseHandler {

	return &UseCaseHandler{
		service: usecase,
	}
}
