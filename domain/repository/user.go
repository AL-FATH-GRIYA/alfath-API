package repository

import (
	"fauzanbintang/alfath/db"
	"fauzanbintang/alfath/domain"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type UserRepository interface {
	GetAll(ctx *gin.Context, src *[]domain.User) (err error)
	GetDetail(ctx *gin.Context, src *domain.User, id int64) (err error)
	Register(ctx *gin.Context, tx *bun.Tx, src *domain.User) (err error)
	GetByEmail(ctx *gin.Context, src *domain.User, email string) (err error)
}

type userRepository struct {
	db *bun.DB
}

func NewUserRepository() UserRepository {
	return &userRepository{
		db: db.GetConn(),
	}
}

func (r *userRepository) GetAll(ctx *gin.Context, src *[]domain.User) (err error) {
	err = r.db.NewSelect().
		Model(src).
		Scan(ctx)
	if err != nil {
		return err
	}

	return
}

func (r *userRepository) GetDetail(ctx *gin.Context, src *domain.User, id int64) (err error) {
	err = r.db.NewSelect().
		Model(src).
		Where("id = ?", id).
		Scan(ctx)
	if err != nil {
		return err
	}

	return
}

func (r *userRepository) Register(ctx *gin.Context, tx *bun.Tx, src *domain.User) (err error) {
	_, err = tx.NewInsert().
		Model(src).
		Exec(ctx)
	if err != nil {
		return err
	}

	return
}

func (r *userRepository) GetByEmail(ctx *gin.Context, src *domain.User, email string) (err error) {
	err = r.db.NewSelect().
		Model(src).
		Where("email = ?", email).
		Scan(ctx)
	if err != nil {
		return err
	}

	return
}
