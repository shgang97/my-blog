package models

import (
	"html/template"
	"my-blog/backend/config"
	"time"
)

/*
@author: shg
@since: 2022/10/13
@desc: //TODO
*/

type Post struct {
	Id          int       `json:"id"`
	UserId      int       `json:"userId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	CreateTime  time.Time `json:"createTime"`
	UpdateTime  time.Time `json:"updateTime"`
	Status      string    `json:"status"`
}

type PagePost struct {
	Posts    []Post `json:"posts"`
	Page     int    `json:"page"`
	Total    int    `json:"total"`
	PageSize int    `json:"pageSize"`
}

func (v Post) TableName() string {
	return "AppBlogPosts"
}

type PostMore struct {
	Pid          int
	Title        string
	Slug         string
	Content      template.HTML
	Markdown     string
	CategoryId   int
	CategoryName string
	UserId       int
	UserName     string
	ViewCount    int
	Type         int
	CreateAt     string
	UpdateAt     string
}

type PostDetail struct {
	Post
	Markdown string `json:"markdown" gorm:"column:Content"`
}

type PostRe struct {
	Pid        int    `json:"pid"`
	Title      string `json:"title"`
	Slug       string `json:"slug"`
	Content    string `json:"content"`
	Markdown   string `json:"markdown"`
	CategoryId int    `json:"categoryId"`
	UserId     int    `json:"userId"`
	Type       int    `json:"type"`
}

type SearchRes struct {
	Pid   int    `json:"pid"`   // 文章ID
	Title string `json:"title"` // 文章标题
}

type PostRes struct {
	config.Viewer
	config.System
	config.Comment
	Article PostMore
}

type WritingRes struct {
	Title      string
	CdnURL     string
	Categories []Category
}
