package core

import (
	"log"

	"github.com/yikuanzz/go-blog/config"
	"github.com/yikuanzz/go-blog/utils"
	"gopkg.in/yaml.v3"
)

// InitConf 从 YAML 文件加载配置
func InitConf() *config.Config {
	c := &config.Config{}
	yamlConf, err := utils.LoadYAML()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("Failed to unmarshal YAML configuration: %v", err)
	}
	return c
}
