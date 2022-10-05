package auth

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = "secretkeyhere"

type JWTClime struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

func GenerateJWT(email string, username string) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClime{
		Email:    email,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(jwtKey))

	return
}

func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(signedToken,
		&JWTClime{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		})

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*JWTClime)

	if !ok {
		err = errors.New("could not parse claims")
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}

	return
}
