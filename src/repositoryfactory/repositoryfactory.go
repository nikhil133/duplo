package repositoryfactory

import (
	"fmt"

	"github.com/nikhil133/duplo/src/dbstorefactory"
	"github.com/nikhil133/duplo/src/repository"
	"github.com/nikhil133/duplo/src/repositoryinterface"
)

func NewWeatherRepository() repositoryinterface.WeatherRepository {
	dbrepo := dbstorefactory.NewMySql()
	db := dbrepo.GetMysqlConnection()
	fmt.Println("DB connection ", db)
	return &repository.WeatherRepository{
		DB: db,
	}
}
