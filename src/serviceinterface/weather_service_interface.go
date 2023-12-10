package serviceinterface

import (
	"github.com/nikhil133/duplo/src/service"
	utils "github.com/nikhil133/duplo/utils/request"
)

type WeatherServiceInterface interface {
	GetForcast(req utils.Request) (*service.Forecast, error)
	GetCoordinates() ([]service.Coordinate, error)
	DeleteCoordinates(id uint) error
}
