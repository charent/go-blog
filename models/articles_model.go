package models

import (
	"go-blog/database"
	"go-blog/models/entity"
)

type Articles entity.Articles

// FindArticlesOrderByPublishTime 按照时间逆序查找文章
// limit start从0开始，end为闭区间，即是[start,end)
func (article *Articles) FindArticlesOrderByPublishTime(userId int, start int, end int) (rows int, err error)  {
	res := database.MysqlDB.Raw(
		// 按照 publish_time 逆排序
		"select * from articles where where owner_id = ? and deleted = false order by publish_time desc limit ?,?;",
		userId, start, end).Scan(article)
	rows = int(res.RowsAffected)
	if res.Error != nil {
		rows = 0
		err = res.Error
		return
	}
	return
}

// GetLatestArticles 获取最近更新的文章，start和end用于分页
func (article *Articles) GetLatestArticles(start int, end int) (rows int, err error)  {
	res := database.MysqlDB.Raw(
		"select * from articles where deleted = false and private = false order by publish_time desc limit ?,?;",
		start, end).Scan(article)
	rows = int(res.RowsAffected)
	if res.Error != nil {
		rows = 0
		err = res.Error
		return
	}
	return
}