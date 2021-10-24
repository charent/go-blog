package model

import (
	"go-blog/model/entity"
	"go-blog/utils/mylog"
)

type CategorySecondModel struct {

}

type CategorySecond = entity.CategorySecond

// CategoryFirstSecond 第一分类和第二分类对应
type CategoryFirstSecond struct {
	CategoryFirstName	string
	CategorySecondName 	string
	ArticleCount       	int
}

// InsertCategorySecond 插入二级分类
func (c *CategorySecondModel) InsertCategorySecond(category *entity.CategorySecond) (rowsAffected int)  {

	var nameCount int
	res := DB.Raw(
		"select count(category_name) as name_count from category_second where owner_id = ? and first_id = ? and category_name = ?",
		category.OwnerId, category.FirstId, category.CategoryName).Scan(&nameCount)

	// 已经存在该分类了，直接返回
	if nameCount != 0 {
		return
	}

	// 创建新的二级分类
	res = DB.Create(category)

	if res.Error != nil {
		mylog.Error.Printf("execute sql error, message: %v", res.Error.Error())
		return
	}

	rowsAffected = int(res.RowsAffected)
	return
}


// FindCategoryFirstJoinSecond 找出用户id下的第二分类和第二分类的名字对应
func (c *CategorySecondModel) FindCategoryFirstJoinSecond(userId int) (category *[]CategoryFirstSecond)  {
	var findCategory []CategoryFirstSecond
	res := DB.Raw(
		// 注意有+号每一行后面都要多一个空格
		"select c1.category_name as category_first_name, c2.category_name as category_second_name, c2.article_count " +
			"from category_first as c1 join category_second as c2 on c1.cf_id = c2.first_id " +
			"where c1.owner_id = ? and c2.owner_id = ? order by c1.category_name, c2.category_name;", userId, userId,
		).Scan(&findCategory)

	if res.Error != nil {
		mylog.Error.Printf("execute sql error, message: %v", res.Error.Error())
		return
	}

	if res.RowsAffected == 0 {
		return
	}

	category = &findCategory
	return
}

// FindCategorySecondByUserId 根据用户id获取用户文章的二级分类
func (c *CategorySecondModel) FindCategorySecondByUserId(userId int) (categories *[]CategorySecond)  {
	var findCategories []CategorySecond
	res := DB.Raw(
		"select * from category_second where owner_id = ?;", userId).Scan(&findCategories)

	if res.Error != nil {
		mylog.Error.Printf("execute sql error, message: %v", res.Error.Error())
		return
	}

	if res.RowsAffected == 0 {
		return
	}
	categories = &findCategories
	return
}

// FindIdByNameAndOwnerId 根据二级分类名字和用户id获取该分类的id，（发布文章的时候用）
func (c *CategorySecondModel ) FindIdByNameAndOwnerId(userId int, categoryFirstName string, categorySecondName string) (csId int)  {
	res := DB.Raw(
		"select cs_id from category_second as c2 join category_first as c1 on c1.cf_id = c2.first_id " +
			"where c1.owner_id = ? and c2.owner_id = ? and c1.category_name = ? and c2.category_name = ? ;",
		userId, userId, categoryFirstName, categorySecondName,
		).Scan(&csId)

	if res.Error != nil {
		mylog.Error.Printf("execute sql error, message: %v", res.Error.Error())
		return
	}

	return
}