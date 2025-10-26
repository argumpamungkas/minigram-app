package routers

import (
	"minigram-app-backend/controllers/auth_controller"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", auth_controller.AuthRegister)
		authGroup.POST("/login", auth_controller.AuthLogin)
	}

	return router
}
