package middlewares

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"go-blog/config"
	"go-blog/services"
	"go-blog/utils/log"
	"time"
)

// AuthMiddleware API认证的中间件
var AuthMiddleware *jwt.GinJWTMiddleware

func init()  {
	var err error
	AuthMiddleware, err =  jwt.New(&jwt.GinJWTMiddleware{
		// 名字，后台管理的API才需要认证
		Realm: "manager zone",
		Key: []byte(config.Server.Api.ApiKey),

		// token过期时间
		Timeout: time.Minute * 30,
		MaxRefresh: time.Minute * 15,
		Authenticator: jwtLoginAuth,
		Unauthorized: jwtUnAuth,
	})

	if err != nil{
		log.Error.Print("jwt AuthMiddleware init error, message: %v", err)
	}
}

type authUser struct {
	UserId int
	UserName string
	RoleId int
}

// 登录认证
func jwtLoginAuth(c *gin.Context) (interface{}, error)  {
	var userService services.LoginUser

	//解析出请求携带的json
	err := c.ShouldBindJSON(&userService)
	if err != nil {
		//c.JSONP(http.StatusBadRequest, gin.H{"success": false,"message": err.Error()})
		return nil, jwt.ErrFailedAuthentication
	}
	user := userService.UserLogin()

	// 登录成功
	if user != nil{

		return &authUser{
			UserId: user.User_id,
			UserName: user.Name,
			RoleId: user.Role_id,
		}, nil
	}

	// 认证失败
	return nil, jwt.ErrFailedAuthentication
}

// 认证不通过
func jwtUnAuth(c *gin.Context, code int, message string)  {
	c.JSON(code, gin.H{
		"success": false,
		"code": code,
		"message": message,
	})
}



