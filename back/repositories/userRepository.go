package repositories

import "src/models"

type IUserRepository interface {
	Save(user models.User) error
	Get(id string) (*models.User, error)
}
