package repositories

import "go-file-system/src/models"

type IUserRepository interface {
	Save(user models.User) error
	Get(id string) (*models.User, error)
	GetAll() ([]models.User, error)
}
