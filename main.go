package main

import (
	"github.com/awbw/2040/db"
	"github.com/awbw/2040/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	db.ConnectDatabase()
}

func main() {
	router := gin.New()

	routes.GameRoutes(router)
}
