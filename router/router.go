package router

import (
	"go-blog/config"
	"go-blog/controllers/adminController"
	"go-blog/middlewares"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()

	//设置允许跨域
	router.Use(middlewares.Cors())
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	var authMiddleware = middlewares.AuthMiddleware
	// 需要认证的 api分组
	authApi := router.Group("/", authMiddleware.MiddlewareFunc())
	{
		authApi.POST("/admin/manager", adminController.AdminLogin)
	}

	// 公共api
	publicApi := router.Group("/")
	{
		publicApi.POST("/user/login", authMiddleware.LoginHandler)
	}


	//启动Gin
	router.Run(config.Server.Api.Host + ":" + config.Server.Api.Port)
}
