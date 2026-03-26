package routes

import (
	"study-api/services"

	"github.com/gin-gonic/gin"
)

func RegisterUsersRoutes(router *gin.Engine) {
	r := router.Group("/users")
	{
		r.GET("/", services.GetUsers)
	}
}
