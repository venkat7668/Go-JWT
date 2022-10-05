package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/venkat7668/Go-JWT/auth"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if len(tokenString) == 0 {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}
		if err := auth.ValidateToken(tokenString); err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		context.Next()
	}
}
