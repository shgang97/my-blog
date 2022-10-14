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
	Id                   string    `json:"id" gorm:"column:Id"`
	Title                string    `json:"title" gorm:"column:Title"`
	Slug                 string    `json:"slug" gorm:"column:Slug"`
	Description          string    `json:"description" gorm:"column:Description"`
	Cover                string    `json:"cover" gorm:"column:Cover"`
	CopyrightType        int       `json:"copyrightType" gorm:"column:CopyrightType"`
	Original             string    `json:"original" gorm:"column:Original"`
	OriginalAvatar       string    `json:"originalAvatar" gorm:"column:OriginalAvatar"`
	OriginalTitle        string    `json:"originalTitle" gorm:"column:OriginalTitle"`
	OriginalLink         string    `json:"originalLink" gorm:"column:OriginalLink"`
	ViewCount            int       `json:"viewCount" gorm:"column:ViewCount"`
	CreationTime         time.Time `json:"creationTime" gorm:"column:CreationTime"`
	LastModificationTime time.Time `json:"lastModificationTime" gorm:"column:LastModificationTime"`
}

func (v Post) TableName() string {
	return "AppBlogPosts"
}

type PostMore struct {
	Id                   string        `json:"id"`
	Title                string        `json:"title"`
	Slug                 string        `json:"slug"`
	Description          string        `json:"description"`
	Cover                string        `json:"cover"`
	Markdown             string        `json:"markdown"`
	Content              template.HTML `json:"content"`
	CopyrightType        int           `json:"copyrightType"`
	Original             string        `json:"original"`
	OriginalAvatar       string        `json:"originalAvatar"`
	OriginalTitle        string        `json:"originalTitle"`
	OriginalLink         string        `json:"originalLink"`
	Categories           []Category    `json:"categories"`
	ViewCount            int           `json:"viewCount"`
	Type                 int           `json:"type"`
	CreationTime         string        `json:"creationTime"`
	LastModificationTime string        `json:"lastModificationTime"`
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
