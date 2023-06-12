package main

import (
	"fauzanbintang/alfath/controllers"
	"fauzanbintang/alfath/db"
	"fauzanbintang/alfath/domain/repository"
	"fauzanbintang/alfath/middlewares"
	"fauzanbintang/alfath/services"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	dbInstance := db.InitDB()
	defer dbInstance.Close()

	rp := repository.InitRepositoryInstance()

	userSrv := services.NewUserService(rp.User)
	userCtl := controllers.NewUserController(userSrv)

	brandSrv := services.NewBrandService(rp.Brand)
	brandCtl := controllers.NewBrandController(brandSrv)

	users := r.Group("/users")
	{
		users.GET("", userCtl.GetAll)
		users.POST("register", middlewares.RequireAuth, userCtl.Register)
		users.POST("login", userCtl.Login)
		users.GET(":id", userCtl.GetDetail)
	}

	brands := r.Group("/brands")
	{
		brands.GET("", middlewares.RequireAuth, brandCtl.GetAll)
		brands.POST("", middlewares.RequireAuth, brandCtl.Create)
		brands.GET(":id", middlewares.RequireAuth, brandCtl.GetDetail)
		brands.PUT(":id", middlewares.RequireAuth, brandCtl.Update)
		brands.DELETE(":id", middlewares.RequireAuth, brandCtl.Delete)
	}

	r.Run(":3000")
}
