package model

import (
	"go-blog/model/entity"
	"go-blog/utils/mylog"
)

type CategoryFirstModel struct {

}
type CategoryFirst = entity.CategoryFirst

// AddArticleCount 该分类添加了新文章，article_count+1
func (c *CategoryFirstModel ) AddArticleCount(userId int,  cfId int) (rowsAffected int) {
	res := DB.Exec(
		"update category_first set article_count = article_count + 1 " +
			"where owner_id = ? and cf_id = ? ;", userId, cfId,
		)
	if res.Error != nil {
		mylog.Error.Printf("execute sql error, message: %v", res.Error.Error())
		return
	}
	rowsAffected = int(res.RowsAffected)
	return
}

// SubArticleCount 该分类删除了文章，article_count-1
func (c *CategoryFirstModel ) SubArticleCount(userId int,  cfId int) (rowsAffected int) {
	res := DB.Exec(
		"update category_first set article_count = article_count - 1 " +
			"where owner_id = ? and cf_id = ? ;", userId, cfId,
	)
	if res.Error != nil {
		mylog.Error.Printf("execute sql error, message: %v", res.Error.Error())
		return
	}
	rowsAffected = int(res.RowsAffected)
	return
}

// FindCfIdByCategoryName 根据类名找到id
func (c *CategoryFirstModel) FindCfIdByCategoryName(userId int, categoryName string) (cfId int)  {
	res := DB.Raw(
		"select cf_id from category_first where owner_id = ? and category_name = ? limit 1;",
		userId, categoryName).Scan(&cfId)
	if res.Error != nil {
		mylog.Error.Printf("execute sql error, message: %v", res.Error.Error())
		return
	}
	return
}


// InsertCategoryFirst 插入一个一级分类
func (c *CategoryFirstModel) InsertCategoryFirst(category *CategoryFirst) (rowsAffected int) {
	// 检查数据库中，该用户是否已经有了该分类
	var nameCount int
	res := DB.Raw(
		"select count(category_name) as name_count from category_first where owner_id = ? and category_name = ?",
		category.OwnerId, category.CategoryName).Scan(&nameCount)

	// 已经存在该分类了，直接返回
	if nameCount != 0 {
		return
	}

	// 创建新的一级分类
	res = DB.Create(category)

	if res.Error != nil {
		mylog.Error.Printf("execute sql error, message: %v", res.Error.Error())
		return
	}

	rowsAffected = int(res.RowsAffected)
	return
}

// FindCategoryFirstByUserId 根据用户id获取用户文章的一级分类
func (c *CategoryFirstModel) FindCategoryFirstByUserId(userId int) (categories *[]CategoryFirst)  {
	var findCategories []CategoryFirst
	res := DB.Raw(
		"select cf_id, category_name, article_count from category_first where owner_id = ?;", userId).Scan(&findCategories)

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

// UpdateCategoryName 重命名一级分类名字
func (c *CategoryFirstModel) UpdateCategoryName(userId int, cfId int, newName string) (rowsAffected int) {
	res := DB.Exec(
		"update category_first set category_name = ? " +
			"where owner_id = ? and cf_id = ?;", newName, userId, cfId,
	)
	if res.Error != nil {
		mylog.Error.Printf("execute sql error, message: %v", res.Error.Error())
		return
	}
	rowsAffected = int(res.RowsAffected)
	return
}

// DeleteCategory 删除一级分类
func (c *CategoryFirstModel) DeleteCategory(userId int, cfId int,) (rowsAffected int) {

	// 删除的分类要求：该分类下的文章数必须为0
	res := DB.Exec(
		"delete from category_first " +
			"where owner_id = ? and cfId = ? and article_count = 0;", userId, cfId,
	)
	if res.Error != nil {
		mylog.Error.Printf("execute sql error, message: %v", res.Error.Error())
		return
	}

	rowsAffected = int(res.RowsAffected)
	return
}