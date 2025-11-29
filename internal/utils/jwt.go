package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/tibin-peter/Turf-Booking-System/internal/model"
)

var jwtkey = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	Email string
	Role  string
	jwt.RegisteredClaims
}

func GenerateAccessToken(u *model.User) (string, error) {
	exp := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Email: u.Email,
		Role:  u.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, err := token.SignedString(jwtkey)
	return s, err
}
func GenerateRefreshToken(u *model.User) (string, int64) {
	exp := time.Now().Add(24 * time.Hour).Unix()
	claims := &Claims{
		Email: u.Email,
		Role:  u.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(exp, 0)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, _ := token.SignedString(jwtkey)
	return s, exp
}
func ValidateToken(token string) (*Claims, bool) {
	Claims := &Claims{}
	t, err := jwt.ParseWithClaims(token, Claims, func(t *jwt.Token) (any, error) {
		return jwtkey, nil
	})
	return Claims, err == nil && t.Valid
}
