package compat

import (
	"test/service/gogf/model"

	"github.com/gogf/gf/v2/frame/g"
)

// 检测记录：
// 1.检测记录搜索翻页(sWord,pageNo,pageSize)
type DetectRecordSearchReq struct {
	g.Meta  `path:"/detect_record/search" method:"get" summary:"查询检测记录(搜索分页)" tags:"检测记录"`
	Name    string `json:"name"     v:"length:0,30#检测记录搜索词|检测记录搜索词长度为:{min}到:{max}位" dc:"检测记录搜索词"`
	OrderBy string `json:"order_by" v:"length:0,30#排序字段不正确"            d:"id"                 dc:"排序字段"`
	Desc    string `json:"desc"     v:"in:desc,asc#排序方式不正确"            d:"desc"               dc:"排序方式"`
	model.CommonPaginationReq
}
type DetectRecordSearchRes struct {
	// List    []*entity.DetectRecord `json:"list"             dc:"检测记录列表(搜索分页)"` //
	OrderBy string                 `json:"order_by"        d:"id"        dc:"排序字段"`
	Desc    string                 `json:"desc"            d:"desc"      dc:"排序方式"`
	model.CommonPaginationRes
}