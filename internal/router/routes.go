package router

import (
	"github.com/8123-molina/golang-api-access_users_routes/internal/controllers"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("api/v1")
	{
		v1.GET("/access", controllers.GetUsersAccessController)
	}
}
