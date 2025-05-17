package controller

// controller 包中文件放置handler 方法

import (
	"github.com/faiz/llm-code-review/api/request"
	"github.com/faiz/llm-code-review/common/app"
	"github.com/faiz/llm-code-review/common/errcode"
	"github.com/faiz/llm-code-review/library"
	"github.com/faiz/llm-code-review/logic/service"
	"github.com/gin-gonic/gin"
)

type BuildController struct {
	appDemoService appService.DemoAppService
}

func NewBuildController(app appService.DemoAppService) *BuildController {
	return &BuildController{
		appDemoService: app,
	}
}

func (build *BuildController) TestPagination(c *gin.Context) {
	data := []struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{
		{
			Name: "faiz",
			Age:  18,
		},
		{
			Name: "lyy",
			Age:  10,
		},
	}
	p := app.NewPagination(c)
	p.SetTotalRows(len(data))
	app.NewResponse(c).SetPagination(p).Success(data)
}

func (build *BuildController) TestGormLog(c *gin.Context) {
	ids, err := build.appDemoService.GetAllIdentities(c)
	if err != nil {
		app.NewResponse(c).Error(errcode.ErrServer.WithCause(err))
	}
	app.NewResponse(c).Success(ids)
}

func (build *BuildController) CreateDemoOrder(c *gin.Context) {
	var orderReq request.DemoOrderReq
	if err := c.ShouldBindJSON(&orderReq); err != nil {
		app.NewResponse(c).Error(errcode.ErrParams.WithCause(err))
		return
	}
	rep, err := build.appDemoService.CreateDemoOrder(c, &orderReq)
	if err != nil {
		app.NewResponse(c).Error(errcode.ErrServer.WithCause(err))
	}
	app.NewResponse(c).Success(rep)
}

func (build *BuildController) TestHttpTool(c *gin.Context) {
	whois := library.NewWhoisLib(c)
	res, err := whois.GetHostIpDetail()
	if err != nil {
		// 上层很难根据具体的错误类型返回特定的响应码
		app.NewResponse(c).Error(errcode.ErrServer.WithCause(err))
		return
	}
	app.NewResponse(c).Success(res)
}
