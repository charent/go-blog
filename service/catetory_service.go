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
	CategoryName string
	ArticleCount int
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
		tempCate.CategoryName = first.CategoryName
		tempCate.ArticleCount = first.ArticleCount
		tempList = append(tempList, tempCate)
	}

	category = &tempList
	return
}
