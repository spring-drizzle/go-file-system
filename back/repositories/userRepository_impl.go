package repositories

import (
	"src/errors"
	"src/models"

	"src/vendor/gorm.io/gorm"
)

const entityType = "user"

type userRepository struct {
	database *gorm.DB
}

func NewUserRepo(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

func (repo *userRepository) Save(user models.User) error {
	id := repo.database.Create(user)
	if id == nil {
		return &errors.FailedInsertError{Entity: entityType}
	}
	return nil
}

func (repo *userRepository) Get(id string) (*models.User, error) {
	var user *models.User
	repo.database.Find(user, id)
	if user != nil {
		return user, nil
	}

	return nil, &errors.NotFoundError{Entity: entityType, Id: id}
}
