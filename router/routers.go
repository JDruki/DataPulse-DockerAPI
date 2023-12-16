package router

import (
	router "Auto/router/auto"
	"github.com/gin-gonic/gin"
)

type Router interface {
	Route(r *gin.Engine)
}

// RegisterRouter 注册路由初始化方法
type RegisterRouter struct {
}

func New() *RegisterRouter {
	return &RegisterRouter{}
}

// Route 路由封装类
func (*RegisterRouter) Route(ro Router, r *gin.Engine) {
	ro.Route(r)
}

func InitRouter(r *gin.Engine) {
	rg := New()
	rg.Route(&router.LogRouter{}, r)
}
