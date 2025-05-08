package main

import (
	"github.com/yikuanzz/go-blog/core"
	"github.com/yikuanzz/go-blog/global"
)

func main() {
	global.Config = core.InitConf()
	global.Log = core.InitLogger()

	core.RunServer()
}
