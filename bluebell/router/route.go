package router

import (
	"bluebell/controller"
	"bluebell/logger"
	"bluebell/middlewares"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	v1 := r.Group("/api/v1") // 创建API v1版本的路由组
	//注册业务路由
	v1.POST("/signup", controller.SignUpHandler)
	v1.POST("/login", controller.LoginHandler)

	v1.Use(middlewares.JWTAuthMiddleware()) //应用JWT认证中间件
	{
		// 社区相关路由
		v1.GET("/community", controller.CommunityHandler)           // 获取社区列表
		v1.GET("/community/:id", controller.CommunityDetailHandler) // 获取特定社区详情
		// 帖子相关路由
		v1.POST("/post", controller.CreatePostHandler)       // 创建新帖子
		v1.GET("/post/:id", controller.GetPostDetailHandler) // 获取特定帖子详情
		v1.GET("/posts", controller.GetPostListHandler)      // 获取帖子列表
		v1.GET("/posts2", controller.GetPostListHandler2)    //根据时间或分数获取帖子列表
		v1.POST("/vote", controller.PostVoteController)      // 帖子投票
	}
	pprof.Register(r) //注册pprof相关路由
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
