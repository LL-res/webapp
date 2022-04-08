package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const TokenExpireduaration = time.Hour * 2

var mySecret = []byte("mySecrethh")

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenToken(username string) (string, error) {
	c := MyClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireduaration).Unix(),
			Issuer:    "bluebell",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	fmt.Println(token)
	return token.SignedString(mySecret)
}

func ParseToken(tokenString string) (*MyClaims, error) {
	var mc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (interface{}, error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid {
		return mc, nil
	}
	return nil, errors.New("invalid token")
}
