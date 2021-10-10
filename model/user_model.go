package model

import (
	"go-blog/database"
	"go-blog/model/entity"
)

// User 将User 作为entity.Users的别名
type User = entity.BlogUser

type UserModel struct {

}

// Insert 插入用户
func (u *UserModel)Insert(user *User) (rows int, err error){
	res := database.MysqlDB.Create(user)
	rows = int(res.RowsAffected)
	if res.Error != nil{
		err = res.Error
		return
	}
	return
}

// FindUserByName 查找用户
func (u *UserModel) FindUserByName(userName string) (user *User, err error)  {
	var findUser User
	res := database.MysqlDB.Raw("select * from blog_user where user_name = ? and deleted = false limit 1;", userName).Scan(&findUser)
	if res.Error != nil{
		err = res.Error
		return
	}
	user = &findUser
	return
}

// 删除用户



