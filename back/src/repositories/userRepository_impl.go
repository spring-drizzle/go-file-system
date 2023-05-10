package repositories

import (
	"go-file-system/src/errors"
	"go-file-system/src/models"

	"gorm.io/gorm"
)

const entityType = "user"

type userRepository struct {
	database *gorm.DB
}

func NewUserRepo(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

func (repo *userRepository) Get(id string) (*models.User, error) {
	var user models.User
	err := repo.database.First(&user, "id=?", id).Error
	if err != nil {
		return nil, &errors.FailedSelectError{EntityType: "User", Err: err}
	}

	return &user, nil
}

func (repo *userRepository) GetAll() ([]models.User, error) {
	var users []models.User
	err := repo.database.Find(&users).Error

	if err != nil {
		return nil, &errors.FailedSelectError{EntityType: "User", Err: err}
	}

	return users, nil
}

func (repo *userRepository) Save(user models.User) error {
	err := repo.database.Create(&user).Error
	if err != nil {
		return &errors.FailedInsertError{Entity: entityType, Err: err}
	}
	return nil
}
