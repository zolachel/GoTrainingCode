package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth(context *gin.Context) {
	log.Println("start middleware")

	authKey := context.GetHeader("Authorization")

	if authKey != "Bearer token123" {
		context.JSON(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
		context.Abort()
		return
	}

	context.Next()

	log.Println("end middleware")
}
