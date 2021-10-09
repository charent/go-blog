package entity

/*
本段代码由工具go-mygen生成，自己再修改（该工具会把bool生成为int32）
go get -u -x github.com/yezihack/go-mygen
go-mygen -h localhost -P 3306 -u root -p root -d go_blog
但是工具生成CURD的代码使用Golang的database/sql来执行，取出的结果要自己拼接为对象，代码可读性差，
而且功能缺失（如：文章按发布时间排序、分页等，

也尝试过：https://github.com/xxjwxc/gormt 大佬自动生成CURD代码
该工具生成的代码虽然使用gorm框架执行，但却是使用ROM框架执行（ORM最后还得转换为原生SQL执行），效率可能受影响（特别是性能差的服务器）

最后决定自己使用gorm框架执行原生sql语句，兼顾性能和可读性

 */

// ArticleLabels 文章对应标签表
type ArticleLabels struct {
	AlId      int `gorm:"al_id"`      // 标签id 主键
	ArticleId int `gorm:"article_id"` // 文章id
	LabelId   int `gorm:"label_id"`   // 标签id

}

// Articles 文章表
type Articles struct {
	Abstract       string `gorm:"abstract"`    // 摘要
	ArticleId      int  `gorm:"article_id"`  // 文章id
	CategoryId     int  `gorm:"category_id"` // 文章分类id
	Deleted        bool  `gorm:"deleted"`		// 是否删除
	LastUpdateTime string `gorm:"last_update_time"` // 上次更新时间
	MdPath         string `gorm:"mdPath"`           // markdown文件路径
	OwnerId        int  `gorm:"owner_id"`         // 文章拥有者id
	Private		   bool  `gorm:"private"`		// 是否是私有文章
	PublishTime    string `gorm:"publish_time"`     // 发布时间
	Tittle         string `gorm:"tittle"`           // 标题
	Visited        int  `gorm:"visited"`          // 访客数量

}

// CategoriesFirst 文章分类表
type CategoriesFirst struct {
	CategoryName string `gorm:"category_name"` // 分类名字
	CfId         int  `gorm:"cf_id"`         // 一级分类id

}

// CategoriesSecond 文章分类表
type CategoriesSecond struct {
	CategoryName string `gorm:"category_name"` // 分类名字
	CsId         int  `gorm:"cs_id"`         // 分类id
	FirstId      int  `gorm:"first_id"`      // 一级分类Id

}

// Labels 标签表
type Labels struct {
	LabelId   int  `gorm:"label_id"`   // 标签id
	LabelName string `gorm:"label_name"` // 标签名字

}

// Operations 操作表
type Operations struct {
	OpId     int  `gorm:"op_id"`     // 操作id
	OpName   string `gorm:"op_name"`   // 操作名字
	OpNameZh string `gorm:"opName_zh"` // 操作名字中文

}

// RoleOperation 角色操作表
type RoleOperation struct {
	OpId   int `gorm:"op_id"`   // 操作id
	RoId   int `gorm:"ro_id"`   // id
	RoleId int `gorm:"role_id"` // 角色id

}

// Roles 角色表
type Roles struct {
	Comments string `gorm:"comments"`  // 备注
	RoleId   int  `gorm:"role_id"`   // 角色id
	RoleName string `gorm:"role_name"` // 角色名字

}

// Users 用户表
type Users struct {
	Deleted       bool  `gorm:"deleted"`         // 是否删除
	LastLoginIp   string `gorm:"last_login_ip"`   // 上次登录IP
	LastLoginTime string `gorm:"last_login_time"` // 上次登录时间
	PasswordHash  string `gorm:"password_hash"`   // 密码
	RoleId        int  `gorm:"role_id"`         // 角色id
	Salted        string `gorm:"salted"`          // 盐值
	UserId        int  `gorm:"user_id"`         // 用户id
	UserName      string `gorm:"user_name"`       // 用户名

}
