package models

import (
	"github.com/nikhil133/duplo/src/dbstore"
	"gorm.io/gorm"
)

type Coordinate struct {
	gorm.Model
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

func Migrate() {
	database := dbstore.InitConnection()
	defer database.Close()
	database.Debug().AutoMigrate(&Coordinate{})

}
