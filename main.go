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

	branchSrv := services.NewBranchService(rp.Branch)
	branchCtl := controllers.NewBranchController(branchSrv)

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

	branches := r.Group("/branches")
	{
		branches.GET("", middlewares.RequireAuth, branchCtl.GetAll)
		branches.POST("", middlewares.RequireAuth, branchCtl.Create)
		branches.GET(":id", middlewares.RequireAuth, branchCtl.GetDetail)
		branches.PUT(":id", middlewares.RequireAuth, branchCtl.Update)
		branches.DELETE(":id", middlewares.RequireAuth, branchCtl.Delete)
	}

	r.Run(":3000")
}
