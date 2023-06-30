package middleware

import (
	"fmt"
	"jerseyhub/pkg/helper"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AdminAuthMiddleware(c *gin.Context) {
	// Get the access token from the header.
	accessToken := c.Request.Header.Get("Authorization")
	// Check if the access token is valid.
	_, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte("accesssecret"), nil
	})

	if err == nil {
		fmt.Println("heyy")
		c.Next()
	}

	// // Extract the claims from the token.
	// claims, ok := token.Claims.(jwt.MapClaims)
	// if !ok {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
	// 	c.Abort()
	// 	return
	// }

	// // Check the expiry manually.
	// expiryUnix, ok := claims["exp"].(float64)
	// if !ok {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
	// 	c.Abort()
	// 	return
	// }

	// expiryTime := time.Unix(int64(expiryUnix), 0)
	// if expiryTime.After(time.Now()) {
	// 	c.Next()
	// 	return
	// }
	fmt.Println("yes my boy")
	refreshToken := c.Request.Header.Get("RefreshToken")

	// Check if the refresh token is valid.
	_, err = jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		return []byte("refreshsecret"), nil
	})
	if err != nil {
		// The refresh token is invalid.
		c.AbortWithStatus(401)
		return
	}
	// The access token is invalid. Check the refresh token.

	// Extract the claims from the token.
	// claims, ok = refresh.Claims.(jwt.MapClaims)
	// if !ok {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
	// 	c.Abort()
	// 	return
	// }

	// // Check the expiry manually.
	// expiryUnix, ok = claims["exp"].(float64)
	// if !ok {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
	// 	c.Abort()
	// 	return
	// }

	// expiryTime = time.Unix(int64(expiryUnix), 0)
	// if expiryTime.Before(time.Now()) {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Expired authorization token"})
	// 	c.Abort()
	// 	return
	// }

	// The refresh token is valid. Generate a new access token.
	newAccessToken, err := CreateNewAccessTokenAdmin()
	if err != nil {
		fmt.Println(err)
		// An error occurred while generating the new access token.
		c.AbortWithStatus(500)
		return
	}

	// Set the new access token in the response header.
	fmt.Println("accesstoken validity extended")
	c.Header("Authorization", "Bearer "+newAccessToken)
	c.Next()
}

func CreateNewAccessTokenAdmin() (string, error) {
	claims := &helper.AuthCustomClaims{
		Role: "admin",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	newAccessToken, err := token.SignedString([]byte("accesssecret"))
	if err != nil {
		return "", err
	}
	fmt.Println("created and returned")
	return newAccessToken, nil
}
