package controller

import "bluebell/models"

// bluebell/controller/docs_models.go

// _ResponsePostList 帖子列表接口响应数据
type _ResponsePostList struct {
	Code    Recode                  `json:"code"`    // 业务响应状态码
	Message string                  `json:"message"` // 提示信息
	Data    []*models.ApiPostDetail `json:"data"`    // 数据
}

type _ResponseCommunityList struct {
	Code    Recode              `json:"code"`    // 业务响应状态码
	Message string              `json:"message"` // 提示信息
	Data    []*models.Community `json:"data"`    // 数据
}
