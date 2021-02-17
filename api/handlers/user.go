package handlers

import (
	"net/http"

	"github.com/cetRide/api-rideyu/forms"
	"github.com/gin-gonic/gin"
)

func GetUserRoutes(r *gin.Engine, h *UseCaseHandler) {
	r.POST("/test", h.createUser)
}

func (h *UseCaseHandler) createUser(c *gin.Context) {
	var form forms.UserForm
	err := c.BindJSON(&form)
	if err != nil {
		panic(err)
	}
	user, err := h.service.RegisterUser(c.Request.Context(), &form)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusCreated, user)

}
