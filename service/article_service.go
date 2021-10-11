package service

import (
	"go-blog/model"
	"go-blog/model/entity"
	"math"
)

type Article =  entity.Article
type ArticleService struct {

}
var ArticleModel model.ArticleModel

type HomeArticleJson struct {
	ArticleId 		int
	Title			string
	Abstract		string
	PublishTime		string
	Visited			int
}

// GetHomeArticles 获取首页展示的最近更新文章
func (a *ArticleService) GetHomeArticles(page int) (nPage int, articles *[]HomeArticleJson)  {

	// 每一页显示多少个文章
	var articlePerPage int = 10

	// 计算limit的开始和结束
	var start int = page * articlePerPage
	var end int = start + articlePerPage
	totalArticle := ArticleModel.GetCountOfArticle()

	// 总页数向上取整
	nPage = int(math.Ceil( float64(totalArticle) / float64(articlePerPage)))

	// start大于总文章数，不查了
	if start > totalArticle {
		return
	}

	findArticles:= ArticleModel.GetLatestArticle(start, end)

	var returnArticle []HomeArticleJson

	// 过滤掉不需要的字段
	for _, article := range *findArticles {

		var temp HomeArticleJson
		temp.ArticleId = article.ArticleId
		temp.Title = article.Title
		temp.Abstract = article.Abstract
		temp.PublishTime = article.PublishTime
		temp.Visited = article.Visited
		returnArticle = append(returnArticle,temp)
	}

	articles = &returnArticle

	return
}
