package controller

import (
	"fmt"
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

type renameCategoryJson struct {
	CategoryId		int 		`json:"categoryId" binding:"required"`
	NewName			string		`json:"newName" binding:"required"`
}

// RenameFirstCategory 重命名一级分类
func (m *ManagerController) RenameFirstCategory(c *gin.Context) {
	userId := extractUserId(c)

	var reqJson renameCategoryJson
	err := c.ShouldBindJSON(&reqJson)

	if err != nil {
		PublicHandler.ParamError(c, err.Error(), nil)
		return
	}

	rowsAffected := categoryService.RenameFirstCategoryName(userId, reqJson.CategoryId, reqJson.NewName)

	if rowsAffected != 1 {
		PublicHandler.ParamError(c, "重命名一级分类失败，请检查参数", reqJson)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"category": reqJson,
	})
}

// RenameSecondCategory 重命名二级分类
func (m *ManagerController) RenameSecondCategory(c *gin.Context) {
	userId := extractUserId(c)

	var reqJson renameCategoryJson
	err := c.ShouldBindJSON(&reqJson)

	if err != nil {
		PublicHandler.ParamError(c, err.Error(), nil)
		return
	}

	rowsAffected := categoryService.RenameSecondCategoryName(userId, reqJson.CategoryId, reqJson.NewName)

	if rowsAffected != 1 {
		PublicHandler.ParamError(c, "重命名二级分类失败，请检查参数", reqJson)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"category": reqJson,
	})
}

type deleteJson struct {
	Id		int		`json:"id" binding:"required"`
}

// DeleteFirstCategory 删除一级分类
func (m *ManagerController) DeleteFirstCategory(c *gin.Context) {
	userId := extractUserId(c)

	var reqJson deleteJson
	err := c.ShouldBindJSON(&reqJson)

	if err != nil {
		PublicHandler.ParamError(c, err.Error(), nil)
		return
	}
	rowsAffected := categoryService.DeleteFirstCategory(userId, reqJson.Id)

	if rowsAffected != 1 {
		PublicHandler.ParamError(c, "删除失败，该分类下的文章数可能不为0，请检查参数", reqJson)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"category": reqJson,
	})

}

// DeleteSecondCategory 删除二级分类
func (m *ManagerController) DeleteSecondCategory(c *gin.Context) {
	userId := extractUserId(c)

	var reqJson deleteJson
	err := c.ShouldBindJSON(&reqJson)

	if err != nil {
		PublicHandler.ParamError(c, err.Error(), nil)
		return
	}
	rowsAffected := categoryService.DeleteSecondCategory(userId, reqJson.Id)

	if rowsAffected != 1 {
		PublicHandler.ParamError(c, "删除失败，该分类下的文章数可能不为0，请检查参数", reqJson)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"category": reqJson,
	})
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
		"categoryFirstName":reqJson.CategoryFirstName,
		"categorySecondName": reqJson.CategorySecondName,
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
		PublicHandler.ParamError(c, "一级分类 " + reqJson.CategoryName + " 已经存在", reqJson.CategoryName)
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
	userId := extractUserId(c)
	var reqJson service.PublishArticleJson
	err := c.ShouldBindJSON(&reqJson)
	fmt.Printf("%v", err)

	if err != nil {
		PublicHandler.ParamError(c, err.Error(), nil)
		return
	}

	articleId := articleService.InsertArticle(userId, &reqJson)

	if articleId <= 0 {
		PublicHandler.ParamError(c, "插入失败", nil)
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"code": http.StatusOK,
		"articleId":articleId,
		"message": "发布成功",
	})


}

// getUserId 从认证token中去除用户id
func extractUserId(c *gin.Context) (userId int) {
	claims := jwt.ExtractClaims(c)

	// 先转string再转int
	uid, _ := claims["userId"].(string)
	userId, _ = strconv.Atoi(uid)
	return
}
