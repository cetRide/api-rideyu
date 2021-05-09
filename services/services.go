package services

import (
	"context"

	"github.com/cetRide/api-rideyu/forms"
	repo "github.com/cetRide/api-rideyu/infrastructure/repository"
	"github.com/cetRide/api-rideyu/model"
	"github.com/gin-contrib/sessions"
)

type RepoHandler struct {
	repository repo.Repository
}

//Returns new NewRepoHandler
func NewRepoHandler(repository repo.Repository) *RepoHandler {

	return &RepoHandler{
		repository: repository,
	}
}

type UseCase interface {
	user
	post
}
type (
	user interface {
		RegisterUser(context.Context, *forms.UserForm) (map[string]interface{}, error)
		Login(context.Context, *forms.LoginForm, sessions.Session) (map[string]interface{}, error)
		Logout(context.Context, sessions.Session) (map[string]interface{}, error)
		Follow(context.Context, *forms.FollowersForm) (*model.Follower, error)
		UnFollow(context.Context, *forms.FollowersForm) (map[string]interface{}, error)
		ViewListOfFollowing(context.Context, *forms.FollowersForm) ([]*model.ListOfFollowers, error)
		ViewListOfFollowers(context.Context, *forms.FollowersForm) ([]*model.ListOfFollowers, error)
	}

	post interface {
		CreatePost(context.Context, *forms.PostForm) (map[string]interface{}, error)
		CommentPost(context.Context, *forms.CommentForm) (map[string]interface{}, error)
		ReplyComment(context.Context, *forms.CommentForm) (map[string]interface{}, error)
		FetchComments(context.Context, int64) (map[string]interface{}, error)
		FetchPosts(context.Context) (map[string]interface{}, error)
		FetchPostCommentsCount(context.Context, string) (map[string]interface{}, error)
		FetchPostLikesCount(context.Context, string) (map[string]interface{}, error)
	}
)
