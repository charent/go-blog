package entity

/*
本段代码由工具go-mygen生成，自己再修改（该工具会把bool生成为int）
go get -u -x github.com/yezihack/go-mygen
go-mygen -h localhost -P 3306 -u root -p root -d go_blog
但是工具生成CURD的代码使用Golang的database/sql来执行，取出的结果要自己拼接为对象，代码可读性差，
而且功能缺失（如：文章按发布时间排序、分页等，

也尝试过：https://github.com/xxjwxc/gormt 大佬自动生成CURD代码
该工具生成的代码虽然使用gorm框架执行，但却是使用ROM框架执行（ORM最后还得转换为原生SQL执行），效率可能受影响（特别是性能差的服务器）

最后决定自己使用gorm框架执行原生sql语句，兼顾性能和可读性

 */

// Article 文章表
type Article struct {
	Abstract       string `gorm:"abstract"`         // 摘要
	ArticleId      int    `gorm:"article_id;primaryKey"`       // 文章id
	CategoryId     int    `gorm:"category_id"`      // 文章分类id
	Deleted        int    `gorm:"deleted"`          // 是否删除
	LastUpdateTime string `gorm:"last_update_time"` // 上次更新时间
	OwnerId        int    `gorm:"owner_id"`         // 文章拥有者id
	Private        int    `gorm:"private"`          // 是否是私有文章
	PublishTime    string `gorm:"publish_time"`     // 发布时间
	Title          string `gorm:"title"`            // 标题
	Visited        int    `gorm:"visited"`          // 访客数量

}
func (a *Article) TableName() string {
	// 映射表名，否则会在表名后面+s => articles
	return "article"
}

// ArticleLabel 文章对应标签表
type ArticleLabel struct {
	AlId      int   `gorm:"al_id;primaryKey"`      // 标签id 主键
	ArticleId int   `gorm:"article_id"` // 文章id
	LabelId   int   `gorm:"label_id"`   // 标签id

}
func (a *ArticleLabel) TableName() string {
	return "article_label"
}



// BlogRole 角色表
type BlogRole struct {
	Comments string `gorm:"comments"`  // 备注
	RoleId   int    `gorm:"role_id;primaryKey"`   // 角色id
	RoleName string `gorm:"role_name"` // 角色名字

}
func (b *BlogRole) TableName() string {
	return "blog_role"
}



// BlogUser 用户表
type BlogUser struct {
	Deleted       int    `gorm:"deleted"`         // 是否删除
	LastLoginIp   string `gorm:"last_login_ip"`   // 上次登录IP
	LastLoginTime string `gorm:"last_login_time"` // 上次登录时间
	PasswordHash  string `gorm:"password_hash"`   // 密码
	RoleId        int    `gorm:"role_id"`         // 角色id
	Salted        string `gorm:"salted"`          // 盐值
	UserId        int    `gorm:"user_id;primaryKey"`         // 用户id
	UserName      string `gorm:"user_name"`       // 用户名

}
func (b *BlogUser) TableName() string {
	return "blog_user"
}

// CategoryFirst 一级文章分类表
type CategoryFirst struct {
	ArticleCount int	`gorm:"article_count"` // 当前分类的文章数
	CategoryName string `gorm:"category_name"` // 分类名字
	CfId         int    `gorm:"cf_id;primaryKey"`         // 一级分类id
	OwnerId      int    `gorm:"owner_id"`      // 一级分类拥有者id

}
func (c *CategoryFirst) TableName() string {
	return "category_first"
}


// CategorySecond 二级文章分类表
type CategorySecond struct {
	ArticleCount int	`gorm:"article_count"` // 当前分类的文章数
	CategoryName string `gorm:"category_name"` // 分类名字
	CsId         int    `gorm:"cs_id;primaryKey"`         // 分类id
	FirstId      int    `gorm:"first_id"`      // 一级分类Id
	OwnerId      int    `gorm:"owner_id"`      // 二级级分类拥有者id

}
func (c *CategorySecond) TableName() string {
	return "category_second"
}

// Draft 用户文章草稿表
type Draft struct {
	OwnerId      int32  `gorm:"owner_id;primaryKey"`       // 该草稿的拥有者，和userId对应
	Title        string `gorm:"title"`          // 标题
	Abstract     string `gorm:"abstract"`       // 摘要
	Content      string `gorm:"content"`        // markdown草稿文件内容
	LastSaveTime string `gorm:"last_save_time"` // 上次保存草稿的时间
	Deleted      int32  `gorm:"deleted"`        // 是否删除

}
func (d *Draft) TableName() string {
	return "draft"
}


// Label 标签表
type Label struct {
	LabelId   int    `gorm:"label_id;primaryKey"`   // 标签id
	LabelName string `gorm:"label_name"` // 标签名字

}
func (l *Label) TableName() string {
	return "label"
}


// Markdown markdown存储表
type Markdown struct {
	ArticleId int    `gorm:"article_id;primaryKey"` // 文章id，和article表的文章id对应
	Content   string `gorm:"content"`    // markdown文件内容
	Deleted   int    `gorm:"deleted"`    // 是否删除

}
func (m *Markdown) TableName() string {
	return "markdown"
}


// Operation 操作表
type Operation struct {
	OpId     int    `gorm:"op_id;primaryKey"`     // 操作id
	OpName   string `gorm:"op_name"`   // 操作名字
	OpNameZh string `gorm:"opName_zh"` // 操作名字中文

}
func (o *Operation) TableName() string {
	return "operation"
}


// RoleOperation 角色操作表
type RoleOperation struct {
	OpId   int   `gorm:"op_id;primaryKey"`   // 操作id
	RoId   int   `gorm:"ro_id"`   // id
	RoleId int   `gorm:"role_id"` // 角色id

}
func (r *RoleOperation) TableName() string {
	return "role_operation"
}

