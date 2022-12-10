# 创建数据库
CREATE DATABASE blog CHARACTER SET utf8;
USE blog;

# 用户表
CREATE TABLE `blog_user`
(
    `pk`        int PRIMARY KEY AUTO_INCREMENT COMMENT '主键',
    `id`        varchar(64) NOT NULL COMMENT 'id',
    `username`  varchar(64) NOT NULL COMMENT '用户名',
    `password`  varchar(64) NOT NULL COMMENT '密码',
    `email`     varchar(64) NOT NULL COMMENT '用户email',
    `avatar`    varchar(64) NOT NULL COMMENT '用户头像url',
    `ip`        varchar(64) NOT NULL COMMENT '用户ip',
    `create_at` datetime    NOT NULL COMMENT '创建时间',
    `update_at` datetime    NOT NULL COMMENT '创建时间',
    CONSTRAINT id UNIQUE (`id`),
    CONSTRAINT username UNIQUE (`username`),
    CONSTRAINT email UNIQUE (`email`)
) COMMENT '用户表';

# 文章表
CREATE TABLE `blog_article`
(
    `pk`            int PRIMARY KEY AUTO_INCREMENT COMMENT '主键',
    `id`            varchar(64) NOT NULL COMMENT 'id',
    `user_id`       varchar(64) NOT NULL COMMENT '用户id',
    `title`         varchar(64) NOT NULL COMMENT '文章标题',
    `content`       text        NOT NULL COMMENT '文章内容',
    `cover`         varchar(64) NOT NULL COMMENT '文章封面图url',
    `view_count`    int         NOT NULL DEFAULT 0 COMMENT '浏览量',
    `like_count`    int         NOT NULL DEFAULT 0 COMMENT '点赞数',
    `comment_count` int         NOT NULL DEFAULT 0 COMMENT '评论数',
    `create_at`     datetime    NOT NULL COMMENT '创建时间',
    `update_at`     datetime    NOT NULL COMMENT '创建时间',
    CONSTRAINT id UNIQUE (`id`)
) COMMENT '文章表';

# 分类表
CREATE TABLE `blog_category`
(
    `pk`          int PRIMARY KEY AUTO_INCREMENT COMMENT '主键',
    `id`          varchar(64) NOT NULL COMMENT 'id',
    `name`        varchar(64) NOT NULL COMMENT '分类名称',
    `description` text        NOT NULL COMMENT '分类描述',
    `parent_id`   varchar(64) NOT NULL COMMENT '父分类id',
    `create_at`   datetime    NOT NULL COMMENT '创建时间',
    `update_at`   datetime    NOT NULL COMMENT '创建时间',
    CONSTRAINT id UNIQUE (`id`),
    CONSTRAINT name UNIQUE (`name`)
) COMMENT '分类表';

# 标签表
CREATE TABLE `blog_tag`
(
    `pk`          int PRIMARY KEY AUTO_INCREMENT COMMENT '主键',
    `id`          varchar(64) NOT NULL COMMENT 'id',
    `name`        varchar(64) NOT NULL COMMENT '标签名称',
    `description` text        NOT NULL COMMENT '标签描述',
    `create_at`   datetime    NOT NULL COMMENT '创建时间',
    `update_at`   datetime    NOT NULL COMMENT '创建时间',
    CONSTRAINT id UNIQUE (`id`),
    CONSTRAINT name UNIQUE (`name`)
) COMMENT '标签表';

# 评论表
CREATE TABLE `blog_comment`
(
    `pk`            int PRIMARY KEY AUTO_INCREMENT COMMENT '主键',
    `id`            varchar(64) NOT NULL COMMENT 'id',
    `user_id`       varchar(64) NOT NULL COMMENT '发表用户id',
    `target_id`     varchar(64) NOT NULL COMMENT '评论目标id',
    `type`          int         NOT NULL COMMENT '评论类型：1-文章评论，2-评论评论，3-回复评论',
    `like_count`    int         NOT NULL DEFAULT 0 COMMENT '点赞数',
    `comment_count` int         NOT NULL DEFAULT 0 COMMENT '评论数',
    `create_at`     datetime    NOT NULL COMMENT '创建时间',
    `update_at`     datetime    NOT NULL COMMENT '创建时间',
    CONSTRAINT id UNIQUE (`id`)
) COMMENT '评论表';

# 文章分类表
CREATE TABLE `blog_article_category`
(
    `pk`          int PRIMARY KEY AUTO_INCREMENT COMMENT '主键',
    `id`          varchar(64) NOT NULL COMMENT 'id',
    `article_id`  varchar(64) NOT NULL COMMENT '文章id',
    `category_id` varchar(64) NOT NULL COMMENT '分类id',
    `create_at`   datetime    NOT NULL COMMENT '创建时间',
    `update_at`   datetime    NOT NULL COMMENT '创建时间',
    CONSTRAINT id UNIQUE (`id`)
) COMMENT '文章分类表';

# 文章标签表
CREATE TABLE `blog_article_tag`
(
    `pk`         int PRIMARY KEY AUTO_INCREMENT COMMENT '主键',
    `id`         varchar(64) NOT NULL COMMENT 'id',
    `article_id` varchar(64) NOT NULL COMMENT '文章id',
    `tag_id`     varchar(64) NOT NULL COMMENT '标签id',
    `create_at`  datetime    NOT NULL COMMENT '创建时间',
    `update_at`  datetime    NOT NULL COMMENT '创建时间',
    CONSTRAINT id UNIQUE (`id`)
) COMMENT '文章标签表';



