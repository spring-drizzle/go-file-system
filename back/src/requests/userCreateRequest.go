package requests

type UserCreateRequest struct {
	Name     string `validate:"required,min=1,max=50" json:"name"`
	Surname  string `validate:"required,min=1,max=50" json:"surname"`
	Username string `validate:"required,min=1,max=50" json:"username"`
	Password string `validate:"required,min=1,max=50" json:"password"`
}
