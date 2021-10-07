package entity

import (
	"go-blog/database"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Id          int
	OwnerId    int
	CategoryId int
	Tittle     string
	Abstract     string
	MdPath           string
	PublishTime    string
	LastUpdateTime string
	Visited        int
	Deleted         	bool
}

//FindArticlesOrderByPublishTime 按照时间逆序查找文章
//limit start从0开始，end为闭区间，即是[start,end)
func (article *Article) FindArticlesOrderByPublishTime(userId int, start int, end int) (rows int, err error)  {
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