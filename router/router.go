package router

import (
	"github.com/gin-gonic/gin"
	"go-blog/config"
	"go-blog/controller"
	"go-blog/middleware"
	"go-blog/utils/mylog"
)

func InitRouter() {
	router := gin.Default()

	//设置允许跨域
	router.Use(middleware.Cors())
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	var authMiddleware = middleware.AuthMiddleware

	// 404 路由
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": 404, "message": "Page not found"})
	})

	var homeController controller.HomeController

	// 公共api
	publicApi := router.Group("/")
	{
		publicApi.POST("/user/login", authMiddleware.LoginHandler)
		publicApi.GET("/home/articles", homeController.GetHomeArticles)
	}

	var ManagerController controller.ManagerController

	// 需要认证的 api分组 manager/*
	authApi := router.Group("/manager")
	authApi.Use(authMiddleware.MiddlewareFunc())
	{
		authApi.POST("/home", ManagerController.Home)
	}

	//启动Gin
	err := router.Run(config.Server.Api.Host + ":" + config.Server.Api.Port)

	if err != nil {
		mylog.Error.Printf("gin router start error, message: %v", err)
	}
}
