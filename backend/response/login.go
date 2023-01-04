package response

import "backend/model"

/*
@author: shg
@since: 2022/12/4
@desc: //TODO
*/

type LoginResponse struct {
	Token    string          `json:"token"`
	UserInfo *model.UserInfo `json:"userInfo"`
}
