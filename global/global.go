package global

import (
	"github.com/yikuanzz/go-blog/config"
	"go.uber.org/zap"
)

var (
	Config *config.Config
	Log    *zap.Logger
)
