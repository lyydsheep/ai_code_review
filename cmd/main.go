package main

import (
	"github.com/faiz/llm-code-review/common/enum"
	log "github.com/faiz/llm-code-review/common/logger"
	"github.com/faiz/llm-code-review/config"
	"github.com/faiz/llm-code-review/dal/cache"
	"github.com/faiz/llm-code-review/dal/dao"
	"github.com/gin-gonic/gin"
)

func init() {
	config.InitConfig()
	log.InitLogger()
	cache.RedisInit()
	dao.InitGormLogger()
	dao.InitDB()
}

func main() {
	if config.App.Env == enum.ModePROD {
		gin.SetMode(gin.ReleaseMode)
	}
	app := InitializeApp()
	if err := app.Run(config.App.Host); err != nil {
		panic(err)
	}
}
