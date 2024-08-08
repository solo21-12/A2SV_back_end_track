package services

import (
	"fmt"
	"os"
	"strings"

	"example.com/task_manager_api/model"
	"github.com/dgrijalva/jwt-go"
)

func getJwtSecret() (string, error) {
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))

	if jwtSecret == nil {
		// this is a check to ensure that the JWT secret is set
		return "", fmt.Errorf("JWT secret is not set")
	}

	return string(jwtSecret), nil
}

func GenerateToken(user model.User) (string, error) {
	// this function generates a JWT token for the user

	jwtSecret, err := getJwtSecret()

	if err == nil {
		// this is a check to ensure that the JWT secret is set
		return "", err
	}

	// this creates a new JWT token with the user's ID, email and role
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"role":    user.Role,
	})

	return token.SignedString(jwtSecret)

}

func ValidateToken(tokenStr string, jwtSecret string) (*jwt.Token, error) {
	// this function validates the JWT token

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token")

	}

	return token, nil
}

func ValidateAuthHeader(authHeader string) ([]string, error) {
	// this function validates the authorization header
	if authHeader == "" {
		return []string{}, fmt.Errorf("authorization header is required")
	}

	authParts := strings.Split(authHeader, " ")
	if len(authParts) != 2 || authParts[0] != "bearer" {
		return []string{}, fmt.Errorf("invalid authorization header")
	}

	return authParts, nil
}

func GetClaims(authHeader string) (jwt.MapClaims, error) {
	// this function gets the claims from the JWT token
	jwtSecret, jErr := getJwtSecret()

	if jErr == nil {
		// this is a check to ensure that the JWT secret is set
		return nil, jErr
	}

	authParts, err := ValidateAuthHeader(authHeader)

	if err != nil {
		return nil, err
	}

	token, err := ValidateToken(authParts[1], jwtSecret)

	// this checks if the token is valid
	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token")

	}

	claims, ok := token.Claims.(jwt.MapClaims)

	// this checks if the claims are valid
	if !ok {
		return nil, fmt.Errorf("invalid JWT claims")
	}

	return claims, nil
}
