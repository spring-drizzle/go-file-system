package controllers

import (
	"net/http"
	"src/errors"
	"src/requests"
	"src/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.IUserService
}

func (controller *UserController) setupUserRoutes(router *gin.RouterGroup) {
	userRoutesGroup := router.Group("/user")
	userRoutesGroup.GET("/Get/:id", controller.Get)
	userRoutesGroup.POST("/Save", controller.Save)
}

func (controller *UserController) Save(context *gin.Context) {
	request := requests.UserCreateRequest{}
	bindErr := context.ShouldBindJSON(&request)
	if bindErr != nil {
		context.JSON(http.StatusBadRequest, nil)
		panic(errors.BadRequestError{RequestType: "SaveUser", Err: bindErr})
	}

	saveErr := controller.UserService.Save(request)
	if saveErr != nil {
		panic(saveErr)
	}
	context.JSON(http.StatusOK, nil)
}

func (controller *UserController) Get(context *gin.Context) {
	id := context.Param("id")
	if id == "" {
		context.JSON(http.StatusBadRequest, nil)
		panic(errors.BadRequestError{RequestType: "GetUser", Err: &errors.InvalidParamError{ParamName: "id", ParamValue: id}})
	}
	user, err := controller.UserService.Get(id)

	if err == nil {
		panic(err)
	}
	context.Header("Content-Type", "application/json")
	context.JSON(http.StatusOK, user)
}
