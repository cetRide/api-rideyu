package routers

import (
	"net/http"

	handler "github.com/cetRide/api-rideyu/api/handlers"
	"github.com/gin-gonic/gin"
)

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  404,
		"message": "Route Not Found",
	})
}
func NewRouter(h *handler.UseCaseHandler) *gin.Engine {

	router := gin.Default()
	// router.Use(middleware.CORS)
	router.NoRoute(notFound)
	handler.GetUserRoutes(router, h)
	return router
}
