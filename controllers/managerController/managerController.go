package managerController

import (
	"github.com/gin-gonic/gin"
	"go-blog/services"
	"net/http"
)

func ManagerHome(c *gin.Context) {
	var userService services.LoginUser

	//解析出请求携带的json
	err := c.ShouldBindJSON(&userService)
	if err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{"success": false,"message": err.Error()})
		return
	}
	authUser := userService.UserLogin()

	if authUser != nil {
		// 登录成功
		c.JSONP(http.StatusOK,gin.H{"success": true})
	}else {
		// 登录失败
		c.JSONP(http.StatusUnauthorized,gin.H{"success": false, "message": "用户名或者密码错误"})
	}

}
