package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/yikuanzz/go-blog/global"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	// 设置 gin 模式
	gin.SetMode(global.Config.System.Env)
	Router := gin.Default()

	// TODO: something

	return Router
}
