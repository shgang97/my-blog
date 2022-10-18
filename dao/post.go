package dao

import (
	"log"
	"my-blog/models"
)

/*
@author: shg
@since: 2022/10/15
@desc: //TODO
*/

func GetPostsPage(page int, pageSize int) ([]models.Post, error) {
	page = (page - 1) * pageSize
	rows, err := Db.Query("select * from post limit ?,?", page, pageSize)
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

func GetPostTotal() int {
	row := Db.QueryRow("select count(0) from post")
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
