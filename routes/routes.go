package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	controllers "github.com/karnatakaPolls/controllers"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Routes(router *gin.Engine) {
	pollsRouter := router.Group("/polls")
	pollsRouter.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	pollsRouter.GET("/", welcome)
	pollsRouter.GET("/voters", controllers.GetAllVoters)
	pollsRouter.POST("/voter", controllers.CreateVoter)
	pollsRouter.GET("/voter/:voterId", controllers.GetSingleVoter)
	pollsRouter.PUT("/voter/:voterId", controllers.EditVoter)
	pollsRouter.DELETE("/voter/:voterId", controllers.DeleteVoter)
}
func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome To API",
	})
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  404,
		"message": "Route Not Found",
	})
}
