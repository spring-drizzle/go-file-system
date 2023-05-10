package services

import (
	"go-file-system/src/models"
	"go-file-system/src/requests"
)

type IUserService interface {
	Save(request requests.UserCreateRequest) error
	Get(id string) (*models.User, error)
	GetAll() ([]models.User, error)
}
