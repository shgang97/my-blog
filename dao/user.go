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

func GetUserNameById(id int) string {
	row := Db.QueryRow("select user_name from user where id = ?", id)
	if row.Err() != nil {
		log.Println(row.Err())
	}
	var categoryName string
	_ = row.Scan(&categoryName)
	return categoryName
}

func GetUser(userName, password string) *models.User {
	row := Db.QueryRow("select * from user where user_name = ? and password = ?", userName, password)
	if row.Err() != nil {
		log.Println(row.Err())
		return nil
	}
	var user models.User
	err := row.Scan(&user.Id, &user.UserName, &user.Password, &user.Avatar, &user.CreateAt, &user.UpdateAt)
	if err != nil {
		log.Println("dao GetUser:", err)
		return nil
	}
	return &user
}
