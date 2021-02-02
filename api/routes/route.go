package routers

import (
	"github.com/cetRide/api-rideyu/api/controllers"
	"github.com/gorilla/mux"
)

func NewRouter(h *controllers.UseCaseHandler) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	controllers.GetUserRoutes(router, h)
	return router
}
