package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"os"
)

var (
	DB *gorm.DB
)

func init() {

	err := initMySQL()
	if err != nil {
		fmt.Println(err.Error())
		panic("MYSQL CONNECT ERROR")
	}
	log.Println("[INIT MYSQL CONNECTS] success")

}

func initMySQL() (err error) {

	host := os.Getenv("DB_HOSTNAME")
	user := os.Getenv("DB_USERNAME")
	passwd := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_DATABASE")
	port := os.Getenv("DB_PORT")
	debug := os.Getenv("DB_DEBUG")

	dsn := user + ":" + passwd + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}

	//开启debug模式
	if debug == "true" {
		DB.LogMode(true)
	}

	//禁用orm表名复数
	DB.SingularTable(true)
	//空闲时最大连接数
	DB.DB().SetMaxIdleConns(10)
	//设置数据库最大连接数
	DB.DB().SetMaxOpenConns(5000)

	return DB.DB().Ping()
}

func Close() {
	DB.Close()
}
