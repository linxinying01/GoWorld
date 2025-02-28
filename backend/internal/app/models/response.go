package models

import "time"

type CommonResponse struct {
	Code      int         `json:"code"`                 // 状态码
	Msg       string      `json:"msg"`                  // 提示信息
	Data      interface{} `json:"data,omitempty"`       // 数据（使用omitempty，如果为空则不返回）
	Timestamp string      `json:"timestamp,omitempty"`  // 响应时间戳
	RequestId string      `json:"request_id,omitempty"` // 请求ID
}

// 带分页的响应结构体
type PaginatedResponse struct {
	CommonResponse
	PageInfo Pagination `json:"pagination"` // 分页信息
}

// 分页元数据
type Pagination struct {
	Page      int `json:"page"`       // 当前页码
	PageSize  int `json:"page_size"`  // 每页数量
	Total     int `json:"total"`      // 总记录数
	TotalPage int `json:"total_page"` // 总页数
}

// 成功响应构造函数
func NewSuccessResponse(data interface{}) CommonResponse {
	return CommonResponse{
		Code:      200,
		Msg:       "成功",
		Data:      data,
		Timestamp: time.Now().Format(time.RFC3339),
	}
}

// 错误响应构造函数
func NewErrorResponse(code int, msg string) CommonResponse {
	return CommonResponse{
		Code:      code,
		Msg:       msg,
		Timestamp: time.Now().Format(time.RFC3339),
	}
}

// 带分页的成功响应
func NewPaginatedResponse(page, pageSize, total int, data interface{}) PaginatedResponse {
	totalPage := 0
	if pageSize > 0 {
		totalPage = (total + pageSize - 1) / pageSize
	}

	return PaginatedResponse{
		CommonResponse: NewSuccessResponse(data),
		PageInfo: Pagination{
			Page:      page,
			PageSize:  pageSize,
			Total:     total,
			TotalPage: totalPage,
		},
	}
}
