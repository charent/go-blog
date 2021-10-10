package middleware

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"go-blog/config"
	"go-blog/service"
	"go-blog/utils/mylog"
	"time"
)

// AuthMiddleware API认证的中间件
var AuthMiddleware *jwt.GinJWTMiddleware

var identityKey = "userId"

func init()  {
	var err error
	AuthMiddleware, err =  jwt.New(&jwt.GinJWTMiddleware{
		// 名字，后台管理的API才需要认证
		Realm: "manager zone",
		Key: []byte(config.Server.Api.ApiKey),

		// token过期时间
		Timeout: time.Minute * 30,
		MaxRefresh: time.Minute * 15,
		PayloadFunc: addJwtPayload,  	// 登录成功后生成返回token时会调用该函数，在token的payload字段中添加用户id
		Authenticator: jwtLoginAuth, 	// 登录时调用该函数，账户密码验证通过后会返回一个jwt token
		Unauthorized: jwtUnAuth,		// 认证不通过时会调用该函数
	})

	if err != nil{
		mylog.Error.Print("jwt AuthMiddleware init error, message: %v", err)
	}
}

// 认证通过的User，UserId作为其唯一标记，该标记将放在jwt token中
type authUser struct {
	UserId int
}
var UserService service.UserService
// 登录认证
func jwtLoginAuth(c *gin.Context) (interface{}, error)  {

	// 登录接口post上来的用户名和密码 struct
	var loginUser service.LoginUser

	//解析出请求携带的json
	err := c.ShouldBindJSON(&loginUser)
	if err != nil {
		//c.JSONP(http.StatusBadRequest, gin.H{"success": false,"message": err.Error()})
		return nil, jwt.ErrFailedAuthentication
	}
	user := UserService.UserLogin(&loginUser)

	// 登录成功
	if user != nil{
		return &authUser{
			UserId: user.UserId,
		}, nil
	}

	// 认证失败
	return nil, jwt.ErrFailedAuthentication
}

// 为clams添加用户标记userId，post请求携带的token可识别出是哪个用户
func addJwtPayload(data interface{}) jwt.MapClaims {
	if v, ok := data.(*authUser); ok {
		return jwt.MapClaims{
			identityKey: v.UserId,
		}
	}
	return jwt.MapClaims{}
}

// API认证
func apiAuh(c *gin.Context)   {

}

// 认证不通过
func jwtUnAuth(c *gin.Context, code int, message string)  {
	c.JSON(code, gin.H{
		"success": false,
		"code": code,
		"message": message,
	})
}



