package middleware

import (
	"net/http"
	"strings"

	"github.com/gillgillint/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	s := context.Request.Header.Get("Authorization")
	token := strings.TrimPrefix(s, "Bearer ")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorization"})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorization"})
		return
	}

	context.Set("userId", userId)
	context.Next()
}
