package jwt

import (
	"errors"
	"fmt"
	"go-mail-sender/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type CustomClaims struct {
	UserID uuid.UUID `json:"id"`
	Email  string    `json:"email"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(email string, userID uuid.UUID, c *config.Config) (string, error) {
	accessSigningKey := []byte(c.JWTKey)
	accessToken := jwt.New(jwt.SigningMethodHS256)
	accessClaims := accessToken.Claims.(jwt.MapClaims)
	accessClaims["authorized"] = true
	accessClaims["email"] = email
	accessClaims["id"] = userID
	accessClaims["exp"] = time.Now().Add(time.Minute * 43200).Unix()
	accessTokenString, err := accessToken.SignedString(accessSigningKey)
	if err != nil {
		return "", err
	}
	return accessTokenString, nil
}

func ValidateJWTToken(tokenString string, cfg *config.Config) (*CustomClaims, error) {
	claims := &CustomClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(cfg.JWTKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid JWT token")
}
