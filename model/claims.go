package model

import "github.com/dgrijalva/jwt-go"

type JWTCustomClaims struct {
	UserId string
	Role   string
	jwt.StandardClaims
}
