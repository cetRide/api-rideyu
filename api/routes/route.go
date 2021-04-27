package routers

import (
	"net/http"

	handler "github.com/cetRide/api-rideyu/api/handlers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
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
	//router.Use(middleware.CORS())
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
	router.NoRoute(notFound)
	handler.GetUserRoutes(router, h)
	handler.PostsRoutes(router, h)
	return router
}
