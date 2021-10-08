package routes

import (
	"bluebell/controller"
	_ "bluebell/docs" // 千万不要忘了导入把你上一步生成的docs
	"bluebell/logger"
	"bluebell/middelwares"
	"bluebell/settings"
	"net/http"

	"github.com/gin-gonic/gin"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func SetUp(mode string) *gin.Engine {
	//设置成发布的模式
	//gin.SetMode(gin.ReleaseMode)
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.Use(middelwares.Cors())

	r.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, settings.Conf.Version)
	})

	v1 := r.Group("/api/v1")
	v1.GET("/", controller.HelloHandle)
	//swagger
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	//注册业务路由
	//注册接口
	v1.POST("signup", controller.SignUpHandle)

	//登录接口
	v1.POST("login", controller.LoginHandle)

	//使用中间件
	v1.Use(middelwares.JWTAuthMiddleware())
	{
		v1.GET("/community", controller.CommunityHandle)
		v1.GET("/communityDetail/:id", controller.CommunityDetailHandle)

		v1.POST("/post", controller.CreatePostHandler)
		v1.GET("/getPost/:id", controller.GetPostHandler)
		v1.GET("/getPostList", controller.GetPostListHandler)
		v1.GET("/getPostList2", controller.GetPostListHandler2)

		//投票
		v1.POST("/vote", controller.PostVoteController)
	}
	//登录验证
	//v1.GET("/ping", middelwares.JWTAuthMiddleware(), func(c *gin.Context) {
	//	username := c.MustGet("username").(string)
	//	c.JSON(http.StatusOK, gin.H{
	//		"msg": username,
	//	})
	//})

	return r
}
