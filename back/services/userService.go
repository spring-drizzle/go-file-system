package services

import (
	"src/models"
	"src/requests"
)

type IUserService interface {
	Save(request requests.UserCreateRequest) error
	Get(id string) (*models.User, error)
}
