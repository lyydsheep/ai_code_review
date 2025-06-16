package main

import (
	"context"
	"github.com/faiz/llm-code-review/logic/service"
	"github.com/gin-gonic/gin"
)

type App struct {
	*gin.Engine
	consumer *service.LLMService
}

func (app *App) Run(host string) error {
	ctx := context.Background()
	app.consumer.Run(ctx)
	if err := app.Engine.Run(host); err != nil {
		return err
	}
	return nil
}

func NewApp(g *gin.Engine, consumer *service.LLMService) *App {
	return &App{Engine: g, consumer: consumer}
}
