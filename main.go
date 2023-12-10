package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nikhil133/duplo/config"
	_ "github.com/nikhil133/duplo/docs"
	"github.com/nikhil133/duplo/routes"
	"github.com/nikhil133/duplo/src/migration"
	"github.com/nikhil133/duplo/src/repositoryfactory"
)

// @title     Duplo Weather Forecast
func main() {

	config.LoadConfig()
	repositoryfactory.NewWeatherRepository()

	//database.GetInstancemysql()
	router := gin.Default()

	migration.Migration()
	routes.SetupRoutes(router)

}
