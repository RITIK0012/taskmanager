package utils

import (
	"time"
	"github.com/golang-jwt/jwt"
)

var jwtKey =[]byte("supersecretkey")

func GenerateToken(username string) (string,error){
	claims := jwt.MapClaims{
		"username": username,
		"exp": time.Now().Add(time.Hour*2).Unix(),
	}
	token :=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	return token.SignedString(jwtKey)
}


func ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
}