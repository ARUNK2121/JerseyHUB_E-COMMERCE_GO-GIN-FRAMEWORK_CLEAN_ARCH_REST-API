package helper

import (
	"time"

	"jerseyhub/pkg/utils/models"

	"github.com/golang-jwt/jwt"
)

func GenerateTokenAdmin(admin models.AdminDetailsResponse) (string, string, error) {
	accessTokenClaims := &AuthCustomClaims{
		Id:    admin.ID,
		Email: admin.Email,
		Role:  "admin",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 20).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	refreshTokenClaims := &AuthCustomClaims{
		Id:    admin.ID,
		Email: admin.Email,
		Role:  "admin",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := accessToken.SignedString([]byte("accesssecret"))
	if err != nil {
		return "", "", err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte("refreshsecret"))
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}
