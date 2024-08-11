package infrastructure

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	domain "github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/Domain"
)

func CreateAccessToken(user domain.UserDTO, secret string) (accessToken string, err error) {
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
