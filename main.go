package main

import (
	"log"

	"github.com/gin-gonic/gin"
	config "github.com/karnatakaPolls/config"
	_ "github.com/karnatakaPolls/docs"
	routes "github.com/karnatakaPolls/routes"
)

// @title  Karnataka Polls Assignment
// @version	1.0
// @description A polling service API in Go using Gin framework
// @host 	localhost:4747
// @BasePath /polls

func main() {
	config.Connect()
	// Init Router
	router := gin.Default()
	// Route Handlers / Endpoints
	routes.Routes(router)
	log.Fatal(router.Run(":4747"))
}
