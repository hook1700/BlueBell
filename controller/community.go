package controller

import (
	"bluebell/logic"
	"strconv"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// CommunityHandle 返回标题列表
// @Summary 返回标题列表
// @Description 返回标题列表
// @Tags 社区
// @Param Authorization header string ture "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseCommunityList
// @Router /community [get]
func CommunityHandle(c *gin.Context) {
	//参数检验
	//业务处理
	cs, err := logic.ListCommunity()
	zap.L().Error("logic.ListCommunity is fail", zap.Error(err))
	if err != nil {
		ResponseError(c, CodeCommunityIsNull)
		return
	}

	//返回数据
	ResponseSuccess(c, cs)
}

func CommunityDetailHandle(c *gin.Context) {
	//参数检验
	strId := c.Param("id")
	id, err := strconv.ParseInt(strId, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	//业务处理
	community, err := logic.CommunityDetail(id)

	if err != nil {
		zap.L().Error("logic.CommunityDetail is fail", zap.Error(err))
		ResponseError(c, CodeGetCommunityFail)
	}
	//返回数据
	ResponseSuccess(c, community)
}
