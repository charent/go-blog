package entity

import (
	"go-blog/database"
	"go-blog/models/entity"
)

// User 将Users表改为User定义
type User entity.Users

// Insert 插入用户
func (user *User) Insert() (rows int, err error){
	res := database.MysqlDB.Create(user)
	rows = int(res.RowsAffected)
	if res.Error != nil{
		err = res.Error
		return
	}
	return
}

// FindUserByName 查找用户
func (user *User) FindUserByName() (rows int, err error)  {
	res := database.MysqlDB.Raw("select * from users where name = ? and deleted = false limit 1;", user.Name).Scan(user)
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



