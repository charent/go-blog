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
	var publicController = controller.PublicHandler
	// 404 路由
	router.NoRoute(publicController.PageNotFound)

	var homeController controller.HomeController

	// 公共api
	publicApi := router.Group("/")
	{
		// 用户登录
		publicApi.POST("/user/login", authMiddleware.LoginHandler)

		// 获取首页的文章（按时间逆序）,注意：articles复数
		publicApi.GET("/home/articles/:page", homeController.GetHomeArticles)

		// 获取文章详情，注意：article单数
		publicApi.GET("/home/article/:articleId", homeController.GetArticleDetail)

	}

	var ManagerController controller.ManagerController

	// 需要认证的api分组, 路径：manager/*
	authApi := router.Group("/manager")
	authApi.Use(authMiddleware.MiddlewareFunc())
	{
		authApi.POST("/home", ManagerController.Home)

		// 文章分类
		authApi.GET("/category/first", ManagerController.GetCategoryFirst)
		authApi.PUT("/category/first", ManagerController.PutCategoryFirst)
		authApi.GET("/category/second", ManagerController.GetCategorySecond)
		authApi.PUT("/category/second", ManagerController.PutCategorySecond)

		// 一级分类重命名和删除
		authApi.POST("/category/first/rename", ManagerController.RenameFirstCategory)
		authApi.DELETE("/category/first", ManagerController.DeleteFirstCategory)

		// 二级分类重命名和删除
		authApi.POST("/category/second/rename", ManagerController.RenameSecondCategory)
		authApi.DELETE("/category/second", ManagerController.DeleteSecondCategory)

		// 文章发布、管理
		authApi.PUT("/article", ManagerController.PutArticle)

	}

	//启动Gin
	err := router.Run(config.Server.Api.Host + ":" + config.Server.Api.Port)

	if err != nil {
		mylog.Error.Printf("gin router start error, message: %v", err)
	}
}
