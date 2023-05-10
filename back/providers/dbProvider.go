package providers

import (
	"src/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDbConnection() *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.GetConnectionString()), &gorm.Config{})
	if err == nil {
		panic(err)
	}

	return db
}
