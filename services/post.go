package services

import (
	"context"
	"errors"
	"net/http"

	"github.com/cetRide/api-rideyu/forms"
	"github.com/cetRide/api-rideyu/model"
	"github.com/cetRide/api-rideyu/utils"
)

func (a *RepoHandler) CreatePost(ctx context.Context, form *forms.PostForm) (map[string]interface{}, error) {
	post := &model.Post{
		UserId:      form.UserId,
		Description: form.Description,
	}

	_, err := a.repository.SavePost(ctx, post)

	if err != nil {
		return nil, utils.NewErrorWithCodeAndMessage(
			err,
			http.StatusInternalServerError,
			"Failed to save post",
			"Failed to save post",
		)
	}

	reponse := make(map[string]interface{})
	reponse["success"] = true
	reponse["message"] = "Post saved"
	return reponse, nil
}

func (a *RepoHandler) CommentPost(ctx context.Context, form *forms.CommentForm) (map[string]interface{}, error) {

	if form.Comment == "" {
		return nil, utils.NewErrorWithCodeAndMessage(
			errors.New("Empty comment!"),
			http.StatusInternalServerError,
			"Failed to save comment",
			"Failed to save comment",
		)
	}
	comment := &model.Comment{
		UserId:  form.UserId,
		PostId:  form.PostId,
		Comment: form.Comment,
	}

	_, err := a.repository.SaveComment(ctx, comment)

	if err != nil {
		return nil, utils.NewErrorWithCodeAndMessage(
			err,
			http.StatusInternalServerError,
			"Failed to save comment",
			"Failed to save comment",
		)
	}

	reponse := make(map[string]interface{})
	reponse["success"] = true
	reponse["message"] = "Comment saved."
	return reponse, nil
}

func (a *RepoHandler) ReplyComment(ctx context.Context, form *forms.CommentForm) (map[string]interface{}, error) {

	if form.Comment == "" {
		return nil, utils.NewErrorWithCodeAndMessage(
			errors.New("Empty reply!"),
			http.StatusInternalServerError,
			"Failed to save reply",
			"Failed to save reply",
		)
	}
	comment := &model.Comment{
		UserId:  form.UserId,
		PostId:  form.PostId,
		Comment: form.Comment,
	}

	_, err := a.repository.ReplyComment(ctx, comment, form.CommentId)

	if err != nil {
		return nil, utils.NewErrorWithCodeAndMessage(
			err,
			http.StatusInternalServerError,
			"Failed to save reply",
			"Failed to save reply",
		)
	}

	reponse := make(map[string]interface{})
	reponse["success"] = true
	reponse["message"] = "Reply saved."
	return reponse, nil
}
