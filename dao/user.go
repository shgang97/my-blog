package dao

import "log"

/*
@author: shg
@since: 2022/10/15
@desc: //TODO
*/

func GetUserNameById(id int) string {
	row := Db.QueryRow("select name from user where id = ?", id)
	if row.Err() != nil {
		log.Println(row.Err())
	}
	var categoryName string
	_ = row.Scan(&categoryName)
	return categoryName
}
