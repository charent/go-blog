package service

import (
	"go-blog/model"
	"go-blog/model/entity"
	"go-blog/utils/mylog"
	"golang.org/x/crypto/bcrypt"
)

// User 注意这里不要写=
// type xx = XX 是类型别名，type xx XX 是类型定义
type User = entity.BlogUser

// LoginUser 登录接口要post的用户名和密码
type LoginUser struct {
	UserName        string 	`json:"userName" binding:"required"`
	Password        string 	`json:"password" binding:"required"`
}

type UserService struct {

}
var UserModel model.UserModel

func (u *UserService)UserLogin(loginUser *LoginUser, ip string) (authUser *User) {

	// 根据用户名查找用户
	user, err := UserModel.FindUserByName(loginUser.UserName)

	if err != nil {
		return
	}

	//将用户输入的密码和数据库的盐值相加
	pwdSalted := loginUser.Password + user.Salted
	//hashedPwd := HashAndSalt(pwdSalted)
	//print(hashedPwd+"\n")

	// 将数据库的密文和盐值相加的密码对比，对比一致
	if ValidatePasswords(user.PasswordHash, pwdSalted) {
		authUser = (*User)(user)

		// 更新登录时间和ip
		_ = UserModel.UpdateLoginInfo(user.UserId, ip)
	}
	return
}

// HashAndSalt 加密密码
func HashAndSalt(password string) string {
	hash, err := bcrypt.GenerateFromPassword([] byte(password), bcrypt.DefaultCost)
	if err != nil {
		mylog.Error.Printf("hash error, message: %v", err)
	}
	return string(hash)
}

// ValidatePasswords 验证密码
func ValidatePasswords(hashedPwd string, databasePwd string) bool {
	byteHash := []byte(hashedPwd)
	bytePwd := []byte(databasePwd)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePwd)
	if err != nil {
		return false
	}
	return true
}