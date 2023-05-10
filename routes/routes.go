package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	controllers "github.com/karnatakaPolls/controllers"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Routes(router *gin.Engine) {
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.NoRoute(notFound)

	pollsRouter := router.Group("/polls")
	pollsRouter.GET("/", welcome)

	pollsRouter.POST("/voter", controllers.CreateVoter)
	pollsRouter.GET("/voter/:AadhaarID", controllers.GetSingleVoter)
	pollsRouter.GET("/voters", controllers.GetAllVoters)
	pollsRouter.PUT("/voter/:AadhaarID", controllers.EditVoter)
	pollsRouter.DELETE("/voter/:AadhaarID", controllers.DeleteVoter)

	pollsRouter.POST("/candidate", controllers.CreateCandidate)
	pollsRouter.GET("/candidate/:AadhaarID", controllers.GetSingleCandidate)
	pollsRouter.GET("/candidates", controllers.GetAllCandidate)
	pollsRouter.PUT("/candidate/:AadhaarID", controllers.EditCandidate)
	pollsRouter.DELETE("/candidate/:AadhaarID", controllers.DeleteCandidate)

	pollsRouter.GET("/constituency/:name", controllers.GetSingleConstituency)
	pollsRouter.POST("/constituency", controllers.CreateConstituency)
	pollsRouter.GET("/constituencies", controllers.GetAllConstituency)
	pollsRouter.PUT("/constituency/:name", controllers.EditConstituency)
	pollsRouter.DELETE("/constituency/:name", controllers.DeleteConstituency)

	pollsRouter.PUT("/vote/:VoterAadhaarID/:CanAadhaarID", controllers.Vote)

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
