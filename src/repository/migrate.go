package repository

import "github.com/nikhil133/duplo/src/models"

func (r *WeatherRepository) Migrate() {
	r.DB.Debug().AutoMigrate(&models.Coordinate{})
}
