package requests

type UserCreateRequest struct {
	Name     string `validate:"required, max=50" json:"name"`
	Surname  string `validate:"required, max=50" json:"surname"`
	Username string `validate:"required, max=50" json:"username"`
	Password string `validate:"required, max=50" json:"password"`
}
