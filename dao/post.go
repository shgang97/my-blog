package dao

import (
	"database/sql"
	"log"
	"my-blog/models"
)

/*
@author: shg
@since: 2022/10/15
@desc: //TODO
*/

func GetPostsPage(slug string, page int, pageSize int) ([]models.Post, error) {
	page = (page - 1) * pageSize
	var err error
	var rows *sql.Rows
	if slug == "" {
		rows, err = Db.Query("select * from post limit ?, ?", page, pageSize)
	} else {
		rows, err = Db.Query("select * from post where slug=? limit ?, ?", slug, page, pageSize)
	}
	if err != nil {
		log.Println("failed to query from table post:", err)
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.Pid, &post.Title, &post.Content, &post.Markdown, &post.CategoryId, &post.UserId, &post.ViewCount, &post.Type, &post.Slug, &post.CreateAt, &post.UpdateAt)
		if err != nil {
			log.Println("failed to convert to Post:", err)
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func GetPostTotal(slug string) int {
	var row *sql.Row
	if slug == "" {
		row = Db.QueryRow("select count(0) from post")
	} else {
		row = Db.QueryRow("select count(0) from post where slug=?", slug)
	}
	var count int
	err := row.Scan(&count)
	if err != nil {
		log.Println("failed to query from table post:", err)
		return 0
	}
	return count
}

func GetPostsPageByCategoryId(categoryId, page, pageSize int) ([]models.Post, error) {
	page = (page - 1) * pageSize
	rows, err := Db.Query("select * from post where category_id = ? limit ?,?", categoryId, page, pageSize)
	if err != nil {
		log.Println("failed to query from table post by categoryId:", err)
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.Pid, &post.Title, &post.Content, &post.Markdown, &post.CategoryId, &post.UserId, &post.ViewCount, &post.Type, &post.Slug, &post.CreateAt, &post.UpdateAt)
		if err != nil {
			log.Println("failed to convert to Post:", err)
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func GetPostTotalByCategoryId(categoryId int) int {
	row := Db.QueryRow("select count(0) from post where category_id = ?", categoryId)
	var count int
	err := row.Scan(&count)
	if err != nil {
		log.Println("failed to query from table post:", err)
		return 0
	}
	return count
}

func GetPostByPid(pid int) (*models.Post, error) {
	row := Db.QueryRow("select * from post where pid = ?", pid)
	var post models.Post
	if row.Err() != nil {
		return &post, row.Err()
	}
	err := row.Scan(&post.Pid, &post.Title, &post.Content, &post.Markdown, &post.CategoryId, &post.UserId, &post.ViewCount, &post.Type, &post.Slug, &post.CreateAt, &post.UpdateAt)
	if err != nil {
		return &post, err
	}
	return &post, nil
}

func GetAllPosts() ([]models.Post, error) {
	rows, err := Db.Query("select * from post")
	if err != nil {
		log.Println("failed to query from table post by categoryId:", err)
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.Pid, &post.Title, &post.Content, &post.Markdown, &post.CategoryId, &post.UserId, &post.ViewCount, &post.Type, &post.Slug, &post.CreateAt, &post.UpdateAt)
		if err != nil {
			log.Println("failed to convert to Post:", err)
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func SavePost(post *models.Post) int {
	res, err := Db.Exec("insert into post (title, content, markdown, category_id, user_id, view_count, type, slug, create_at, update_at) "+
		"values (?,?,?,?,?,?,?,?,?,?)", post.Title, post.Content, post.Markdown, post.CategoryId, post.UserId, post.ViewCount, post.Type, post.Slug, post.CreateAt, post.UpdateAt)
	if err != nil {
		log.Println("insert post failed")
		log.Println(err)
	}
	pid, _ := res.LastInsertId()
	return int(pid)
}

func UpdatePost(post *models.Post) {
	_, err := Db.Exec("update post set title=?, content=?, markdown=?,category_id=?,user_id=?,view_count=?,type=?,slug=?,create_at=?,update_at=? where pid=?", post.Title, post.Content, post.Markdown, post.CategoryId, post.UserId, post.ViewCount, post.Type, post.Slug, post.CreateAt, post.UpdateAt, post.Pid)
	if err != nil {
		log.Println("insert post failed")
		log.Println(err)
		return
	}
}
