package model

import (
	"go-blog/database"
	"go-blog/model/entity"
	"go-blog/utils/mylog"
)

type Article = entity.Article
var DB = database.MysqlDB
type ArticleModel struct {

}

// InsertArticle 插入新文章
func (a *ArticleModel) InsertArticle(article *Article) (articleId int) {

	res := DB.Create(article)

	if res.Error != nil {
		mylog.Error.Printf("execute sql error, message: %v", res.Error.Error())
		return
	}

	if res.RowsAffected == 0 {
		articleId = -1
		return
	}
	articleId = article.ArticleId
	return
}


// FindArticleByArticleId 根据文章id获取文章信息
func (a *ArticleModel) FindArticleByArticleId(articleId int) (article *Article)  {
	var findArticle Article
	res := DB.Raw("select * from article where article_id = ? and deleted = false and private = false;", articleId).Scan(&findArticle)

	if res.Error != nil {
		mylog.Error.Printf("execute sql error, message: %v", res.Error.Error())
		return
	}

	if res.RowsAffected == 0 {
		return
	}
	article = &findArticle
	return
}

// FindArticleOrderByPublishTime 按照时间逆序查找文章
// limit start从0开始，end为闭区间，即是[start,end)
func (a *ArticleModel) FindArticleOrderByPublishTime(userId int, start int, end int) (articles *[]Article)  {
	var findArticles []Article
	res := DB.Raw(
		// 按照 publish_time 逆排序
		"select * from article where where owner_id = ? and deleted = false order by publish_time desc limit ?,?;",
		userId, start, end).Scan(&findArticles)

	if res.Error != nil {
		mylog.Error.Printf("execute sql error, message: %v", res.Error.Error())
		return
	}

	if res.RowsAffected == 0 {
		return
	}

	articles = &findArticles
	return
}

func (a *ArticleModel) GetCountOfArticle() (n int)  {

	res := DB.Raw(
		"select count(*) as n from article where deleted = false and private = false;",
		).Scan(&n)

	if res.Error != nil {
		mylog.Error.Printf("execute sql error, message: %v", res.Error.Error())
		return
	}
	return
}

// FindLatestArticle 获取最近更新的文章，start和end用于分页
func  (a *ArticleModel) FindLatestArticle(start int, end int) (articles *[]Article)  {
	var findArticles []entity.Article
	res := database.MysqlDB.Raw(
		"select * from article where deleted = false and private = false order by publish_time desc limit ?,?;",
		start, end).Scan(&findArticles)

	if res.Error != nil {
		mylog.Error.Printf("execute sql error, message: %v", res.Error.Error())
		return
	}

	if res.RowsAffected == 0 {
		return
	}

	articles = &findArticles
	return
}