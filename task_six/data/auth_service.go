package data

import (
	"context"
	"fmt"
	"os"
	"strings"

	"example.com/task_manager_api/model"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func getJwtSecret() ([]byte, error) {
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))

	if len(jwtSecret) == 0 {
		return nil, fmt.Errorf("JWT secret is not set")
	}

	return jwtSecret, nil
}

func validatePassword(user model.User, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) == nil
}

func GenerateToken(user model.User) (string, error) {
	jwtSecret, err := getJwtSecret()
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"role":    user.Role,
	})

	return token.SignedString(jwtSecret)
}

func ValidateToken(tokenStr string, jwtSecret []byte) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	return token, nil
}

func ValidateAuthHeader(authHeader string) ([]string, error) {
	if authHeader == "" {
		return nil, fmt.Errorf("authorization header is required")
	}

	authParts := strings.Split(authHeader, " ")
	if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
		return nil, fmt.Errorf("invalid authorization header")
	}

	return authParts, nil
}

func GetClaims(authHeader string) (jwt.MapClaims, error) {
	jwtSecret, err := getJwtSecret()
	if err != nil {
		return nil, err
	}

	authParts, err := ValidateAuthHeader(authHeader)
	if err != nil {
		return nil, err
	}

	token, err := ValidateToken(authParts[1], jwtSecret)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid JWT claims")
	}

	return claims, nil
}

func LoginUser(user model.UserLogin, userService UserService, ctx context.Context) (model.Auth, error) {

	curUser, err := userService.GetUser(ctx, user.Email)
	if err != nil {
		return model.Auth{}, err
	}

	if !validatePassword(curUser, user.Password) {
		return model.Auth{}, fmt.Errorf("invalid email or password")
	}

	token, err := GenerateToken(curUser)
	if err != nil {
		return model.Auth{}, fmt.Errorf("internal server error: %v", err)
	}

	return model.Auth{
		User:  user,
		Token: token,
	}, nil
}
