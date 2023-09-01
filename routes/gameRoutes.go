package routes

import (
	"github.com/awbw/2040/controllers"
	"github.com/gin-gonic/gin"
)

func GameRoutes(r *gin.Engine) {
	// Game
	r.GET("/games/:id", controllers.Game.Get)

	// Players
	//r.GET("/games/:id/players")
	//r.GET("/games/:id/players/:player_id")

	// Units
	//r.GET("/games/:id/players/:player_id/units")
	//r.GET("/games/:id/players/:player_id/units/:unit_id")

}
