package home

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-blog/services/articles"
)

// GetHomeArticles 获取首页最新更新的文章
func GetHomeArticles(c *gin.Context) {
	newArticles := articles.GetHomeArticles()
	fmt.Printf("%v", newArticles)
	c.Next()
}
