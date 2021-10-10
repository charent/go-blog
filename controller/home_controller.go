package controller

import (
	"github.com/gin-gonic/gin"
	"go-blog/service"
)

type HomeController struct {

}

var ArticleService service.ArticleService

// GetHomeArticles 获取首页最新更新的文章
func (h *HomeController)GetHomeArticles(c *gin.Context) {
	latestArticles := ArticleService.GetHomeArticles()

	c.JSON(200, gin.H{
		"code": 200,
		"articles": *latestArticles,
	})
}
