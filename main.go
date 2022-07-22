package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func main() {
	// 线上
	dsn := "root:123456@tcp(172.17.0.1:3306)/hxp?charset=utf8mb4&parseTime=True&loc=Local"
	// 线下
	// dsn := "root:123456@tcp(127.0.0.1:3306)/gin_wall?charset=utf8mb4&parseTime=True&loc=Local"
	//全局模式
	var err error
	DB, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Println("连接数据库失败")
		return
	}
	fmt.Println("连接成功")

}
