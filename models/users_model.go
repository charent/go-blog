package models

import (
	"go-blog/database"
	"go-blog/models/entity"
)

// Users 将User 作为entity.Users的别名
type Users entity.Users

// Insert 插入用户
func (user *Users) Insert() (rows int, err error){
	res := database.MysqlDB.Create(user)
	rows = int(res.RowsAffected)
	if res.Error != nil{
		err = res.Error
		return
	}
	return
}

// FindUserByName 查找用户
func (user *Users) FindUserByName() (rows int, err error)  {
	res := database.MysqlDB.Raw("select * from users where user_name = ? and deleted = false limit 1;", user.UserName).Scan(user)
	rows = int(res.RowsAffected)
	if res.Error != nil{
		err = res.Error
		return
	}
	return
}

// 删除用户
//func (user *User) delete() (id int, err error)  {
//	res := database.MysqlDB.Update()
//}



