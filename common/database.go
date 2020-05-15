package common

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"test.blog.com/testBlog/model"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	user := viper.GetString("datasource.user")
	password := viper.GetString("datasource.password")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	dbName := viper.GetString("datasource.dbName")
	charset := viper.GetString("datasource.charset")

	connArgs := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		user,
		password,
		host,
		port,
		dbName,
		charset,
	)
	db, err := gorm.Open("mysql", connArgs)
	if err != nil {
		panic("fail to connect to mysql, err" + err.Error())
	}
	db.AutoMigrate(&model.User{})
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
