package util

/*
@author: shg
@since: 2022/10/17
@desc: //TODO
*/

import (
	gojwt "github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

var jwtKey []byte

func init() {
	jwtKey = []byte(os.Getenv("JWT_SECRET"))
}

type Claims struct {
	Uid int
	gojwt.StandardClaims
}

func Award(uid *int) (string, error) {
	// 过期时间 默认7天
	expireTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		Uid: *uid,
		StandardClaims: gojwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	// 生成token
	token := gojwt.NewWithClaims(gojwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

// ParseToken 解析token
func ParseToken(tokenStr string) (*gojwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := gojwt.ParseWithClaims(tokenStr, claims, func(token *gojwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, claims, err
}
