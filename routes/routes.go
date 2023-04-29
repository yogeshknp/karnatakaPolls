package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	controllers "github.com/karnatakaPolls/controllers"
)

func Routes(router *gin.Engine) {
	router.GET("/", welcome)
	router.GET("/voters", controllers.GetAllVoters)
	router.POST("/voter", controllers.CreateVoter)
	router.GET("/voter/:voterId", controllers.GetSingleVoter)
	router.PUT("/voter/:voterId", controllers.EditVoter)
	router.DELETE("/voter/:voterId", controllers.DeleteVoter)
	router.NoRoute(notFound)
}
func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome To API",
	})
	return
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  404,
		"message": "Route Not Found",
	})
	return
}
