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

type BranchController interface {
	GetAll(ctx *gin.Context)
	GetDetail(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type branchController struct {
	branchSrv services.BranchService
}

func NewBranchController(branchSrv services.BranchService) BranchController {
	return &branchController{
		branchSrv: branchSrv,
	}
}

func (ctl *branchController) GetAll(ctx *gin.Context) {
	var branchs []domain.Branch
	err := ctl.branchSrv.GetAll(ctx, &branchs)
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
		"data":    branchs,
	})
}

func (ctl *branchController) GetDetail(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 1<<6)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
			"error":   err,
		})
		return
	}

	var branch domain.Branch
	err = ctl.branchSrv.GetDetail(ctx, &branch, id)
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
		"data":    branch,
	})
}

func (ctl *branchController) Create(ctx *gin.Context) {
	var req requests.BranchForm

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "error bind json",
			"error":   err,
		})

		return
	}

	var branch domain.Branch
	if err = copier.Copy(&branch, req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
			"error":   err,
		})
		return
	}

	err = ctl.branchSrv.Create(ctx, &branch)
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
		"data":    branch,
	})
}

func (ctl *branchController) Update(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 1<<6)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
			"error":   err,
		})
		return
	}

	var req requests.BranchForm

	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "error bind json",
			"error":   err,
		})

		return
	}

	var branch domain.Branch
	if err = copier.Copy(&branch, req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
			"error":   err,
		})
		return
	}

	branch.ID = id

	err = ctl.branchSrv.Update(ctx, &branch)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Successfully update data",
		"data":    branch,
	})
}

func (ctl *branchController) Delete(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 1<<6)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
			"error":   err,
		})
		return
	}

	var branch domain.Branch
	branch.ID = id
	err = ctl.branchSrv.Delete(ctx, &branch)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Successfully delete data",
	})
}
