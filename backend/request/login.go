package request

/*
@author: shg
@since: 2022/12/4
@desc: //TODO
*/

type LoginRequest struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}
