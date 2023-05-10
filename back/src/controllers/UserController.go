package controllers

import (
	"go-file-system/src/errors"
	"go-file-system/src/requests"
	"go-file-system/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.IUserService
}

func (controller *UserController) setupUserRoutes(router *gin.RouterGroup) {
	userRoutesGroup := router.Group("/user")
	userRoutesGroup.GET("/get/:id", controller.Get)
	userRoutesGroup.GET("/getAll", controller.GetAll)
	userRoutesGroup.POST("/save", controller.Save)
}

func (controller *UserController) Get(context *gin.Context) {
	id := context.Param("id")
	if id == "" {
		context.JSON(http.StatusBadRequest, errors.BadRequestError{RequestType: "GetUser", Err: &errors.InvalidParamError{ParamName: "id", ParamValue: id}})
		return
	}
	user, err := controller.UserService.Get(id)

	if err != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}
	context.Header("Content-Type", "application/json")
	context.JSON(http.StatusOK, user)
}

func (controller *UserController) GetAll(context *gin.Context) {
	users, err := controller.UserService.GetAll()

	if err != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}

	context.JSON(http.StatusOK, users)
}

func (controller *UserController) Save(context *gin.Context) {
	request := requests.UserCreateRequest{}
	bindErr := context.ShouldBindJSON(&request)
	if bindErr != nil {
		context.JSON(http.StatusBadRequest, errors.BadRequestError{RequestType: "SaveUser", Err: bindErr})
		return
	}

	saveErr := controller.UserService.Save(request)
	if saveErr != nil {
		context.JSON(http.StatusInternalServerError, saveErr)
		return
	}
	context.JSON(http.StatusOK, "User inserted")
}
