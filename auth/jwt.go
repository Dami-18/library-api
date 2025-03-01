package auth

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("secretkey")

type JWTClaim struct {
	Username string `json:"username"`
	ID       uint    `json:"userId" gorm:"primaryKey"`
	jwt.StandardClaims
}

func GenerateJWT(username string, userId uint) (tokenString string, err error) {
	expirationTime := time.Now().Add(2 * time.Hour)
	claims := &JWTClaim{
		Username: username,
		ID: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func ValidateTokenAndGetClaims(signedToken string) (claim JWTClaim, err error) {
	token, err := jwt.ParseWithClaims(
	  signedToken,
	  &JWTClaim{},
	  func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	  },
	)
	if err != nil {
	  return
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
	  err = errors.New("error in parsing claims")
	  return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
	  err = errors.New("token expired")
	  return
	}

	claim = *claims
	return
  }
