package services

import (
	"context"
	"database/sql"
	"fauzanbintang/alfath/db"
	"fauzanbintang/alfath/domain"
	"fauzanbintang/alfath/domain/repository"
	"fauzanbintang/alfath/dto/responses"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/uptrace/bun"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetAll(ctx *gin.Context, src *[]domain.User) (res []responses.UserDetail, err error)
	GetDetail(ctx *gin.Context, src *domain.User, id int64) (res responses.UserDetail, err error)
	Register(ctx *gin.Context, src *domain.User) (res responses.UserDetail, err error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (srv *userService) GetAll(ctx *gin.Context, src *[]domain.User) (res []responses.UserDetail, err error) {
	if err = srv.userRepo.GetAll(ctx, src); err != nil {
		return
	}

	for _, user := range *src {
		userDetail := responses.UserDetail{
			ID:        user.ID,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}

		res = append(res, userDetail)
	}

	return
}

func (srv *userService) GetDetail(ctx *gin.Context, src *domain.User, id int64) (res responses.UserDetail, err error) {
	if err = srv.userRepo.GetDetail(ctx, src, id); err != nil {
		return
	}

	if err = copier.Copy(&res, src); err != nil {
		return
	}

	return
}

func (srv *userService) Register(ctx *gin.Context, src *domain.User) (res responses.UserDetail, err error) {
	if err = db.GetConn().RunInTx(ctx, &sql.TxOptions{}, func(c context.Context, tx bun.Tx) error {
		hash, err := bcrypt.GenerateFromPassword([]byte(src.Password), 10)
		if err != nil {
			return err
		}

		src.Password = string(hash)
		src.Role = "admin"

		if err = srv.userRepo.Register(ctx, &tx, src); err != nil {
			return err
		}

		if err = copier.Copy(&res, src); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return
	}

	return
}
