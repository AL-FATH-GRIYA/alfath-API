package services

import (
	"context"
	"database/sql"
	"fauzanbintang/alfath/db"
	"fauzanbintang/alfath/domain"
	"fauzanbintang/alfath/domain/repository"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type BrandService interface {
	GetAll(ctx *gin.Context, src *[]domain.Brand) (err error)
	GetDetail(ctx *gin.Context, src *domain.Brand, id int64) (err error)
	Create(ctx *gin.Context, src *domain.Brand) (err error)
	Update(ctx *gin.Context, src *domain.Brand) (err error)
	Delete(ctx *gin.Context, src *domain.Brand) (err error)
}

type brandService struct {
	brandRepo repository.BrandRepository
}

func NewBrandService(brandRepo repository.BrandRepository) BrandService {
	return &brandService{
		brandRepo: brandRepo,
	}
}

func (srv *brandService) GetAll(ctx *gin.Context, src *[]domain.Brand) (err error) {
	if err = srv.brandRepo.GetAll(ctx, src); err != nil {
		return
	}

	return
}

func (srv *brandService) GetDetail(ctx *gin.Context, src *domain.Brand, id int64) (err error) {
	if err = srv.brandRepo.GetDetail(ctx, src, id); err != nil {
		return
	}

	return
}

func (srv *brandService) Create(ctx *gin.Context, src *domain.Brand) (err error) {
	if err = db.GetConn().RunInTx(ctx, &sql.TxOptions{}, func(c context.Context, tx bun.Tx) error {
		if err = srv.brandRepo.Create(ctx, &tx, src); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return
	}

	return
}

func (srv *brandService) Update(ctx *gin.Context, src *domain.Brand) (err error) {
	if err = db.GetConn().RunInTx(ctx, &sql.TxOptions{}, func(c context.Context, tx bun.Tx) error {
		if err = srv.brandRepo.Update(ctx, &tx, src, src.ID); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return
	}

	return
}

func (srv *brandService) Delete(ctx *gin.Context, src *domain.Brand) (err error) {
	if err = db.GetConn().RunInTx(ctx, &sql.TxOptions{}, func(c context.Context, tx bun.Tx) error {
		if err = srv.brandRepo.Delete(ctx, &tx, src, src.ID); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return
	}

	return
}
