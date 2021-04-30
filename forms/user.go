package forms

type (
	UserForm struct {
		Username string `binding:"required" json:"username"`
		Email    string `binding:"required" json:"email"`
		// Phone    string `binding:"required" json:"phone"`
		Password string `binding:"required" json:"password"`
		// FirstName string `binding:"required" json:"firstname"`
		// LastName  string `binding:"required" json:"lastname"`
	}

	LoginForm struct {
		UsernameOrEmail string `binding:"required" json:"usernameoremail"`
		Password        string `binding:"required" json:"password"`
	}
	FollowersForm struct {
		ID        int64 `json:"id"`
		Follower  int64 `json:"follower"`
		Following int64 `json:"following"`
	}
)
