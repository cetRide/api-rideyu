package handlers

import (
	"net/http"
	"strconv"

	"github.com/cetRide/api-rideyu/forms"
	"github.com/cetRide/api-rideyu/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (h *UseCaseHandler) createPost(c *gin.Context) {

	var form forms.PostForm
	err := c.BindJSON(&form)
	if err != nil {
		appError := utils.NewErrorWithCodeAndMessage(
			err,
			http.StatusBadRequest,
			"Invalid form provided",
			"Failed to bind to user request",
		)

		appError.LogErrorMessages()

		c.JSON(appError.HttpStatus(), appError.JsonResponse())
		return
	}
	session := sessions.Default(c)
	form.UserId = session.Get("id").(int64)
	response, err := h.service.CreatePost(c.Request.Context(), &form)
	if err != nil {
		appError := utils.NewError(
			err,
			"Failed to save post",
		)

		appError.LogErrorMessages()

		c.JSON(appError.HttpStatus(), appError.JsonResponse())
		return
	}
	c.JSON(http.StatusCreated, response)
}

func (h *UseCaseHandler) commentPost(c *gin.Context) {

	var form forms.CommentForm
	err := c.BindJSON(&form)
	if err != nil {
		appError := utils.NewErrorWithCodeAndMessage(
			err,
			http.StatusBadRequest,
			"Invalid form provided",
			"Failed to bind to user request",
		)

		appError.LogErrorMessages()

		c.JSON(appError.HttpStatus(), appError.JsonResponse())
		return
	}
	session := sessions.Default(c)
	form.UserId = session.Get("id").(int64)
	postId, err := strconv.ParseInt(c.Params.ByName("post-id"), 10, 64)

	if err != nil {
		appError := utils.NewError(
			err,
			"Failed to save comment",
		)

		appError.LogErrorMessages()

		c.JSON(appError.HttpStatus(), appError.JsonResponse())
		return
	}

	form.PostId = postId
	response, err := h.service.CommentPost(c.Request.Context(), &form)
	if err != nil {
		appError := utils.NewError(
			err,
			"Failed to save comment",
		)

		appError.LogErrorMessages()

		c.JSON(appError.HttpStatus(), appError.JsonResponse())
		return
	}
	c.JSON(http.StatusCreated, response)
}

func (h *UseCaseHandler) replyComment(c *gin.Context) {

	var form forms.CommentForm
	err := c.BindJSON(&form)
	if err != nil {
		appError := utils.NewErrorWithCodeAndMessage(
			err,
			http.StatusBadRequest,
			"Invalid form provided",
			"Failed to bind to user request",
		)

		appError.LogErrorMessages()

		c.JSON(appError.HttpStatus(), appError.JsonResponse())
		return
	}
	session := sessions.Default(c)
	form.UserId = session.Get("id").(int64)
	postId, err := strconv.ParseInt(c.Params.ByName("post-id"), 10, 64)

	if err != nil {
		appError := utils.NewError(
			err,
			"Failed to save reply",
		)

		appError.LogErrorMessages()

		c.JSON(appError.HttpStatus(), appError.JsonResponse())
		return
	}

	commentId, err := strconv.ParseInt(c.Params.ByName("comment-id"), 10, 64)

	if err != nil {
		appError := utils.NewError(
			err,
			"Failed to save reply",
		)

		appError.LogErrorMessages()

		c.JSON(appError.HttpStatus(), appError.JsonResponse())
		return
	}

	form.PostId = postId
	form.CommentId = commentId
	response, err := h.service.ReplyComment(c.Request.Context(), &form)
	if err != nil {
		appError := utils.NewError(
			err,
			"Failed to save reply",
		)

		appError.LogErrorMessages()

		c.JSON(appError.HttpStatus(), appError.JsonResponse())
		return
	}
	c.JSON(http.StatusCreated, response)
}

func PostsRoutes(r *gin.Engine, h *UseCaseHandler) {
	r.Group("/post")
	r.POST("/save-post", h.createPost)
	r.POST("/comment-post/:post-id", h.commentPost)
	r.POST("/reply-comment/:post-id/:comment-id", h.replyComment)
}
