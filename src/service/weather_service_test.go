package service_test

import (
	"database/sql"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nikhil133/duplo/mocks"
	"github.com/nikhil133/duplo/src/models"
	"github.com/nikhil133/duplo/src/service"
	"github.com/nikhil133/duplo/utils/constant"
	utils "github.com/nikhil133/duplo/utils/request"
	mocket "github.com/selvatico/go-mocket"
)

func TestGetForcast(t *testing.T) {
	constant.Test = true
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mocket.Catcher.Register()
	mocket.Catcher.Logging = false
	m := mocks.NewMockWeatherRepository(ctrl)
	c := models.Coordinate{}
	m.EXPECT().Insert(&c).Return(nil).AnyTimes()
	s := service.WeatherService{
		Repo: m,
	}
	req := utils.Request{}
	s.GetForcast(req)
}

func TestGetForcastError(t *testing.T) {
	constant.Test = true
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock WeatherRepository
	m := mocks.NewMockWeatherRepository(ctrl)

	// Set up expectations for the WeatherRepository's Insert method
	c := models.Coordinate{}
	mockError := sql.ErrConnDone
	m.EXPECT().Insert(&c).Return(mockError).AnyTimes()

	// Create a WeatherService with the mock repository
	s := service.WeatherService{
		Repo: m,
	}

	// Call the GetForcast method
	req := utils.Request{}
	s.GetForcast(req)

	// Assert that the error returned matches the expected error
}

func TestGetCoordinates(t *testing.T) {
	constant.Test = true

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock WeatherRepository
	m := mocks.NewMockWeatherRepository(ctrl)

	// Set up expectations for the WeatherRepository's Insert method
	c := []service.Coordinate{}
	m.EXPECT().Get(&c).Return(nil).AnyTimes()

	// Create a WeatherService with the mock repository
	s := service.WeatherService{
		Repo: m,
	}

	// Call the GetForcast method
	s.GetCoordinates()

}

func TestGetCoordinatesError(t *testing.T) {
	constant.Test = true

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock WeatherRepository
	m := mocks.NewMockWeatherRepository(ctrl)

	// Set up expectations for the WeatherRepository's Insert method
	c := []service.Coordinate{}
	mockErr := sql.ErrNoRows
	m.EXPECT().Get(&c).Return(mockErr).AnyTimes()

	// Create a WeatherService with the mock repository
	s := service.WeatherService{
		Repo: m,
	}

	// Call the GetForcast method
	s.GetCoordinates()

}

func TestDeleteCoordinates(t *testing.T) {
	constant.Test = true

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	id := uint(1)
	// Create a mock WeatherRepository
	m := mocks.NewMockWeatherRepository(ctrl)
	m.EXPECT().Delete(id).Return(nil).AnyTimes()
	s := service.WeatherService{
		Repo: m,
	}

	// Call the GetForcast method
	s.DeleteCoordinates(id)

}
