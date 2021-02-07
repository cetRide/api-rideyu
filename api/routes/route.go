package routers

import (
	"github.com/cetRide/api-rideyu/api/controllers"
	"github.com/cetRide/api-rideyu/api/middleware"
	"github.com/gorilla/mux"
)

func NewRouter(h *controllers.UseCaseHandler) *mux.Router {

	router := mux.NewRouter()
	router.Use(middleware.CORS)
	controllers.GetUserRoutes(router, h)
	return router
}
