package forms

type (
	PostForm struct {
		UserId      int64
		Description string `json:"description"`
	}
	CommentForm struct {
		UserId    int64
		PostId    int64
		CommentId int64
		Comment   string `json:"comment"`
	}
)
