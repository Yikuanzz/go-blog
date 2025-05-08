package global

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-redis/redis"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/yikuanzz/go-blog/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config     *config.Config
	Log        *zap.Logger
	DB         *gorm.DB
	ESClient   *elasticsearch.TypedClient
	Redis      redis.Client
	BlackCache local_cache.Cache
)
