//go:build wireinject

package main

import (
	"github.com/faiz/llm-code-review/api/controller"
	"github.com/faiz/llm-code-review/api/router"
	"github.com/faiz/llm-code-review/common/middleware"
	"github.com/faiz/llm-code-review/dal/cache"
	"github.com/faiz/llm-code-review/dal/dao"
	"github.com/faiz/llm-code-review/logic/appService"
	"github.com/faiz/llm-code-review/logic/domainService"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitializeApp() *gin.Engine {
	wire.Build(router.RegisterRoutersAndMiddleware,
		middleware.GetHandlerFunc, controller.NewBuildController,
		wire.Bind(new(appService.DemoAppService), new(*appService.DemoAppServiceV1)), appService.NewDemoAppServiceV1,
		wire.Bind(new(domainService.DemoDomainService), new(*domainService.DemoDomainServiceV1)), domainService.NewDemoDomainServiceV1,
		wire.Bind(new(dao.DemoDAO), new(*dao.DemoDAOV1)), wire.Bind(new(cache.DemoCache), new(*cache.DemoCacheV1)),
		dao.NewDemoDAO, cache.NewCacheV1,
	)
	return nil
}
