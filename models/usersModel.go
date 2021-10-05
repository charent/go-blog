package models

import (
	"go-blog/database"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	User_id         int
	Name            string
	Password        string
	Salted          string
	Role_id         int
	Last_login_time string
	Last_login_ip   string
	Deleted         bool
}

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



