package routers

import (
	"minigram-app-backend/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", controllers.AuthRegister)
		authGroup.POST("/login", controllers.AuthLogin)
	}

	return router
}
