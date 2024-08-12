package infrastructure

import (
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	domain "github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/Domain"
	"github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/bootstrap"
)

func CreateAccessToken(user domain.UserDTO, secret []byte) (accessToken string, err error) {
	expTime := time.Now().Add(time.Minute * 30).Unix()
	claims := &domain.JWTCustome{
		ID:    user.ID.Hex(),
		Email: user.Email,
		Role:  user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(secret)

	if err != nil {
		return "", err
	}

	return t, err

}

func GetJwtSecret() ([]byte, error) {
	env := *bootstrap.NewEnv()

	if len(env.JWT_SECRET) == 0 {
		return nil, fmt.Errorf("JWT secret is not set")
	}

	return []byte(env.JWT_SECRET), nil
}

func ValidateToken(tokenStr string, jwtSecret []byte) (*domain.JWTCustome, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &domain.JWTCustome{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	claims, ok := token.Claims.(*domain.JWTCustome)
	if !ok {
		return nil, fmt.Errorf("invalid JWT claims")
	}

	return claims, nil
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

func GetClaims(authHeader string) (*domain.JWTCustome, error) {
	// Retrieve the JWT secret key
	jwtSecret, err := GetJwtSecret()
	if err != nil {
		return nil, fmt.Errorf("failed to get JWT secret: %v", err)
	}

	// Validate and parse the Authorization header
	authParts, err := ValidateAuthHeader(authHeader)
	if err != nil {
		return nil, fmt.Errorf("invalid authorization header: %v", err)
	}

	// Validate the JWT token
	claims, err := ValidateToken(authParts[1], jwtSecret)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	return claims, nil
}
