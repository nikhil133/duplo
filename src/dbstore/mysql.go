package dbstore

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/nikhil133/duplo/utils/constant"
	mocket "github.com/selvatico/go-mocket"
	"github.com/spf13/viper"
)

type MySql struct {
	Db *gorm.DB
}

func InitConnection() *gorm.DB {
	if constant.Test {
		mocket.Catcher.Register() // Safe register. Allowed multiple calls to save
		mocket.Catcher.Logging = true
		// GORM
		db, _ := gorm.Open(mocket.DriverName, "connection_string") // Can be any connection string
		return db
	}
	fmt.Println("I am here")
	user := viper.GetString("mysql.user")
	password := viper.GetString("mysql.password")
	host := viper.GetString("mysql.host")
	dbname := viper.GetString("mysql.dbname")
	var mysqlHost = fmt.Sprint(user, ":", password, "@(", host, ")/", dbname, "?parseTime=true")
	fmt.Println(mysqlHost)
	db, err := gorm.Open("mysql", mysqlHost)
	if err != nil {
		panic("DB error " + err.Error())
		//logging.Logger.WithError(err).WithField("err", err).Fatal("Database not connected")
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(40)
	return db
}

func (m *MySql) GetMysqlConnection() *gorm.DB {
	if m.Db == nil {
		m.Db = InitConnection()
	}
	return m.Db
}

func (m *MySql) Close(db *gorm.DB) {
	if db != nil {
		db.Close()
	}

}
