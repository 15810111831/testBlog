package common

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"test.blog.com/testBlog/model"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	user := "root"
	password := "abc"
	host := "127.0.0.1"
	port := 3306
	dbName := "test_blog"
	connArgs := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		port,
		dbName,
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
