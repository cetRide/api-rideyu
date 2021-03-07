package handlers

import (
	"net/http"

	"github.com/cetRide/api-rideyu/forms"
	"github.com/cetRide/api-rideyu/utils"
	"github.com/gin-gonic/gin"
)

func (h *UseCaseHandler) createUser(c *gin.Context) {
	var form forms.UserForm
	err := c.BindJSON(&form)
	if err != nil {
		appError := utils.NewErrorWithCodeAndMessage(
			err,
			http.StatusBadRequest,
			"Invalid form provided",
			"Failed to bind to user form",
		)

		appError.LogErrorMessages()

		c.JSON(appError.HttpStatus(), appError.JsonResponse())
		return
	}

	user, err := h.service.RegisterUser(c.Request.Context(), &form)
	if err != nil {
		appError := utils.NewError(
			err,
			"Failed to create user",
		)

		appError.LogErrorMessages()

		c.JSON(appError.HttpStatus(), appError.JsonResponse())
		return
	}
	c.JSON(http.StatusCreated, user)

}

func (h *UseCaseHandler) login(c *gin.Context) {
	var form forms.LoginForm
	err := c.BindJSON(&form)
	if err != nil {
		appError := utils.NewErrorWithCodeAndMessage(
			err,
			http.StatusBadRequest,
			"Invalid form provided",
			"Failed to bind to login user form",
		)

		appError.LogErrorMessages()

		c.JSON(appError.HttpStatus(), appError.JsonResponse())
		return
	}

	user, err := h.service.Login(c.Request.Context(), &form)
	if err != nil {
		appError := utils.NewError(
			err,
			"Failed to create user",
		)

		appError.LogErrorMessages()

		c.JSON(appError.HttpStatus(), appError.JsonResponse())
		return
	}
	c.JSON(http.StatusCreated, user)

}

func (h *UseCaseHandler) followUser(c *gin.Context) {
	var form forms.FollowersForm
	err := c.BindJSON(&form)
	if err != nil {
		appError := utils.NewErrorWithCodeAndMessage(
			err,
			http.StatusBadRequest,
			"Invalid form provided",
			"Failed to bind to user form",
		)

		appError.LogErrorMessages()

		c.JSON(appError.HttpStatus(), appError.JsonResponse())
		return
	}

	follow, err := h.service.Follow(c.Request.Context(), &form)
	if err != nil {
		appError := utils.NewError(
			err,
			"Failed to follow",
		)

		appError.LogErrorMessages()

		c.JSON(appError.HttpStatus(), appError.JsonResponse())
		return
	}
	c.JSON(http.StatusCreated, follow)

}

func (h *UseCaseHandler) unFollowUser(c *gin.Context) {
	var form forms.FollowersForm
	err := c.BindJSON(&form)
	if err != nil {
		appError := utils.NewErrorWithCodeAndMessage(
			err,
			http.StatusBadRequest,
			"Invalid form provided",
			"Failed to bind to user form",
		)

		appError.LogErrorMessages()

		c.JSON(appError.HttpStatus(), appError.JsonResponse())
		return
	}

	response, err := h.service.UnFollow(c.Request.Context(), &form)
	if err != nil {
		appError := utils.NewError(
			err,
			"Failed to unfollow",
		)

		appError.LogErrorMessages()

		c.JSON(appError.HttpStatus(), appError.JsonResponse())
		return
	}
	c.JSON(http.StatusCreated, response)
}

func (h *UseCaseHandler) viewFollowing(c *gin.Context) {
	var form forms.FollowersForm
	err := c.BindJSON(&form)
	if err != nil {
		appError := utils.NewErrorWithCodeAndMessage(
			err,
			http.StatusBadRequest,
			"Invalid form provided",
			"Failed to bind to user form",
		)

		appError.LogErrorMessages()

		c.JSON(appError.HttpStatus(), appError.JsonResponse())
		return
	}

	response, err := h.service.ViewListOfFollowers(c.Request.Context(), &form)
	if err != nil {
		appError := utils.NewError(
			err,
			"Failed to retrieve following",
		)

		appError.LogErrorMessages()

		c.JSON(appError.HttpStatus(), appError.JsonResponse())
		return
	}
	c.JSON(http.StatusCreated, response)
}

func (h *UseCaseHandler) viewFollowers(c *gin.Context) {
	var form forms.FollowersForm
	err := c.BindJSON(&form)
	if err != nil {
		appError := utils.NewErrorWithCodeAndMessage(
			err,
			http.StatusBadRequest,
			"Invalid form provided",
			"Failed to bind to user form",
		)

		appError.LogErrorMessages()

		c.JSON(appError.HttpStatus(), appError.JsonResponse())
		return
	}

	response, err := h.service.ViewListOfFollowers(c.Request.Context(), &form)
	if err != nil {
		appError := utils.NewError(
			err,
			"Failed to retrieve followers",
		)

		appError.LogErrorMessages()

		c.JSON(appError.HttpStatus(), appError.JsonResponse())
		return
	}
	
	c.JSON(http.StatusCreated, response)
}

func GetUserRoutes(r *gin.Engine, h *UseCaseHandler) {
	r.POST("/create-account", h.createUser)
	r.GET("/login", h.login)
	r.POST("/follow", h.followUser)
	r.DELETE("/unfollow", h.unFollowUser)
	r.GET("/view-following", h.viewFollowing)
	r.GET("/view-followers", h.viewFollowers)
}
