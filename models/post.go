package models

import (
	"html/template"
	"my-blog/config"
	"time"
)

/*
@author: shg
@since: 2022/10/13
@desc: //TODO
*/

type Post struct {
	Pid        int
	Title      string
	Content    string
	Markdown   string
	CategoryId int
	UserId     int
	ViewCount  int
	Type       int
	Slug       string
	CreateAt   time.Time
	UpdateAt   time.Time
}

func (v Post) TableName() string {
	return "AppBlogPosts"
}

type PostMore struct {
	Pid          int
	Title        string
	Slug         string
	Content      template.HTML
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
	Pid   int    `orm:"pid" json:"pid"`     // 文章ID
	Title string `orm:"title" json:"title"` // 文章标题
}

type PostRes struct {
	config.Viewer
	config.System
	Article PostMore
}
