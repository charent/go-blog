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
		return
	}

	nPage, latestArticles := ArticleService.GetHomeArticles(page)

	// 查到的文章为0，但是页数不为0（数据库中有文章），那就是请求的页数错了
	if latestArticles == nil {
		if nPage != 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"message": "页数错误",
				"page": page,
			})
			c.Abort()
			return
		} else {
			// 查到的文章为0，而且文章总数也为0
			// 构造一个空数组，防止返回null
			var temp = make([]service.HomeArticleJson, 0)
			latestArticles = &temp
		}

	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"articles": *latestArticles,
		"totalPage":nPage,
		"currentPage": page,
	})
}
