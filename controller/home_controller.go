package controller

import (
	"github.com/gin-gonic/gin"
	"go-blog/service"
	"net/http"
	"strconv"
)

type HomeController struct {

}

var articleService service.ArticleService


func (h HomeController) GetArticleDetail(c *gin.Context)  {
	articleId, err := strconv.Atoi(c.Param("articleId"))

	if err != nil || articleId < 0 {
		PublicHandler.ParamError(c, "文章id错误", articleId)
		c.Abort()
		return
	}

	article, markdown := articleService.GetArticleDetailByArticleId(articleId)

	// 任何一个为nil就是404
	if article == nil || markdown == nil {
		PublicHandler.PageNotFound(c)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"article": article,
		"markdown":markdown.Content,
	})
}


// GetHomeArticles 获取首页最新更新的文章
func (h *HomeController)GetHomeArticles(c *gin.Context) {

	page, err := strconv.Atoi(c.Param("page"))

	if err != nil || page < 0 {
		PublicHandler.ParamError(c, "页数错误", page)
		c.Abort()
		return
	}

	nPage, latestArticles := articleService.GetHomeArticles(page)

	// 查到的文章为0，但是页数不为0（数据库中有文章），那就是请求的页数错了
	if latestArticles == nil {
		if nPage != 0 {
			PublicHandler.ParamError(c, "页数错误", page)
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
