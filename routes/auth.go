package routes

import (
	"study-api/auth"

	"github.com/gin-gonic/gin"
)

func RegisterLoginRoutes(router *gin.Engine) {
	r := router.Group("/login")

	{
		r.POST("/", auth.Login)
	}
}
