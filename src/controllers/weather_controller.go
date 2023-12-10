package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nikhil133/duplo/src/servicefactory"
	"github.com/nikhil133/duplo/src/serviceinterface"
	utils "github.com/nikhil133/duplo/utils/request"
)

type Api struct {
	serv serviceinterface.WeatherServiceInterface
}

// GetForcast godoc
// @Summary Get weather forecast based on latitude and longitude
// @Description Retrieve weather forecast data for a given latitude and longitude.
// @ID get-forecast
// @Produce json
// @Param latitude query number true "Latitude coordinate"
// @Param longitude query number true "Longitude coordinate"
// @Success 200 {object} service.Forecast "Successfully retrieved forecast data"
// @Failure 400 {string} string "Bad Request. Invalid latitude or longitude."
// @Failure 500 {string} string "Internal Server Error. Failed to fetch forecast data."
// @Router /forecast [get]
func GetForcast(c *gin.Context) {
	var (
		req utils.Request
		err error
	)
	req.Latitude, err = strconv.ParseFloat(c.Query("latitude"), 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	req.Longitude, err = strconv.ParseFloat(c.Query("longitude"), 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ws := Api{
		serv: servicefactory.NewWeatherService(),
	}
	weatherService(c, req, ws.serv)

}

func weatherService(c *gin.Context, req utils.Request, ws serviceinterface.WeatherServiceInterface) {
	data, err := ws.GetForcast(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, data)

}

// GetCoordinates godoc
// @Summary Get weather service coordinates
// @Description Retrieve the current weather service coordinates.
// @ID get-coordinates
// @Produce json
// @Success 200 {array} service.Coordinate "Successfully retrieved coordinates"
// @Failure 500 {string} string "Internal Server Error. Failed to fetch coordinates."
// @Router /coordinate/history [get]
func GetCoordinates(c *gin.Context) {
	ws := Api{
		serv: servicefactory.NewWeatherService(),
	}

	GetCoordinateHelper(c, ws.serv)
}

func GetCoordinateHelper(c *gin.Context, ws serviceinterface.WeatherServiceInterface) {
	data, err := ws.GetCoordinates()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, data)
}

// DeleteCoordinates godoc
// @Summary Delete weather service coordinates by ID
// @Description Delete weather service coordinates based on the provided ID.
// @ID delete-coordinates
// @Produce json
// @Param id query integer true "ID of the coordinates to be deleted"
// @Success 200 {string} string "Record deleted successfully"
// @Failure 400 {string} string "Bad Request. Invalid ID format."
// @Failure 500 {string} string "Internal Server Error. Failed to delete coordinates."
// @Router /coordinate [delete]
func DeleteCoordinate(c *gin.Context) {
	reqId := c.Query("id")
	id, err := strconv.Atoi(reqId)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return

	}
	ws := Api{
		serv: servicefactory.NewWeatherService(),
	}
	deleteCoordinateHelper(c, uint(id), ws.serv)

}

func deleteCoordinateHelper(c *gin.Context, id uint, ws serviceinterface.WeatherServiceInterface) {
	err := ws.DeleteCoordinates(uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "record deleted successfully")

}
