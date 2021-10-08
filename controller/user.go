package controller

import (
	"bluebell/dao/mysql"
	"bluebell/logic"
	"bluebell/models"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// HelloHandle 测试接口
// @Summary 测试接口
// @Description 测试接口
// @Tags 测试接口
// @Success 200
// @Router / [get]
func HelloHandle(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

// SignUpHandle 处理注册请求的函数
func SignUpHandle(c *gin.Context) {
	//1.获取参数和参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		//判断错误是否为validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		// 非validator.ValidationErrors类型错误直接返回
		if !ok {
			zap.L().Error("signup param is invalid", zap.Error(err))
			ResponseError(c, CodeInvalidParam)
			//c.JSON(http.StatusOK, gin.H{
			//	"msg": err.Error(),
			//})
			return
		}
		// validator.ValidationErrors类型错误则进行翻译

		//通过封装返回
		ResponseErrorWithMsg(c, CodeServerBusy, removeTopStruct(errs.Translate(trans)))
		//c.JSON(http.StatusOK, gin.H{
		//	"msg": removeTopStruct(errs.Translate(trans)),
		//})
		return
	}
	//fmt.Println(p)
	//手动对参数进行校验
	//if len(p.Username) == 0 || len(p.Password) == 0 || len(p.RePassword) == 0 || p.Password != p.RePassword {
	//
	//	zap.L().Error("signup param is invalid")
	//	c.JSON(http.StatusOK, gin.H{
	//		"msg": "参数错误",
	//	})
	//	return
	//}

	//2.业务处理
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("login Sign failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorExistUser) {
			ResponseError(c, CodeUserExist)
			//c.JSON(http.StatusOK, gin.H{
			//	"msg": "用户注册失败",
			//})
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	//3.返回响应
	ResponseSuccess(c, nil)
	//c.JSON(http.StatusOK, gin.H{
	//	"msg": "ok",
	//})
}

// LoginHandle 用户登录接口
func LoginHandle(c *gin.Context) {
	//1.获取参数和校验参数
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		//判断错误是否为validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		// 非validator.ValidationErrors类型错误直接返回
		if !ok {
			zap.L().Error("login param is invalid", zap.Error(err))
			ResponseError(c, CodeInvalidParam)
			//c.JSON(http.StatusOK, gin.H{
			//	"msg": err.Error(),
			//})
			return
		}
		// validator.ValidationErrors类型错误则进行翻译

		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		//c.JSON(http.StatusOK, gin.H{
		//	"msg": removeTopStruct(errs.Translate(trans)),
		//})
		return
	}

	//2.业务处理
	user, token, err := logic.Login(p)
	if err != nil {

		//后层返回的错误用日志记录，给用户看得错误用json返回
		zap.L().Error("logic fail", zap.String("username", p.Username), zap.Error(err))
		if errors.Is(err, mysql.ErrorNotExistUser) {
			ResponseErrorWithMsg(c, CodeUserNotExist, CodeUserNotExist.Msg())
			return
		}

		if errors.Is(err, mysql.ErrorInvalidPassword) {
			ResponseErrorWithMsg(c, CodeInValidPassword, CodeInValidPassword.Msg())
			return
			//c.JSON(http.StatusOK, gin.H{
			//	"msg": "用户名或密码错误",
			//})
			//return
		}

	}
	//3.返回数据
	ResponseSuccess(c, gin.H{
		"user_id":   user.UserId,
		"user_name": user.Username,
		"token":     token,
	})
	//c.JSON(http.StatusOK, gin.H{
	//	//"msg": "登录成功",
	//
	//})
}
