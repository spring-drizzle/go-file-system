package controllers

import (
	"go-file-system/src/repositories"
	"go-file-system/src/services"

	"github.com/go-playground/validator/v10"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB, validator *validator.Validate) {
	api := router.Group("/api")
	setupUserController(api, db, validator)
}

func setupUserController(router *gin.RouterGroup, db *gorm.DB, validator *validator.Validate) {
	userRepo := repositories.NewUserRepo(db)
	userService := services.NewUserService(userRepo, validator)
	userController := &UserController{UserService: userService}

	userController.setupUserRoutes(router)
}
