package servicefactory

import (
	"github.com/nikhil133/duplo/src/repositoryfactory"
	"github.com/nikhil133/duplo/src/service"
	"github.com/nikhil133/duplo/src/serviceinterface"
)

func NewWeatherService() serviceinterface.WeatherServiceInterface {
	return &service.WeatherService{
		Repo: repositoryfactory.NewWeatherRepository(),
	}
}
