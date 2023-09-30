package routes

import (
	"github.com/awbw/2040/controllers"
	"github.com/awbw/2040/middleware"
	"github.com/gin-gonic/gin"
)

func PressRoutes(r *gin.Engine) {
	pressRoutes := r.Group("/press")
	{
		pressRoutes.GET("/:id")

		pressRoutes.GET("/game/:id")

		pressRoutes.POST(
			"/create",
			middleware.RequireAuth,
			middleware.ValidatePlayer,
			controllers.Press.Create,
		)
	}
}
