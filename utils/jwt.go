package utils

import (
	"be-ruang-warga/internal/user/domain"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWT(user domain.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"role":    user.Role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
