package routers

import (
	"minigram-app-backend/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/auth/register", controllers.AuthRegister)
	router.GET("/auth/user", controllers.GetAllUser)

	return router
}
