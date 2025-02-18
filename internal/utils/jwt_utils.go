package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var (
	accTokenSerct     = []byte("acctk@yinglian-ai")
	accessTokenExpire = time.Hour * 24 // 1一天
)

// Claims 定义结构
type Claims struct {
	UserId int64 `json:"user_id"`
	jwt.RegisteredClaims
}

func GenAccessToken(userId int64) (string, error) {
	claims := Claims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessTokenExpire)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(accTokenSerct)
}
func ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return accTokenSerct, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, err
	}
	return claims, nil
}
