package providers

import (
	"go-file-system/src/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDbConnection() *gorm.DB {
	dialector := postgres.Open(config.GetConnectionString())

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
