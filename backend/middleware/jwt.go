package middleware

import (
	"backend/common"
	"backend/constant"
	"errors"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

/*
@author: shg
@since: 2022/12/4
@desc: //TODO
*/

const (
	// ExpireDuration 过期时间 默认7天
	ExpireDuration = 7 * 24 * time.Hour
)

var SECRET = []byte("shgang's blog")

type Claims struct {
	Id string `json:"id"`
	jwt.StandardClaims
}

func GenToken(id string) (string, error) {
	claims := Claims{
		Id: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ExpireDuration).Unix(), // 过期时间
			Issuer:    "my-blog",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SECRET)
}

func ParseToken(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return SECRET, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

func JWT() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		// 假设Token放在Header的Authorization中，并使用Bearer开头
		tokenStr := ctx.Request.Header.Get("Authorization")
		if tokenStr == "" {
			common.FailWithStatus(ctx, http.StatusUnauthorized, constant.ERROR_AUTH_NO_TOKEN)
			ctx.Abort()
			return
		}

		// 解析token
		claim, err := ParseToken(tokenStr)
		if err != nil {
			common.FailWithStatus(ctx, http.StatusUnauthorized, constant.ERROR_AUTH_INVALID_TOKEN)
			ctx.Abort()
			return
		}
		ctx.Set("userId", claim.Id)
		ctx.Next()
	}

}
