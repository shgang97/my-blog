package dao

import (
	"backend/model"
)

/*
@author: shg
@since: 2022/12/4
@desc: //TODO
*/

func GetUser(username, password string) (*model.User, error) {
	var user model.User
	result := Db.Select("id", "username", "avatar").Where("username = ?", username).Where("password = ?", password).Take(&user)
	if result.RowsAffected > 0 {
		return &user, nil
	}
	return nil, result.Error
}
