package middlewares

import (
	"fauzanbintang/alfath/db"
	"fauzanbintang/alfath/domain"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(ctx *gin.Context) {
	tokenStr, err := ctx.Cookie("Auth")
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("AllYourBase"), nil
	})
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

		var user domain.User
		err = db.GetConn().NewSelect().
			Model(&user).
			Where("id = ?", claims["sub"]).
			Scan(ctx)
		if err != nil {
			fmt.Println("error", err)
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

		ctx.Set("user", user)
	} else {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

	ctx.Next()
}
