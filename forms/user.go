package forms

type UserForm struct {
	FirstName string `binding:"required" json:"firstname"`
	LastName  string `binding:"required" json:"lastname"`
}
