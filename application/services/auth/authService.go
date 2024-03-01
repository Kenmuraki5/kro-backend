package auth

import (
	"errors"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var secretKey []byte

func init() {
	err := godotenv.Load("env/.env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	secretKey = []byte(os.Getenv("SECRET_KEY"))
}

type AuthService struct{}

func (s *AuthService) GenerateToken(userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userId,
		"iss": "your-issuer",
	})

	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (s *AuthService) ValidateToken(tokenString string) (string, *jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return secretKey, nil
	})

	if err != nil {
		return "", nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId, ok := claims["sub"].(string)
		if !ok {
			return "", nil, errors.New("user ID not found in token claims")
		}

		return userId, token, nil
	}

	return "", nil, errors.New("invalid token claims")
}
