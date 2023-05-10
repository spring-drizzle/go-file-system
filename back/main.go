package main

import (
	"net/http"
	"src/controllers"
	"src/providers"
	"src/vendor/github.com/go-playground/validator/v10"

	"src/vendor/github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	buildApp(router)

	s := &http.Server{
		Addr:    ":6969",
		Handler: router,
	}
	s.ListenAndServe()
}

func buildApp(router *gin.Engine) {
	validator := validator.New()
	db := providers.GetDbConnection()
	controllers.SetupRoutes(router, db, validator)
}
