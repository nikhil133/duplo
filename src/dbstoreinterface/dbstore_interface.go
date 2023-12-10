package dbstoreinterface

import "github.com/jinzhu/gorm"

type DbStoreInterface interface {
	GetMysqlConnection() *gorm.DB
	Close(*gorm.DB)
}
