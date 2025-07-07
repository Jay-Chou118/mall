package util

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte("yijiansanlian")

type Claims struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Authority int    `json:"authority"`
	jwt.StandardClaims
}

func GenerateToken(id uint, userName string, authority int) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Hour * 24)
	claims := Claims{
		ID:        id,
		Username:  userName,
		Authority: authority,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "FanOne-Mall",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err

}

// 验证用户token
func ParseToken(token string) (*Claims, error) {
	tokenCliams, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenCliams != nil {
		if claims, ok := tokenCliams.Claims.(*Claims); ok && tokenCliams.Valid {
			return claims, nil
		}
	}
	return nil, err
}
