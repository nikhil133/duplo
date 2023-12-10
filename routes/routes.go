package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nikhil133/duplo/src/controllers"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(router *gin.Engine) {

	api := router.Group("/")

	api.GET("forecast", controllers.GetForcast)
	api.GET("coordinate/history", controllers.GetCoordinates)
	api.DELETE("coordinate", controllers.DeleteCoordinate)
	api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(viper.GetString("server.port"))
}
