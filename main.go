package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/awbw/2040/db"
	"github.com/awbw/2040/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var router *gin.Engine

func init() {
	godotenv.Load()
	db.NewDatabase(os.Getenv("TEST_DSN"))
}

func main() {
	router = gin.New()
	routes.GameRoutes(router)
	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), router)
}
