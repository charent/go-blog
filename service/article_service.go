package service

import (
	"go-blog/model"
	"go-blog/model/entity"
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
func (a *ArticleService) GetHomeArticles() (articles *[]HomeArticleJson)  {

	findArticles, _ := ArticleModel.GetLatestArticle(0, 10)

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
