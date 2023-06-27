package helper

import (
	"time"

	"jerseyhub/pkg/utils/models"

	"github.com/golang-jwt/jwt"
)

type authCustomClaims struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

// func GenerateTokenUsers(user models.UserDetailsResponse) (string, error) {

// 	claims := &authCustomClaimsUsers{
// 		Id:    user.Id,
// 		Email: user.Email,
// 		Role:  "user",
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
// 			IssuedAt:  time.Now().Unix(),
// 		},
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	tokenString, err := token.SignedString([]byte("comebuyjersey"))

// 	if err != nil {
// 		return "", err
// 	}

// 	return tokenString, nil

// }

func GenerateTokenClients(user models.UserDetailsResponse) (string, error) {
	claims := &authCustomClaims{
		Id:    user.Id,
		Email: user.Email,
		Role:  "client",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("comebuyjersey"))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
