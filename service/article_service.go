package service

import (
	"go-blog/model"
	"go-blog/model/entity"
	"math"
)

// Article 这里要用等于，相当于别名
type Article =  entity.Article
type ArticleService struct {

}
var articleModel model.ArticleModel
var markdownModel model.MarkdownModel

type HomeArticleJson struct {
	ArticleId 		int
	Title			string
	Abstract		string
	PublishTime		string
	Visited			int
}

// GetArticleDetailByArticleId 根据文章id或者文章信息和文章内容（markdown文本）
func (a *ArticleService) GetArticleDetailByArticleId(articleId int) (article *Article, markdown *entity.Markdown) {
	article = articleModel.FindArticleByArticleId(articleId)
	markdown = markdownModel.FindMarkdownByArticleId(articleId)

	return
}

// GetHomeArticles 获取首页展示的最近更新文章
func (a *ArticleService) GetHomeArticles(page int) (nPage int, articles *[]HomeArticleJson)  {

	// 每一页显示多少个文章
	var articlePerPage int = 10

	// 计算limit的开始和结束
	var start int = page * articlePerPage
	var end int = start + articlePerPage
	totalArticle := articleModel.GetCountOfArticle()

	// 总页数向上取整
	nPage = int(math.Ceil( float64(totalArticle) / float64(articlePerPage)))

	// start大于总文章数，不查了
	if start > totalArticle {
		articles = nil
		return
	}

	findArticles:= articleModel.GetLatestArticle(start, end)

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
