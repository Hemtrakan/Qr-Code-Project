package utility

import (
	"github.com/dgrijalva/jwt-go"
	"qrcode/access/constant"
	"time"
)

func AuthenticationLogin(id uint ,role string) (Token string ,Error error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	Token, err := token.SignedString([]byte(constant.SecretKey))
	return Token ,err
}


