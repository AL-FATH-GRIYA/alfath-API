package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Group("/users")
	{
		r.GET("", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "successfully get users",
			})
		})
	}

	r.Run(":3000")
}
