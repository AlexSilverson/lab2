package entity

import (
	"github.com/dgrijalva/jwt-go"
)

var SecretKey = "ItIsMyPassword"

type Claims struct {
	Role uint `json:"role"`
	jwt.StandardClaims
}
