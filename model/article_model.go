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

// FindArticleOrderByPublishTime 按照时间逆序查找文章
// limit start从0开始，end为闭区间，即是[start,end)
func (a *ArticleModel) FindArticleOrderByPublishTime(userId int, start int, end int) (articles *[]Article)  {
	res := DB.Raw(
		// 按照 publish_time 逆排序
		"select * from article where where owner_id = ? and deleted = false order by publish_time desc limit ?,?;",
		userId, start, end).Scan(articles)

	if res.Error != nil {
		mylog.Error.Printf("execute sql error, message: %v", res.Error.Error())
		return
	}
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

// GetLatestArticle 获取最近更新的文章，start和end用于分页
func  (a *ArticleModel) GetLatestArticle(start int, end int) (articles *[]Article)  {
	var findArticles []entity.Article
	res := database.MysqlDB.Raw(
		"select * from article where deleted = false and private = false order by publish_time desc limit ?,?;",
		start, end).Scan(&findArticles)

	if res.Error != nil {
		mylog.Error.Printf("execute sql error, message: %v", res.Error.Error())
		return
	}
	articles = &findArticles
	return
}