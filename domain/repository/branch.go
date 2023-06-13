package repository

import (
	"fauzanbintang/alfath/db"
	"fauzanbintang/alfath/domain"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type BranchRepository interface {
	GetAll(ctx *gin.Context, src *[]domain.Branch) (err error)
	GetDetail(ctx *gin.Context, src *domain.Branch, id int64) (err error)
	Create(ctx *gin.Context, tx *bun.Tx, src *domain.Branch) (err error)
	Update(ctx *gin.Context, tx *bun.Tx, src *domain.Branch, id int64) (err error)
	Delete(ctx *gin.Context, tx *bun.Tx, src *domain.Branch, id int64) (err error)
}

type branchRepository struct {
	db *bun.DB
}

func NewBranchRepository() BranchRepository {
	return &branchRepository{
		db: db.GetConn(),
	}
}

func (r *branchRepository) GetAll(ctx *gin.Context, src *[]domain.Branch) (err error) {
	err = r.db.NewSelect().
		Model(src).
		Scan(ctx)
	if err != nil {
		return err
	}

	return
}

func (r *branchRepository) GetDetail(ctx *gin.Context, src *domain.Branch, id int64) (err error) {
	err = r.db.NewSelect().
		Model(src).
		Where("id = ?", id).
		Scan(ctx)
	if err != nil {
		return err
	}

	return
}

func (r *branchRepository) Create(ctx *gin.Context, tx *bun.Tx, src *domain.Branch) (err error) {
	_, err = tx.NewInsert().
		Model(src).
		Exec(ctx)
	if err != nil {
		return err
	}

	return
}

func (r *branchRepository) Update(ctx *gin.Context, tx *bun.Tx, src *domain.Branch, id int64) (err error) {
	_, err = tx.NewUpdate().
		Model(src).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return err
	}

	return
}

func (r *branchRepository) Delete(ctx *gin.Context, tx *bun.Tx, src *domain.Branch, id int64) (err error) {
	_, err = tx.NewDelete().
		Model(src).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return err
	}

	return
}
