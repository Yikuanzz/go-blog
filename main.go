package main

import (
	"github.com/yikuanzz/go-blog/core"
	"github.com/yikuanzz/go-blog/flag"
	"github.com/yikuanzz/go-blog/global"
	"github.com/yikuanzz/go-blog/initialize"
)

// cleanUp 清理函数，用于关闭所有连接
func cleanUp() {
	// 关闭 MySQL 连接
	if global.DB != nil {
		sqlDB, err := global.DB.DB()
		if err == nil {
			sqlDB.Close()
		}
	}

	// 关闭 Redis 连接
	err := global.Redis.Close()
	if err != nil {
		global.Log.Error("Failed to close Redis connection")
	}

	// 因为 ES 客户端会自动管理连接池，所以不需要显式关闭
	// 关闭日志
	if global.Log != nil {
		_ = global.Log.Sync()
	}
}

func main() {
	// 初始化配置
	global.Config = core.InitConf()

	// 初始化日志
	global.Log = core.InitLogger()

	// 初始化其他配置
	initialize.OtherInit()

	// 初始化 MySQL 连接
	global.DB = initialize.InitMySQL()

	// 初始化 Redis 连接
	global.Redis = initialize.ConnectRedis()

	// 初始化 ES 连接
	global.ESClient = initialize.ConnectEs()

	// 关闭连接
	defer cleanUp()

	// 初始化定时任务
	initialize.InitCron()

	// 初始化 CLI
	flag.InitFlag()

	// 启动服务器
	core.RunServer()
}
