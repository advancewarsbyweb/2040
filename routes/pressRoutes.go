package routes

import (
	"github.com/awbw/2040/controllers"
	"github.com/awbw/2040/middleware"
	"github.com/gin-gonic/gin"
)

func PressRoutes(r *gin.Engine) {
	pressRoutes := r.Group("/press").Use(middleware.RequireAuth, middleware.ValidatePlayer)
	{
		pressRoutes.GET("/:id")

		pressRoutes.GET("/player/:playerId", controllers.Press.GetAll)

		pressRoutes.POST("/create", controllers.Press.Create)
	}
}
