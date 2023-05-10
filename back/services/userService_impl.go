package services

import (
	"src/errors"
	"src/models"
	"src/repositories"
	"src/requests"
	"src/vendor/github.com/go-playground/validator/v10"
	"strings"

	"github.com/google/uuid"
)

type userService struct {
	UserRepo  repositories.IUserRepository
	Validator *validator.Validate
}

func NewUserService(repo repositories.IUserRepository, validator *validator.Validate) IUserService {
	return &userService{UserRepo: repo, Validator: validator}
}

func (service *userService) Save(request requests.UserCreateRequest) error {
	validateErr := service.Validator.Struct(request)
	if validateErr != nil {
		return &errors.BadRequestError{RequestType: "SaveUser", Err: validateErr}
	}
	uuid := strings.Replace(uuid.New().String(), "-", "", -1)
	user := models.User{Id: uuid, Name: request.Name, Surname: request.Surname, Username: request.Username, Password: request.Password}

	return service.UserRepo.Save(user)
}

func (service *userService) Get(id string) (*models.User, error) {
	return service.UserRepo.Get(id)
}
