package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"strconv"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// CreatePostHandler 创建一条帖子
func CreatePostHandler(c *gin.Context) {
	//参数绑定校验
	var post = new(models.Post)
	if err := c.ShouldBindJSON(post); err != nil {
		zap.L().Error("c.ShouldBindJSON", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	//创建帖子1.通过token获取用户ID
	userID, err := GetCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	post.AuthorID = userID
	//p := &models.Post{
	//	ID:          snowflake.GenID(),
	//	AuthorID:    c.GetInt64("userId"),
	//	CommunityID: post.CommunityID,
	//	Title:       post.Title,
	//	Content:     post.Content,
	//}
	if err = logic.CreatePost(post); err != nil {
		zap.L().Error("logic.CreatePost fail", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//返回数据
	ResponseSuccess(c, nil)
}

// GetPostHandler  获取一条帖子
func GetPostHandler(c *gin.Context) {
	//参数校验绑定
	strId := c.Param("id")
	pid, err := strconv.ParseInt(strId, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	//业务处理
	data, err := logic.GetPost(pid)
	zap.L().Error("logic.GetPost fail", zap.Error(err))
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	//返回数据
	ResponseSuccess(c, data)

}

// GetPostListHandler 分页获取帖子
func GetPostListHandler(c *gin.Context) {

	//获取分页参数
	page, size := getPage(c)

	//业务处理
	dataList, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("logic.GetPostList fail ", zap.Error(err))
		return
	}
	//返回数据
	ResponseSuccess(c, dataList)
}

// GetPostListHandler2 升级版帖子列表接口
// @Summary 升级版帖子列表接口
// @Description 可按社区按时间或分数排序查询帖子列表接口
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query models.ParamPostList false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /getPostList2 [get]
func GetPostListHandler2(c *gin.Context) {
	// GetPostList2 升级版帖子查询，通过create time和score进行排序
	//1.获取参数
	//2.去redis查询id列表
	//3.根据id查询帖子详细信息  http://127.0.0.1:8081/api/v1/getPostList?page=1&size=10&order=time
	p := &models.ParamPostList{
		Page:  1,
		Size:  10,
		Order: models.OrderTime,
	}
	err := c.ShouldBindQuery(p)
	if err != nil {
		zap.L().Error("GetPostListHandler2 fail", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	//获取分页参数
	//page, size := getPage(c)

	//业务处理
	dataList, err := logic.GetPostListNew(p)
	if err != nil {
		zap.L().Error("logic.GetPostList fail ", zap.Error(err))
		return
	}
	//返回数据
	ResponseSuccess(c, dataList)
}

// GetPostListByCommunityHandler 根据社区去查询帖子列表
//func GetPostListByCommunityHandler(c *gin.Context) {
//	p := &models.ParamPostList{
//		Page:        1,
//		Size:        10,
//		Order:       models.OrderTime,
//		CommunityID: 0,
//	}
//	err := c.ShouldBindQuery(p)
//	if err != nil {
//		zap.L().Error("GetPostListByCommunity fail", zap.Error(err))
//		ResponseError(c, CodeInvalidParam)
//		return
//	}
//
//	//获取分页参数
//	//page, size := getPage(c)
//
//	//业务处理
//	dataList, err := logic.GetPostListByCommunity(p)
//	if err != nil {
//		zap.L().Error("logic.GetPostList fail ", zap.Error(err))
//		return
//	}
//	//返回数据
//	ResponseSuccess(c, dataList)
//}
