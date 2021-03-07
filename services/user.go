package services

import (
	"context"
	"errors"
	"net/http"
	"regexp"

	"github.com/cetRide/api-rideyu/forms"
	"github.com/cetRide/api-rideyu/model"
	"github.com/cetRide/api-rideyu/utils"
	"golang.org/x/crypto/bcrypt"
)

//validations
func validateEmail(email string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	return re.MatchString(email)
}

func validatePassword(password string) map[string]interface{} {

	r, _ := regexp.Compile(`[A-Z]`)
	if !r.MatchString(password) {
		return map[string]interface{}{
			"status": false, "message": "Password should contain a uppercase letter."}
	}

	r, _ = regexp.Compile(`[a-z]`)
	if !r.MatchString(password) {
		return map[string]interface{}{
			"status": false, "message": "Password should contain a lowercase letter."}

	}

	r, _ = regexp.Compile(`[0-9]`)
	if !r.MatchString(password) {
		return map[string]interface{}{
			"status": false, "message": "Password should contain a digit."}
	}
	if len(password) < 8 {
		return map[string]interface{}{
			"status": false, "message": "Password should be atleast 8 characters."}
	}

	return map[string]interface{}{
		"status": true}
}

func (a *RepoHandler) RegisterUser(ctx context.Context, form *forms.UserForm) (*model.User, error) {

	if form.Username == "" {
		return nil, utils.NewErrorWithCodeAndMessage(
			errors.New("username cannot be empty"),
			http.StatusBadRequest,
			"Username cannot be empty",
			"Username=[%v] cannot be empty",
			form.Username,
		)
	}

	if form.Email == "" {
		return nil, utils.NewErrorWithCodeAndMessage(
			errors.New("email address cannot be empty"),
			http.StatusBadRequest,
			"email address cannot be empty",
			"Email=[%v] cannot be empty",
			form.Email,
		)
	}

	if form.Phone == "" {
		return nil, utils.NewErrorWithCodeAndMessage(
			errors.New("Phone number cannot be empty"),
			http.StatusBadRequest,
			"Phone number cannot be empty",
			"Phone=[%v] cannot be empty",
			form.Phone,
		)
	}

	if len(form.Username) > 50 {
		return nil, utils.NewErrorWithCodeAndMessage(
			errors.New("invalid username"),
			http.StatusBadRequest,
			"Username cannot be more than 50 characters",
			"Username=[%v] provided is more that 50 characters",
			form.Username,
		)
	}

	if len(form.Username) > 50 {
		return nil, utils.NewErrorWithCodeAndMessage(
			errors.New("invalid email address"),
			http.StatusBadRequest,
			"Email cannot be more than 50 characters",
			"Email=[%v] provided is more that 50 characters",
			form.Username,
		)
	}

	if !validateEmail(form.Email) {
		return nil, utils.NewErrorWithCodeAndMessage(
			errors.New("invalid email address"),
			http.StatusBadRequest,
			"Invalid email address",
			"Email=[%v] provided is invalid",
			form.Username,
		)
	}

	passwordValidation := validatePassword(form.Password)
	if !passwordValidation["status"].(bool) {
		return nil, utils.NewErrorWithCodeAndMessage(
			errors.New(passwordValidation["message"].(string)),
			http.StatusBadRequest,
			passwordValidation["message"].(string),
			"Password=[%v] provided is invalid",
			form.Username,
		)
	}

	// find phone
	_, err := a.repository.FindByPhone(ctx, form.Phone)

	if err != nil && !utils.IsErrNoRows(err) {
		return nil, utils.NewErrorWithCodeAndMessage(
			err,
			http.StatusBadRequest,
			"Failed to validate user phone number",
			"Phone=[%v] failed to be validated",
			form.Phone,
		)
	}

	if err == nil {
		return nil, utils.NewErrorWithCodeAndMessage(
			errors.New("Phone number already used"),
			http.StatusConflict,
			"Phone number already used",
			"phone=[%v] provided is not available",
			form.Phone,
		)
	}

	//find by email
	_, err = a.repository.FindByEmail(ctx, form.Email)

	if err != nil && !utils.IsErrNoRows(err) {
		return nil, utils.NewErrorWithCodeAndMessage(
			err,
			http.StatusBadRequest,
			"Failed to validate user email address",
			"Email=[%v] failed to be validated",
			form.Email,
		)
	}

	if err == nil {
		return nil, utils.NewErrorWithCodeAndMessage(
			errors.New("Email address already used"),
			http.StatusConflict,
			"Email address already used",
			"email=[%v] provided is not available",
			form.Email,
		)
	}

	//find by username
	_, err = a.repository.FindByUsername(ctx, form.Username)

	if err != nil && !utils.IsErrNoRows(err) {
		return nil, utils.NewErrorWithCodeAndMessage(
			err,
			http.StatusBadRequest,
			"Failed to validate user username",
			"Username=[%v] failed to be validated",
			form.Username,
		)
	}

	if err == nil {
		return nil, utils.NewErrorWithCodeAndMessage(
			errors.New("Username already used"),
			http.StatusConflict,
			"Username already used",
			"username=[%v] provided is not available",
			form.Username,
		)
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)

	saltedPassword := string(hashedPassword)

	user := &model.User{
		Username: form.Username,
		Email:    form.Email,
		Phone:    form.Phone,
		Password: saltedPassword,
	}

	results, err := a.repository.SaveUser(ctx, user)

	if err != nil {
		return nil, utils.NewErrorWithCodeAndMessage(
			err,
			http.StatusInternalServerError,
			"Failed to save user details",
			"Failed to save user details",
		)
	}

	return results, nil
}

func (a *RepoHandler) Login(ctx context.Context, form *forms.LoginForm) (*model.User, error) {

	if form.UsernameOrEmail == "" {
		return nil, utils.NewErrorWithCodeAndMessage(
			errors.New("Username or Email field is empty"),
			http.StatusBadRequest,
			"Username or Email field is empty",
			"UsernameOrEmail=[%v] cannot be empty",
			form.UsernameOrEmail,
		)
	}

	isEmail := validateEmail(form.UsernameOrEmail)

	if isEmail {

		user, err := a.repository.FindByEmail(ctx, form.UsernameOrEmail)

		if err != nil && !utils.IsErrNoRows(err) {
			return nil, utils.NewErrorWithCodeAndMessage(
				err,
				http.StatusInternalServerError,
				"Failed to find user email address",
				"Email=[%v] failed to find user email",
				form.UsernameOrEmail,
			)
		}

		if err != nil && utils.IsErrNoRows(err) {
			return nil, utils.NewErrorWithCodeAndMessage(
				errors.New("User not found"),
				http.StatusNotFound,
				"User not found",
				"email=[%v] provided is not available",
				form.UsernameOrEmail,
			)
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password))

		if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
			return nil, utils.NewErrorWithCodeAndMessage(
				errors.New("Incorrect password"),
				http.StatusForbidden,
				"Incorrect password",
				"password=[%v] provided is not available",
				form.Password,
			)

		}
		if err != nil && err != bcrypt.ErrMismatchedHashAndPassword {
			return nil, utils.NewErrorWithCodeAndMessage(
				err,
				http.StatusInternalServerError,
				"Internal server error",
				"password=[%v] Internal server error",
				form.Password,
			)

		}

		return user, nil

	} else {
		user, err := a.repository.FindByUsername(ctx, form.UsernameOrEmail)

		if err != nil && !utils.IsErrNoRows(err) {
			return nil, utils.NewErrorWithCodeAndMessage(
				err,
				http.StatusInternalServerError,
				"Failed to find username",
				"Username=[%v] failed to find username",
				form.UsernameOrEmail,
			)
		}

		if err != nil && utils.IsErrNoRows(err) {
			return nil, utils.NewErrorWithCodeAndMessage(
				errors.New("Username not found"),
				http.StatusNotFound,
				"User not found",
				"Username=[%v] provided is not available",
				form.UsernameOrEmail,
			)
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password))

		if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
			return nil, utils.NewErrorWithCodeAndMessage(
				errors.New("Incorrect password"),
				http.StatusForbidden,
				"Incorrect password",
				"password=[%v] provided is not available",
				form.Password,
			)

		}

		if err != nil && err != bcrypt.ErrMismatchedHashAndPassword {
			return nil, utils.NewErrorWithCodeAndMessage(
				err,
				http.StatusInternalServerError,
				"Internal server error",
				"password=[%v] Internal server error",
				form.Password,
			)

		}

		return user, nil

	}

}

func (a *RepoHandler) Follow(ctx context.Context, form *forms.FollowersForm) (*model.Follower, error) {
	if form.Follower == 0 {
		return nil, utils.NewErrorWithCodeAndMessage(
			errors.New("Empty follower id"),
			http.StatusBadRequest,
			"Empty follower Id",
			"Follower ID=[%v] is empty",
			form.Follower,
		)
	}

	if form.Follower == 0 {
		return nil, utils.NewErrorWithCodeAndMessage(
			errors.New("Empty user id"),
			http.StatusBadRequest,
			"Empty user Id",
			"User ID=[%v] is empty",
			form.Following,
		)
	}

	followers, err := a.repository.Follow(ctx, form.Follower, form.Following)

	if err != nil {
		return nil, utils.NewErrorWithCodeAndMessage(
			err,
			http.StatusInternalServerError,
			"Internal server error",
			"Follower ID=[%v] Following=[%v] is empty",
			form.Follower, form.Following,
		)
	}
	return followers, nil
}

func (a *RepoHandler) UnFollow(ctx context.Context, form *forms.FollowersForm) (map[string]interface{}, error) {

	if form.ID == 0 {
		return nil, utils.NewErrorWithCodeAndMessage(
			errors.New("Empty id"),
			http.StatusBadRequest,
			"Empty Id",
			"ID=[%v] is empty",
			form.ID,
		)
	}

	results, err := a.repository.UnFollow(ctx, form.ID)
	if err != nil {
		return nil, utils.NewErrorWithCodeAndMessage(
			err,
			http.StatusBadRequest,
			"Unable to unfollow",
			"ID=[%v] : unable to unfollow",
			form.ID,
		)
	}

	return results, nil
}

func (a *RepoHandler) ViewListOfFollowing(ctx context.Context, form *forms.FollowersForm) ([]*model.ListOfFollowers, error) {
	if form.ID == 0 {
		return nil, utils.NewErrorWithCodeAndMessage(
			errors.New("Empty id"),
			http.StatusBadRequest,
			"Empty Id",
			"ID=[%v] is empty",
			form.ID,
		)
	}
	following, err := a.repository.ViewListOfFollowing(ctx, form.Follower)

	if err != nil {
		return nil, utils.NewErrorWithCodeAndMessage(
			err,
			http.StatusBadRequest,
			"Unable to retrieve a list of following",
			"ID=[%v] : retrieve a list of following",
			form.ID,
		)
	}
	return following, nil
}

func (a *RepoHandler) ViewListOfFollowers(ctx context.Context, form *forms.FollowersForm) ([]*model.ListOfFollowers, error) {
	if form.ID == 0 {
		return nil, utils.NewErrorWithCodeAndMessage(
			errors.New("Empty id"),
			http.StatusBadRequest,
			"Empty Id",
			"ID=[%v] is empty",
			form.ID,
		)
	}
	followers, err := a.repository.ViewListOfFollowers(ctx, form.Following)

	if err != nil {
		return nil, utils.NewErrorWithCodeAndMessage(
			err,
			http.StatusBadRequest,
			"Unable to retrieve a list of followers",
			"ID=[%v] : retrieve a list of followers",
			form.ID,
		)
	}

	return followers, nil
}
