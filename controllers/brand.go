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

type BrandController interface {
	GetAll(ctx *gin.Context)
	GetDetail(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type brandController struct {
	brandSrv services.BrandService
}

func NewBrandController(brandSrv services.BrandService) BrandController {
	return &brandController{
		brandSrv: brandSrv,
	}
}

func (ctl *brandController) GetAll(ctx *gin.Context) {
	var brands []domain.Brand
	err := ctl.brandSrv.GetAll(ctx, &brands)
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
		"data":    brands,
	})
}

func (ctl *brandController) GetDetail(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 1<<6)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
			"error":   err,
		})
		return
	}

	var brand domain.Brand
	err = ctl.brandSrv.GetDetail(ctx, &brand, id)
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
		"data":    brand,
	})
}

func (ctl *brandController) Create(ctx *gin.Context) {
	var req requests.BrandForm

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "error bind json",
			"error":   err,
		})

		return
	}

	var brand domain.Brand
	if err = copier.Copy(&brand, req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
			"error":   err,
		})
		return
	}

	err = ctl.brandSrv.Create(ctx, &brand)
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
		"data":    brand,
	})
}

func (ctl *brandController) Update(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 1<<6)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
			"error":   err,
		})
		return
	}

	var req requests.BrandForm

	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "error bind json",
			"error":   err,
		})

		return
	}

	var brand domain.Brand
	if err = copier.Copy(&brand, req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
			"error":   err,
		})
		return
	}

	brand.ID = id

	err = ctl.brandSrv.Update(ctx, &brand)
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
		"data":    brand,
	})
}

func (ctl *brandController) Delete(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 1<<6)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
			"error":   err,
		})
		return
	}

	var brand domain.Brand
	brand.ID = id
	err = ctl.brandSrv.Delete(ctx, &brand)
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
