package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xavimoreno548/BookStore/internal/services"
)

func Authorize() gin.HandlerFunc{
	return func(ctx *gin.Context){
		err := services.TokenValid(ctx)
		if err != nil {
			ctx.String(http.StatusUnauthorized, "Unauthorized")
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

