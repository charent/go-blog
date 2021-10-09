create database if not exists go_blog default char set utf8mb4;
use go_blog;

-- 创建用户表 
create table if not exists users(
    user_id int unsigned primary key auto_increment comment '用户id',
    user_name varchar(16) not null unique comment '用户名',
    password_hash varchar(256) not null comment '密码',
    salted varchar(16) not null comment '盐值',
    role_id int not null comment '角色id',
    last_login_time varchar(32) comment '上次登录时间',
    last_login_ip varchar(32) comment '上次登录IP',
    deleted boolean default false not null comment '是否删除',
    index name_index(user_name) comment '用户名索引'
) engine InnoDB, default char set utf8mb4, comment '用户表';

-- 插入用户admin，账号：admin，密码：admin，插入的密码是加密后的
INSERT INTO `go_blog`.`users` (`user_id`, `user_name`, `password_hash`, `salted`, `role_id`, `last_login_time`, `last_login_ip`, `deleted`)
VALUES (1, 'admin', '$2a$10$kxFrdMf82pEAR8X1Lg88M.U06Z7UwlSSlvDUk90iOfiRA3P7mJ6xi', 'dfj2', '0', '0', '0', '0');

-- drop table users;
-- 创建角色表
create table if not exists roles(
    role_id int unsigned primary key auto_increment comment '角色id',
    role_name varchar(16) not null unique comment '角色名字',
    comments varchar(64) comment '备注'
) engine InnoDB, default char set utf8mb4, comment '角色表';

-- 插入用户admin的角色
INSERT INTO `go_blog`.`roles` (`role_id`, `role_name`, `comments`) VALUES (1, 'admin', '管理员');

-- 创建操作表
create table if not exists operations(
    op_id int unsigned primary key auto_increment comment '操作id',
    op_name varchar(32) not null unique comment '操作名字',
    opName_zh varchar(32) unique comment '操作名字中文'
) engine InnoDB, default char set utf8mb4, comment '操作表';

-- 创建角色操作表（角色拥有哪些操作权限）
create table if not exists role_operation(
    ro_id int unsigned primary key auto_increment comment 'id',
    role_id int unsigned not null comment '角色id',
    op_id int unsigned not null comment '操作id'
) engine InnoDB, default char set utf8mb4, comment '角色操作表';

-- 创建文章一级分类表
create table if not exists categories_first(
    cf_id int unsigned primary key auto_increment comment '一级分类id',
    category_name varchar(16) not null comment '分类名字'
) engine InnoDB, default char set utf8mb4, comment '文章分类表';

-- 创建文章二级分类表
create table if not exists categories_second(
    cs_id int unsigned primary key auto_increment comment '分类id',
    first_id int unsigned not null comment '一级分类Id',
    category_name varchar(16) not null comment '分类名字'
) engine InnoDB, default char set utf8mb4, comment '文章分类表';

-- 创建标签表
create table if not exists labels(
    label_id int unsigned primary key auto_increment comment '标签id',
    label_name varchar(16) not null comment '标签名字'
) engine InnoDB, default char set utf8mb4, comment '标签表';

-- 创建文章对应标签表
create table if not exists article_labels(
    al_id int unsigned primary key auto_increment comment '标签id 主键',
    article_id int unsigned not null comment '文章id',
    label_id int unsigned not null comment '标签id'
) engine InnoDB, default char set utf8mb4, comment '文章对应标签表';

-- 创建文章表
create table if not exists articles(
    article_id int unsigned primary key auto_increment comment '文章id',
    owner_id int unsigned not null comment '文章拥有者id',
    category_id int unsigned not null comment '文章分类id',
    tittle varchar(64) not null comment '标题',
    abstract text comment '摘要',
    mdPath varchar(256) comment 'markdown文件路径',
    publish_time varchar(32) not null comment '发布时间',
    last_update_time varchar(32) not null comment '上次更新时间',
    visited int unsigned not null default 0  comment '访客数量',
    private boolean default false comment '是否是私有文章',
    deleted boolean default false comment '是否删除'
) engine InnoDB, default char set utf8mb4, comment '文章表';


