package middlewares

import (
	"net/http"

	"github.com/frevent/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized action."})
		return
	}

	verifiedToken, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized action."})
		return
	}

	context.Set("token", verifiedToken)
	context.Next()
}
