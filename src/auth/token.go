package auth

import (
	"api/src/config"
	"errors"
	"fmt"
	jwt "github.com/golang-jwt/jwt/v5"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func Generate(userId uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userId"] = userId
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString(config.JwtSecret)
}

func Validate(r *http.Request) error {
	tokenString := getToken(r)
	token, err := jwt.Parse(tokenString, getKey)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("Token Inválido")
}

func GetUserId(r *http.Request) (uint64, error) {
	tokenString := getToken(r)
	token, err := jwt.Parse(tokenString, getKey)
	if err != nil {
		return 0, err
	}

	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, err := strconv.ParseUint(fmt.Sprintf("%.0f", permissions["userId"]), 10, 64)
		if err != nil {
			return 0, nil
		}
		return userID, nil
	}
	return 0, errors.New("Token Inválido")
}

func getToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	splitToken := strings.Split(token, " ")

	if len(splitToken) == 2 {
		return splitToken[1]
	}

	return token
}

func getKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("signin method unexpected! %v", token.Header["alg"])
	}

	return config.JwtSecret, nil
}
