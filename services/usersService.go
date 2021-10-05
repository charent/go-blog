package services

import (
	"go-blog/models"
	"go-blog/utils/log"
	"golang.org/x/crypto/bcrypt"
)

type LoginUser struct {
	Name            string 	`json:"userName" binding:"required"`
	Password        string 	`json:"password" binding:"required"`
}

func (user *LoginUser) UserLogin() (authUser *models.User) {
	authUser = nil
	var userModel models.User
	userModel.Name = user.Name

	// 根据用户名查找用户
	_, err := userModel.FindUserByName()

	if err != nil {
		log.Error.Printf("adminController login error, message: %v", err)
		return
	}

	//将用户输入的密码和数据库的盐值相加
	pwdSalted := user.Password + userModel.Salted
	//hashedPwd := HashAndSalt(pwdSalted)
	//print(hashedPwd+"\n")

	// 将数据库的密文和盐值相加的密码对比
	if ValidatePasswords(userModel.Password, pwdSalted) {
		authUser = &userModel
	}
	return
}

// HashAndSalt 加密密码
func HashAndSalt(password string) string {
	hash, err := bcrypt.GenerateFromPassword([] byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Error.Printf("hash error, message: %v", err)
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