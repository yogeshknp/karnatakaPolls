package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	controllers "github.com/karnatakaPolls/controllers"
)

func Routes(router *gin.Engine) {
	router.GET("/", welcome)
	router.GET("/todos", controllers.GetAllVoters)
	router.POST("/todo", controllers.CreateVoter)
	router.GET("/todo/:todoId", controllers.GetSingleVoter)
	router.PUT("/todo/:todoId", controllers.EditVoter)
	router.DELETE("/todo/:todoId", controllers.DeleteVoter)
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
