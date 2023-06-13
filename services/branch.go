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

type BranchService interface {
	GetAll(ctx *gin.Context, src *[]domain.Branch) (err error)
	GetDetail(ctx *gin.Context, src *domain.Branch, id int64) (err error)
	Create(ctx *gin.Context, src *domain.Branch) (err error)
	Update(ctx *gin.Context, src *domain.Branch) (err error)
	Delete(ctx *gin.Context, src *domain.Branch) (err error)
}

type branchService struct {
	branchRepo repository.BranchRepository
}

func NewBranchService(branchRepo repository.BranchRepository) BranchService {
	return &branchService{
		branchRepo: branchRepo,
	}
}

func (srv *branchService) GetAll(ctx *gin.Context, src *[]domain.Branch) (err error) {
	if err = srv.branchRepo.GetAll(ctx, src); err != nil {
		return
	}

	return
}

func (srv *branchService) GetDetail(ctx *gin.Context, src *domain.Branch, id int64) (err error) {
	if err = srv.branchRepo.GetDetail(ctx, src, id); err != nil {
		return
	}

	return
}

func (srv *branchService) Create(ctx *gin.Context, src *domain.Branch) (err error) {
	if err = db.GetConn().RunInTx(ctx, &sql.TxOptions{}, func(c context.Context, tx bun.Tx) error {
		if err = srv.branchRepo.Create(ctx, &tx, src); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return
	}

	return
}

func (srv *branchService) Update(ctx *gin.Context, src *domain.Branch) (err error) {
	if err = db.GetConn().RunInTx(ctx, &sql.TxOptions{}, func(c context.Context, tx bun.Tx) error {
		if err = srv.branchRepo.Update(ctx, &tx, src, src.ID); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return
	}

	return
}

func (srv *branchService) Delete(ctx *gin.Context, src *domain.Branch) (err error) {
	if err = db.GetConn().RunInTx(ctx, &sql.TxOptions{}, func(c context.Context, tx bun.Tx) error {
		if err = srv.branchRepo.Delete(ctx, &tx, src, src.ID); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return
	}

	return
}
