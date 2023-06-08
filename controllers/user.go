package controllers

import (
	"fauzanbintang/alfath/domain"
	"fauzanbintang/alfath/dto/requests"
	"fauzanbintang/alfath/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type UserController interface {
	GetAll(ctx *gin.Context)
	GetDetail(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type userController struct {
	userSrv services.UserService
}

func NewUserController(userSrv services.UserService) UserController {
	return &userController{
		userSrv: userSrv,
	}
}

func (ctl *userController) GetAll(ctx *gin.Context) {
	var users []domain.User

	res, err := ctl.userSrv.GetAll(ctx, &users)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Successfully load data",
		"data":    res,
	})
}

func (ctl *userController) GetDetail(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 1<<6)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
			"error":   err,
		})
		return
	}

	var user domain.User
	res, err := ctl.userSrv.GetDetail(ctx, &user, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Successfully load data",
		"data":    res,
	})
}

func (ctl *userController) Register(ctx *gin.Context) {
	var req requests.RegisterUser

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "error bind json",
			"error":   err,
		})

		return
	}

	var user domain.User
	if err = copier.Copy(&user, req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
			"error":   err,
		})
		return
	}

	res, err := ctl.userSrv.Register(ctx, &user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"code":    http.StatusCreated,
		"message": "Successfully create data",
		"data":    res,
	})
}
