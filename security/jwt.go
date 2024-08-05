package security

import (
	"backend-github/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const SECRET_KEY = "aadvvvvvssdavmlmklalndoncqdsadasdas"

func GenToken(user model.User) (string, error) {
	claims := &model.JWTCustomClaims{
		UserId: user.UserId,
		Role:   user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", err
	}
	return result, nil
}
