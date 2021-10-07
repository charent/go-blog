package entity

import "gorm.io/gorm"

/*
本段代码由 go-mygen 工具生成，参考：https://github.com/yezihack/go-mygen
但是该工具生成CURD的代码使用Golang的database/sql来执行，取出的结果要自己拼接为对象，代码可读性差，
而且功能缺失（如：文章按发布时间排序、分页等，

也尝试过：https://github.com/xxjwxc/gormt 大佬自动生成CURD代码
该工具生成的代码虽然使用gorm框架执行，但却是使用ROM框架执行（ORM最后还得转换为原生SQL执行），效率可能受影响（特别是性能差的服务器）

最后决定自己使用gorm框架执行原生sql语句，兼顾性能和可读性

 */

// ArticleLabels 文章对应标签表
type ArticleLabels struct {
	gorm.Model
	ArticleId int32 // 文章id
	Id        int32 // 标签id
	LabelId   int32 // 标签id

}

// Articles 文章表
type Articles struct {
	gorm.Model
	Abstract       string // 摘要
	CategoryId     int32  // 文章分类id
	Deleted        int32
	Id             int32  // 文章id
	LastUpdateTime string // 上次更新时间
	MdPath         string // markdown文件路径
	OwnerId        int32  // 文章拥有者id
	PublishTime    string // 发布时间
	Tittle         string // 标题
	Visited        int32  // 访客数量

}

// CategoriesFirst 文章分类表
type CategoriesFirst struct {
	gorm.Model
	CategoryName string // 分类名字
	Id           int32  // 分类id

}

// CategoriesSecond 文章分类表
type CategoriesSecond struct {
	gorm.Model
	CategoryName string // 分类名字
	FirstId      int32  // 一级分类Id
	Id           int32  // 分类id

}

// Labels 标签表
type Labels struct {
	gorm.Model
	Id        int32  // 标签id
	LabelName string // 标签名字

}

// Operations 操作表
type Operations struct {
	gorm.Model
	OpId     int32  // 操作id
	OpName   string // 操作名字
	OpNameZH string // 操作名字中文

}

// RoleOperation 角色操作表
type RoleOperation struct {
	gorm.Model
	Id     int32 // id
	OpId   int32 // 操作id
	RoleId int32 // 角色id

}

// Roles 角色表
type Roles struct {
	gorm.Model
	Comments string // 备注
	RoleId   int32  // 角色id
	RoleName string // 角色名字

}

// Users 用户表
type Users struct {
	gorm.Model
	Deleted       int32  // 是否删除
	LastLoginIP   string // 上次登录IP
	LastLoginTime string // 上次登录时间
	Name          string // 用户名
	Password      string // 密码
	RoleId        int32  // 角色id
	Salted        string // 盐值
	UserId        int32  // 用户id
}
