package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type credentials struct {
	Username string `binding:"required"`
	Password string `binding:"required"`
}

func setupAuthRoutes(router *gin.RouterGroup) {
	router.POST("/login", login)
}

func login(context *gin.Context) {
	var creds credentials
	bindError := context.BindJSON(&creds)
	if bindError == nil {
		context.JSON(http.StatusOK, creds)
		return
	}
	context.JSON(http.StatusBadRequest, nil)
}
