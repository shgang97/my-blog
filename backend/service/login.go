package service

import (
	"backend/dao"
	"backend/middleware"
	"backend/model"
	"backend/request"
	"backend/response"
	"backend/utils"
)

/*
@author: shg
@since: 2022/12/4
@desc: //TODO
*/

func Login(request *request.LoginRequest) (*response.LoginResponse, error) {
	username := request.Username
	password := request.Password
	password = utils.Md5Crypt(password, "salt")
	// 校验用户名和密码查询
	user, err := dao.GetUser(username, password)
	if err != nil {
		return nil, err
	}
	// 生成 token
	token, err := middleware.GenToken(user.Id)
	if err != nil {
		return nil, err
	}
	userInfo := &model.UserInfo{
		Id:       user.Id,
		Username: user.Username,
		Avatar:   user.Avatar,
	}
	loginUserRes := &response.LoginResponse{
		UserInfo: userInfo,
		Token:    token,
	}
	return loginUserRes, nil
}
