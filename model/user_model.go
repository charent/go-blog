package model

import (
	"go-blog/database"
	"go-blog/model/entity"
	"go-blog/utils/mylog"
	"time"
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

func (u *UserModel) UpdateLoginInfo(uerId int, ip string) (RowsAffected int)  {

	res := DB.Exec(
		"update blog_user set last_login_time = ?, last_login_ip = ? where user_id = ?;",
		time.Now().Format("2006-01-02 15:04:05"), ip, uerId,
	)
	if res.Error != nil {
		mylog.Error.Printf("execute sql error, message: %v", res.Error.Error())
		return
	}

	RowsAffected = int(res.RowsAffected)
	return
}



