package dao

import (
	"Auto/config"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	Db  *sql.DB
	err error
)

func init() {
	Db, err = sql.Open("mysql", config.Conf.MC.Name+":"+config.Conf.MC.Password+"@tcp("+config.Conf.MC.Host+")"+"/userdatabase"+"?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("数据库连接失败 : %s\n", err.Error())
		fmt.Println("mysql", config.Conf.MC.Name+":"+config.Conf.MC.Password+"@tcp("+config.Conf.MC.Host+")"+"/userdatabase"+"?charset=utf8mb4&parseTime=True&loc=Local")
		//log.Fatal("无法连接到数据库")

	}
	fmt.Println("mysql", config.Conf.MC.Name+":"+config.Conf.MC.Password+"@tcp("+config.Conf.MC.Host+")"+"/userdatabase"+"?charset=utf8mb4&parseTime=True&loc=Local")

	// 确保连接正常
	err = Db.Ping()
	if err != nil {
		//fmt.Printf("数据库不健康 : %s\n", err.Error())
		log.Fatal("数据库不健康")
		//return
	}
	fmt.Printf("数据库已连接")
}
