package middleware

import (
	"backend-github/model"
	"backend-github/security"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func JWTMiddlewares() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		Claims: &model.JWTCustomClaims{},
		//SigningKey: security.SECRET_KEY,
		SigningKey: []byte(security.SECRET_KEY),
	}
	return middleware.JWTWithConfig(config)
	//return echojwt.Config{
	//	Claims:     &model.JWTCustomClaims{},
	//	SigningKey: []byte(security.SECRET_KEY),
	//}
}
