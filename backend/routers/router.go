package routers

import (
	"minigram-app-backend/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	// router.Group("/auth", func(ctx *gin.Context) {
	router.POST("/auth/register", controllers.AuthRegister)
	router.POST("/auth/login", controllers.AuthLogin)
	// })

	return router
}
