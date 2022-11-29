package middlewares

import (
	auth "jwttuts/controllers/Auth"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenstring := context.GetHeader("Authorization")
		if tokenstring == "" {
			context.JSON(401, gin.H{"error": "no header found"})
			context.Abort()
			return
		}

		err := auth.ValidateToken(tokenstring)
		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		context.Next()
	}
}
