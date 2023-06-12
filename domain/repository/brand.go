package repository

import (
	"fauzanbintang/alfath/db"
	"fauzanbintang/alfath/domain"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type BrandRepository interface {
	GetAll(ctx *gin.Context, src *[]domain.Brand) (err error)
	GetDetail(ctx *gin.Context, src *domain.Brand, id int64) (err error)
	Create(ctx *gin.Context, tx *bun.Tx, src *domain.Brand) (err error)
	Update(ctx *gin.Context, tx *bun.Tx, src *domain.Brand, id int64) (err error)
	Delete(ctx *gin.Context, tx *bun.Tx, src *domain.Brand, id int64) (err error)
}

type brandRepository struct {
	db *bun.DB
}

func NewBrandRepository() BrandRepository {
	return &brandRepository{
		db: db.GetConn(),
	}
}

func (r *brandRepository) GetAll(ctx *gin.Context, src *[]domain.Brand) (err error) {
	err = r.db.NewSelect().
		Model(src).
		Scan(ctx)
	if err != nil {
		return err
	}

	return
}

func (r *brandRepository) GetDetail(ctx *gin.Context, src *domain.Brand, id int64) (err error) {
	err = r.db.NewSelect().
		Model(src).
		Where("id = ?", id).
		Scan(ctx)
	if err != nil {
		return err
	}

	return
}

func (r *brandRepository) Create(ctx *gin.Context, tx *bun.Tx, src *domain.Brand) (err error) {
	_, err = tx.NewInsert().
		Model(src).
		Exec(ctx)
	if err != nil {
		return err
	}

	return
}

func (r *brandRepository) Update(ctx *gin.Context, tx *bun.Tx, src *domain.Brand, id int64) (err error) {
	_, err = tx.NewUpdate().
		Model(src).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return err
	}

	return
}

func (r *brandRepository) Delete(ctx *gin.Context, tx *bun.Tx, src *domain.Brand, id int64) (err error) {
	_, err = tx.NewDelete().
		Model(src).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return err
	}

	return
}
