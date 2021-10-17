package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type myHandler struct {}

// PublicHandler 在controller包可以直接使用
var PublicHandler *myHandler

func init()  {
	var handler myHandler
	PublicHandler = &handler
}

// ParamError 参数错误
func (p *myHandler) ParamError(c *gin.Context, message string, param interface{})  {
	c.JSON(http.StatusBadRequest, gin.H{
		"code": http.StatusBadRequest,
		"message": message,
		"error_param": param,
	})
}

// PageNotFound 404找不到页面返回json
func (p *myHandler) PageNotFound(c *gin.Context)  {
	c.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "message": "Page not found"})
}
