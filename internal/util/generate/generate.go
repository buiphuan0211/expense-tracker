package pgenerate

import (
	"expense-tracker/internal/config"
	"expense-tracker/internal/util/ptime"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// HashPassword ...
func HashPassword(password string) string {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return string(hashPassword)
}

// CheckPasswordHash ...
func CheckPasswordHash(hashPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}

// TokenString ...
func TokenString(data interface{}) (string, error) {
	var (
		envVars = config.GetENV()
		claims  = &jwt.MapClaims{
			"exp":  ptime.Now().Add(time.Hour * 24).Unix(),
			"data": data,
		}
		jwtToken = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	)

	tokenString, err := jwtToken.SignedString([]byte(envVars.Auth.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
