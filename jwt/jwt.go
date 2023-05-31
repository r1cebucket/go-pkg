package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const TokenExpireDuration = time.Hour * 2

var Secret = []byte("人生路漫漫")

type Payload struct {
	Data interface{}
	jwt.StandardClaims
}

// get token
func GenToken(data interface{}) (string, error) {
	timeNow := time.Now()
	cla := Payload{
		data,
		jwt.StandardClaims{
			ExpiresAt: timeNow.Add(TokenExpireDuration).Unix(), // 过期时间
			NotBefore: timeNow.Unix(),
			IssuedAt:  timeNow.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cla)
	fmt.Println("Token = ", token)
	return token.SignedString(Secret) // 进行签名生成对应的token
}

// parse token
func ParseToken(tokenString string) (*Payload, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Payload{}, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Payload); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
