package routes

import "github.com/gin-gonic/gin"

func unitRoutes(r *gin.Engine) {
	r.GET("/units/:id")
}
