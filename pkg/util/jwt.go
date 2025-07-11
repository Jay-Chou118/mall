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

type EmailClaims struct {
	UserID        uint   `json:"user_id"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	OperationType uint   `json:"operation_type"`
	jwt.StandardClaims
}

// 签发email token
func GenerateEmailToken(useId, Operation uint, email, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Hour * 24)
	claims := EmailClaims{
		UserID:        useId,
		Email:         email,
		Password:      password,
		OperationType: Operation,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "FanOne-Mall",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// 解析email token
func ParseEmailToken(token string) (*EmailClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &EmailClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*EmailClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
