package controllers

import (
	"database/sql"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"github.com/nikhil133/duplo/mocks"
	"github.com/nikhil133/duplo/src/service"
	"github.com/nikhil133/duplo/utils/constant"
	utils "github.com/nikhil133/duplo/utils/request"
)

func TestGetForcast(t *testing.T) {
	router := gin.Default()
	api := router.Group("/")

	api.GET("forecast", GetForcast)
	req, _ := http.NewRequest("GET", "/forecast?latitude=4.2&longitude=3.1", nil)

	constant.Test = true
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// GORM

	// Create mocks
	m := mocks.NewMockWeatherServiceInterface(ctrl)

	request := utils.Request{
		Longitude: 3.54,
		Latitude:  4.2,
	}
	f := &service.Forecast{
		Status: service.Status{
			Error:   false,
			Message: "success",
			Code:    http.StatusOK,
		},
	}
	a := Api{
		serv: m,
	}

	// Set up expectations for WeatherService
	m.EXPECT().GetForcast(request).Return(f, nil).AnyTimes()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	weatherService(c, request, a.serv)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetForcastError(t *testing.T) {
	router := gin.Default()
	api := router.Group("/")

	api.GET("forecast", GetForcast)

	req, _ := http.NewRequest("GET", "/forecast?latitude=4.2&longitude=3.1", nil)

	constant.Test = true
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// GORM

	// Create mocks
	m := mocks.NewMockWeatherServiceInterface(ctrl)

	request := utils.Request{
		Longitude: 3.54,
	}
	mockError := errors.New("mock error")
	a := Api{
		serv: m,
	}
	// Set up expectations for WeatherService
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	m.EXPECT().GetForcast(request).Return(nil, mockError)
	weatherService(c, request, a.serv)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestGetForcastValidationError(t *testing.T) {
	router := gin.Default()
	api := router.Group("/")

	api.GET("forecast", GetForcast)

	req, _ := http.NewRequest("GET", "/forecast?latitude=4.2&longitude=", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	req, _ = http.NewRequest("GET", "/forecast?latitude=&longitude=3.1", nil)
	w = httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestDeleteCoordinate(t *testing.T) {
	router := gin.Default()
	api := router.Group("/")

	api.DELETE("coordinate", DeleteCoordinate)

	req, _ := http.NewRequest("DELETE", "/coordinate?id=1", nil)
	w := httptest.NewRecorder()

	constant.Test = true
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWeatherServiceInterface(ctrl)

	a := Api{
		serv: m,
	}
	id := uint(1)
	m.EXPECT().DeleteCoordinates(id).Return(nil)
	a.serv.DeleteCoordinates(id)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteCoordinateError(t *testing.T) {
	router := gin.Default()
	api := router.Group("/")

	api.DELETE("coordinate", DeleteCoordinate)

	req, _ := http.NewRequest("DELETE", "/coordinate?id=0", nil)
	w := httptest.NewRecorder()

	constant.Test = true
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWeatherServiceInterface(ctrl)

	a := Api{
		serv: m,
	}
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	id := uint(0)
	err := sql.ErrNoRows
	m.EXPECT().DeleteCoordinates(id).Return(err)
	deleteCoordinateHelper(c, id, a.serv)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestDeleteCoordinateValidationError(t *testing.T) {
	router := gin.Default()
	api := router.Group("/")

	api.DELETE("coordinate", DeleteCoordinate)

	req, _ := http.NewRequest("DELETE", "/coordinate?id=", nil)
	w := httptest.NewRecorder()

	constant.Test = true
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetCoordinates(t *testing.T) {
	router := gin.Default()
	api := router.Group("/")

	api.GET("coordinate/history", GetCoordinates)

	req, _ := http.NewRequest("GET", "/coordinate/history", nil)
	w := httptest.NewRecorder()

	constant.Test = true
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWeatherServiceInterface(ctrl)

	a := Api{
		serv: m,
	}
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	cord := []service.Coordinate{}
	m.EXPECT().GetCoordinates().Return(cord, nil)
	GetCoordinateHelper(c, a.serv)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetCoordinatesError(t *testing.T) {
	router := gin.Default()
	api := router.Group("/")

	api.GET("coordinate/history", GetCoordinates)

	req, _ := http.NewRequest("GET", "/coordinate/history", nil)
	w := httptest.NewRecorder()

	constant.Test = true
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWeatherServiceInterface(ctrl)

	a := Api{
		serv: m,
	}
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	mockError := sql.ErrNoRows
	m.EXPECT().GetCoordinates().Return(nil, mockError)
	GetCoordinateHelper(c, a.serv)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
