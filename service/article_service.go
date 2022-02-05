package service

import (
	"go-blog/model"
	"go-blog/model/entity"
	"go-blog/utils"
	"go-blog/utils/mylog"
	"math"
)

// Article 这里要用等于，相当于别名
type Article =  entity.Article
type Markdown = entity.Markdown

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

// PublishArticleJson 发布文章发送到服务器的json
type PublishArticleJson struct {
	CategoryFirstName 	string	`json:"categoryFirstName" binding:"required"`
	CategorySecondName	string	`json:"categorySecondName" binding:"required"`
	Title				string	`json:"title" binding:"required"`
	MarkdownText		string	`json:"markdownText" binding:"required"`
	// IsPrivate 不能写 binding:"required"，Gin框架的required tag 要求请求的int bool string不能为初始类型（0，false，"")，否则会报错该字段没有填
	IsPrivate			bool 	`json:"isPrivate"`
}

// InsertArticle 插入新文章
func (a *ArticleService) InsertArticle(userId int, reqJson *PublishArticleJson) (articleId int) {
	var article Article
	var markdown Markdown

	// bool to int
	var private int = 0
	if reqJson.IsPrivate {
		private = 1
	}

	// 获取二级分类的id作为文章的分类id
	categoryFirstId, categorySecondId := categorySecondModel.FindIdByNameAndOwnerId(
		userId, reqJson.CategoryFirstName,reqJson.CategorySecondName)

	var updateTime = utils.GetFormatTime()

	// 构造插入的article对象
	article.OwnerId = userId
	article.Title = reqJson.Title

	//取前128个字符作为摘要
	var abstractEnd int = len(reqJson.MarkdownText)
	if abstractEnd > 128 {
		abstractEnd = 128
	}
	article.Abstract = reqJson.MarkdownText[0: abstractEnd] + "......"

	article.CategoryId = categorySecondId
	article.Private = private
	article.PublishTime = updateTime
	article.LastUpdateTime = updateTime

	// 先插入文章表，获取新插入文章的id
	articleId = articleModel.InsertArticle(&article)

	// articleId小于等于0，没有插入成功
	if articleId <= 0 {
		return
	}

	//构造markdown对象
	markdown.ArticleId = articleId
	markdown.Content = reqJson.MarkdownText

	// 插入markdown内容
	newArticleId := markdownModel.InsertMarkdown(&markdown)
	if newArticleId <= 0 {
		return
	}

	// 更新分类的文章数
	row1 := categoryFirstModel.AddArticleCount(userId, categoryFirstId)
	row2 := categorySecondModel.AddArticleCount(userId, categoryFirstId, reqJson.CategorySecondName)

	if row1 + row2 != 2 {
		mylog.Error.Printf("update article_count error,message: " + reqJson.CategoryFirstName + "," + reqJson.CategorySecondName)
	}

	return
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

	findArticles:= articleModel.FindLatestArticle(start, end)
	if findArticles == nil {
		return
	}

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
