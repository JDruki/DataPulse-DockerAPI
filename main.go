package main

import (
	"Auto/config"
	"Auto/router"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := gin.Default()
	//设置路由
	router.InitRouter(r)

	r.Run(":" + config.Conf.USER.Port) // 监听并服务在 :8080 端口
}
