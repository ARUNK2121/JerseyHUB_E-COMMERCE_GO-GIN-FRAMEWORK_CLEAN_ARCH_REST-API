package helper

// import (
// 	"time"

// 	"jerseyhub/pkg/utils/models"

// 	"crypto/rand"
// 	"encoding/base32"

// 	"github.com/golang-jwt/jwt"
// )

// type AuthCustomClaims struct {
// 	Id    int    `json:"id"`
// 	Email string `json:"email"`
// 	Role  string `json:"role"`
// 	jwt.StandardClaims
// }

// func GenerateTokenClients(user models.UserDetailsResponse) (string, error) {
// 	claims := &AuthCustomClaims{
// 		Id:    user.Id,
// 		Email: user.Email,
// 		Role:  "client",
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

// func GenerateRefferalCode() (string, error) {
// 	// Calculate the required number of random bytes
// 	byteLength := (5 * 5) / 8

// 	// Generate a random byte array
// 	randomBytes := make([]byte, byteLength)
// 	_, err := rand.Read(randomBytes)
// 	if err != nil {
// 		return "", err
// 	}

// 	// Encode the random bytes to base32
// 	encoder := base32.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZ234567").WithPadding(base32.NoPadding)
// 	encoded := encoder.EncodeToString(randomBytes)

// 	// Trim any additional characters to match the desired length
// 	encoded = encoded[:5]

// 	return encoded, nil
// }
