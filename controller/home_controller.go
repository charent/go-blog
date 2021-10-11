package controller

import (
	"github.com/gin-gonic/gin"
	"go-blog/service"
	"net/http"
	"strconv"
)

type HomeController struct {

}

var ArticleService service.ArticleService

// GetHomeArticles 获取首页最新更新的文章
func (h *HomeController)GetHomeArticles(c *gin.Context) {

	page, err := strconv.Atoi(c.Param("page"))

	if err != nil || page < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"message": "参数错误",
			"param": page,
			})
		c.Abort()
	}

	nPage, latestArticles := ArticleService.GetHomeArticles(page)

	if latestArticles == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"message": "页数错误",
			"page": page,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"articles": *latestArticles,
		"totalPage":nPage,
		"currentPage": page,
	})
}
