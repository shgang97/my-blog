package main

import (
	"backend/dao"
	"backend/model"
	"fmt"
)

/*
@author: shg
@since: 2022/11/29
@desc: //TODO
*/

func main() {
	//router.Router()
	var user model.User
	dao.Db.First(&user)
	fmt.Println("user = ", user)

}
