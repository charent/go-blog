package controller

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"go-blog/service"
	"net/http"
	"strconv"
)

type ManagerController struct {

}

var operationAuth service.OperationAuth
var categoryService service.CategoryService

type putCategoryFirstJson struct {
	CategoryName	string	`json:"categoryName" binding:"required"`
}

type putCategorySecondJson struct {
	CategoryFirstName	string	`json:"categoryFirstName" binding:"required"`
	CategorySecondName	string	`json:"categorySecondName" binding:"required"`

}

// GetCategoryFirst 获取用户的一级分类
func (m *ManagerController) GetCategoryFirst(c *gin.Context)  {
	userId := extractUserId(c)

	categoryFirst := categoryService.GetFirstCategoryFirstByUserId(userId)

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"category": categoryFirst,
	})

}

// GetCategorySecond 获取用户的二级分类，同时二级分类包含一级分类的信息
func (m *ManagerController) GetCategorySecond(c *gin.Context)  {
	userId := extractUserId(c)
	category := categoryService.GetCategoryFirstJoinSecondByUserId(userId)

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"category": category,
	})
}


// PutCategorySecond 插入二级分类
func (m *ManagerController) PutCategorySecond(c *gin.Context)  {
	userId := extractUserId(c)
	var reqJson putCategorySecondJson
	err := c.ShouldBindJSON(&reqJson)

	if err != nil {
		PublicHandler.ParamError(c, err.Error(), nil)
		return
	}

	rowsAffected := categoryService.InsertCategorySecond(userId, reqJson.CategoryFirstName, reqJson.CategorySecondName)

	// rowsAffected为0因为没有插入成功，表示数据库中该用户已经有该分类了，正常插入rowsAffected=1，
	if rowsAffected ==  0 {
		PublicHandler.ParamError(c, "一级分类不存在或者二级分类已经存在", reqJson)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"categoryName": reqJson.CategorySecondName,
	})
}

// PutCategoryFirst 插入一级分类
func (m *ManagerController) PutCategoryFirst(c *gin.Context)  {
	userId := extractUserId(c)
	var reqJson putCategoryFirstJson
	err := c.ShouldBindJSON(&reqJson)

	if err != nil {
		PublicHandler.ParamError(c, err.Error(), nil)
		return
	}

	rowsAffected := categoryService.InsertCategoryFirst(userId, reqJson.CategoryName)

	// rowsAffected为0因为没有插入成功，表示数据库中该用户已经有该分类了，正常插入rowsAffected=1，
	if rowsAffected ==  0 {
		PublicHandler.ParamError(c, "一级分类已经存在", reqJson.CategoryName)
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"code": http.StatusOK,
		"categoryName": reqJson.CategoryName,
	})
}


func (m *ManagerController) Home(c *gin.Context) {


}


func (m *ManagerController) PutArticle(c *gin.Context)  {

}

// getUserId 从认证token中去除用户id
func extractUserId(c *gin.Context) (userId int) {
	claims := jwt.ExtractClaims(c)

	// 先转string再转int
	uid, _ := claims["userId"].(string)
	userId, _ = strconv.Atoi(uid)
	return
}
