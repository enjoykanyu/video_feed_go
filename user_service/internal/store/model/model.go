package model

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

// Db 用来承接db变量
var Db *gorm.DB

// InitDb 初始化数据库连接
func InitDb() *gorm.DB {
	var (
		Username = viper.GetString("database.username")
		Password = viper.GetString("database.password")
		Host     = viper.GetString("database.host")
		Port     = viper.GetInt("database.port")
		DbName   = viper.GetString("database.dbname")
	)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", Username, Password, Host, Port, DbName)
	fmt.Println(Username)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal("数据库连接失败,报错信息" + err.Error())
	}
	// 设置连接池，空闲连接
	db.DB().SetMaxIdleConns(50)
	// 打开链接
	db.DB().SetMaxOpenConns(100)
	// 表明禁用后缀加s
	db.SingularTable(true)
	// 启用Logger，显示详细日志
	db.LogMode(viper.GetBool("app.debug"))
	return db
}
