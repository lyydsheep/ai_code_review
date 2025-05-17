package app

import (
	"github.com/faiz/llm-code-review/config"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Pagination struct {
	Page      int `json:"page"`
	PageSize  int `json:"page_size"`
	TotalRows int `json:"total_rows"`
}

func NewPagination(c *gin.Context) *Pagination {
	page, _ := strconv.Atoi(c.GetString("page"))
	if page < 1 {
		page = 1
	}
	pageSize, _ := strconv.Atoi(c.GetString("page_size"))
	if pageSize < 1 || pageSize > config.App.Pagination.MaxSize {
		pageSize = config.App.Pagination.DefaultSize
	}
	return &Pagination{
		Page:      page,
		PageSize:  pageSize,
		TotalRows: 0,
	}
}

func (p *Pagination) Offset() int {
	return (p.Page - 1) * p.PageSize
}

func (p *Pagination) GetPage() int {
	return p.Page
}

func (p *Pagination) SetTotalRows(totalRows int) {
	p.TotalRows = totalRows
}

func (p *Pagination) GetPageSize() int {
	return p.PageSize
}
