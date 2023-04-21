package main

import (
	"github.com/awbw/2040/api/routes"
	"github.com/awbw/2040/conf"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	conf.ConnectDatabase()
	conf.ApplyMigrations()
}

func main() {
	router := gin.New()

	routes.GameRoutes(router)
}
