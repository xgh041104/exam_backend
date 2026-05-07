package jwt_use

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const TokenExpireDuration = time.Hour * 2

// const TokenExpireDuration = time.Second * 60

var Secret = []byte("studyexam")

type MyClaims struct {
	UserName string
	UserType int //0 是管理员和老师  1是学生 2是站点用户
	jwt.StandardClaims
}

// get token
func GetToken(username string, userType int) (string, error) {
	cla := MyClaims{
		username,
		userType,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "lx-jwt",                                   // 签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cla)
	fmt.Println("Token = ", token)
	return token.SignedString(Secret) // 进行签名生成对应的token
}

// parse token
func ParseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
