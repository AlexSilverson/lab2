package utility

import (
	"AlexSilverson/lab2/src/domain/entity"

	"github.com/dgrijalva/jwt-go"
)

func ParseToken(tokenString string) (claims *entity.Claims, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &entity.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("it_is_my_password"), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*entity.Claims)

	if !ok {
		return nil, err
	}

	return claims, nil
}
