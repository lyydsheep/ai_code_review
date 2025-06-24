package main

import (
	"context"
	"github.com/faiz/llm-code-review/common/enum"
	log "github.com/faiz/llm-code-review/common/logger"
	"github.com/faiz/llm-code-review/config"
	"github.com/faiz/llm-code-review/dal/dao"
	"github.com/faiz/llm-code-review/logic/service"
	"github.com/gin-gonic/gin"
)

func init() {
	config.InitConfig()
	log.InitLogger()
	dao.InitGormLogger()
	dao.InitDB()
}

func main() {
	if config.App.Env == enum.ModePROD {
		gin.SetMode(gin.ReleaseMode)
	}
	service.ListenAndSend(context.Background())
	app := InitializeApp()
	if err := app.Run(config.App.Host); err != nil {
		panic(err)
	}
}
