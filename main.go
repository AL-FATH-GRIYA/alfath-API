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

	users := r.Group("/users")
	{
		users.GET("", userCtl.GetAll)
		users.POST("register", middlewares.RequireAuth, userCtl.Register)
		users.POST("login", userCtl.Login)
		users.GET(":id", userCtl.GetDetail)
	}

	r.Run(":3000")
}
