package repository

import (
	"github.com/jinzhu/gorm"
)

/*
type WeatherRepository struct {
	DB dbstoreinterface.DbStoreInterface
}*/

type WeatherRepository struct {
	DB *gorm.DB
}

func (r *WeatherRepository) Insert(obj interface{}) error {
	if err := r.DB.Create(obj).Error; err != nil {
		return err
	}
	return nil
}

func (r *WeatherRepository) Get(obj interface{}) error {

	err := r.DB.Debug().Order("created_at desc").Find(obj).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *WeatherRepository) Delete(id uint) error {
	//database := r.DB.GetMysqlConnection()
	//defer r.DB.Close(database)
	//
	//err := database.Debug().Exec("delete from coordinates where id=?", id).Error
	err := r.DB.Debug().Exec("delete from coordinates where id=?", id).Error
	if err != nil {
		return err
	}
	return nil
}
