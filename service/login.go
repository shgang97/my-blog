package service

import (
	"errors"
	"log"
	"my-blog/dao"
	"my-blog/models"
	"my-blog/util"
)

/*
@author: shg
@since: 2022/10/16
@desc: //TODO
*/

func Login(userName, password string) (*models.LoginRes, error) {
	password = util.Md5Crypt(password, "shgang")
	user := dao.GetUser(userName, password)
	if user == nil {
		log.Println("login failed")
		return nil, errors.New("login failed")
	}
	uid := user.Id
	// 生成token jwt
	token, err := util.Award(&uid)
	if err != nil {
		log.Println("util.Award:", err)
		return nil, errors.New("failed to generate token")
	}
	var userInfo models.UserInfo
	userInfo.Id = user.Id
	userInfo.UserName = user.UserName
	userInfo.Avatar = user.Avatar
	var loginRes = &models.LoginRes{
		Token:    token,
		UserInfo: userInfo,
	}
	return loginRes, nil
}
