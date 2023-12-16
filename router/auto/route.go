package router

import (
	"Auto/handle"
	"github.com/gin-gonic/gin"
)

type LogRouter struct {
}

func (*LogRouter) Route(r *gin.Engine) {
	//Todo
	database := handle.New()
	// 已完成
	r.GET("/getup", database.GetUserDataBase)       //动态查询
	r.POST("/setup", database.SetUserDataBase)      //动态修改
	r.POST("/deleted", database.DeleteUserDataBase) //删除
	// todo
	r.POST("/inset", database.InsetUserDataBase) // 动态插入字段
	//前端获取路由
}
