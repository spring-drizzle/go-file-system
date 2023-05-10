package services

import (
	"go-file-system/src/errors"
	"go-file-system/src/models"
	"go-file-system/src/repositories"
	"go-file-system/src/requests"
	"strings"

	"github.com/go-playground/validator/v10"

	"github.com/google/uuid"
)

type userService struct {
	UserRepo  repositories.IUserRepository
	Validator *validator.Validate
}

func NewUserService(repo repositories.IUserRepository, validator *validator.Validate) IUserService {
	return &userService{UserRepo: repo, Validator: validator}
}

func (service *userService) Get(id string) (*models.User, error) {
	return service.UserRepo.Get(id)
}

func (service *userService) GetAll() ([]models.User, error) {
	return service.UserRepo.GetAll()
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
