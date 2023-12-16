package main

import (
	"Auto/router"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := gin.Default()
	//设置路由
	router.InitRouter(r)

	r.Run(":8080") // 监听并服务在 :8080 端口
}
