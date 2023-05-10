package main

import (
	"go-file-system/src/controllers"
	"go-file-system/src/models"
	"go-file-system/src/providers"
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	buildApp(router)

	s := &http.Server{
		Addr:    "localhost:6969",
		Handler: router,
	}
	s.ListenAndServe()
}

func buildApp(router *gin.Engine) {
	validator := validator.New()
	db := providers.GetDbConnection()
	db.AutoMigrate(&models.User{})
	controllers.SetupRoutes(router, db, validator)
}
