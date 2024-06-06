package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = "qwertyacid12345acidqwerty"

func GenerateJWT(userId int64) (string, time.Time, error) {
	exp := time.Now().Add(time.Hour * 2)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"exp":    exp.Unix(),
	})

	tokenStr, err := token.SignedString([]byte(secretKey))

	return tokenStr, exp, err
}

func VerifyJWT(tokenString string) (int64, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, errors.New("invalid token " + err.Error())
	}

	if !token.Valid {
		return 0, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token")
	}

	userId, _ := claims["userId"].(float64)

	return int64(userId), nil
}
