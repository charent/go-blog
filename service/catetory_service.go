package service

import (
	"go-blog/model"
	"go-blog/model/entity"
)

type CategoryService struct {

}

var categoryFirstModel  model.CategoryFirstModel
var categorySecondModel  model.CategorySecondModel

// InsertCategorySecond 插入二级分类
func (c *CategoryService) InsertCategorySecond(userId int, cFirstName string, cSecondName string) (rowsAffected int) {
	cfId := categoryFirstModel.FindCfIdByCategoryName(userId, cFirstName)

	// cfId = 0表示数据库中没有该一级分类
	if cfId == 0 { return }

	var category entity.CategorySecond
	category.OwnerId = userId
	category.FirstId = cfId
	category.CategoryName = cSecondName

	rowsAffected = categorySecondModel.InsertCategorySecond(&category)
	return
}

// InsertCategoryFirst 插入一个新的一级分类
func (c *CategoryService ) InsertCategoryFirst(userId int, categoryName string) (rowsAffected int)  {
	var category entity.CategoryFirst
	category.OwnerId = userId
	category.CategoryName = categoryName

	rowsAffected = categoryFirstModel.InsertCategoryFirst(&category)
	return
}

// GetCategoryFirstJoinSecondByUserId 获取用户的一级分类和二级分类
func (c *CategoryService) GetCategoryFirstJoinSecondByUserId(userId int)  (category *[]model.CategoryFirstSecond){
	var tempList = make([]model.CategoryFirstSecond, 0)
	findCategory := categorySecondModel.FindCategoryFirstJoinSecond(userId)

	if findCategory == nil {
		category = &tempList
		return
	}
	category = findCategory
	return
}

type GetCategoryFirstJson struct {
	CfId			int
	CategoryName 	string
	ArticleCount	int
}

// GetFirstCategoryFirstByUserId 获取用户的一级分类
func (c *CategoryService) GetFirstCategoryFirstByUserId(userId int) (category *[]GetCategoryFirstJson) {
	var tempList = make([]GetCategoryFirstJson, 0)
	findCategory := categoryFirstModel.FindCategoryFirstByUserId(userId)
	if findCategory == nil {
		category = &tempList
	}


	// 过滤掉不需要的字段
	for _, first := range *findCategory {
		var tempCate GetCategoryFirstJson
		tempCate.CfId = first.CfId
		tempCate.CategoryName = first.CategoryName
		tempCate.ArticleCount = first.ArticleCount
		tempList = append(tempList, tempCate)
	}

	category = &tempList
	return
}

// RenameFirstCategoryName 重命名一级分类
func (c *CategoryService) RenameFirstCategoryName(userId int, cfId int, newName string) (rowsAffected int)  {
	rowsAffected = categoryFirstModel.UpdateCategoryName(userId, cfId, newName)
	return
}

// RenameSecondCategoryName 重命名二级分类
func (c *CategoryService) RenameSecondCategoryName(userId int, csId int, newName string) (rowsAffected int)  {
	rowsAffected = categorySecondModel.UpdateCategoryName(userId, csId, newName)
	return
}

// DeleteFirstCategory 删除一级分类
func (c *CategoryService) DeleteFirstCategory(userId int, cfId int) (rowsAffected int)  {
	rowsAffected = categoryFirstModel.DeleteCategory(userId, cfId)
	return
}

// DeleteSecondCategory 删除二级分类
func (c *CategoryService) DeleteSecondCategory(userId int, csId int) (rowsAffected int)  {
	rowsAffected = categoryFirstModel.DeleteCategory(userId, csId)
	return
}



