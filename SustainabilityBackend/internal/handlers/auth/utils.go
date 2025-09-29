package auth

import (
	"math/rand"
	"strconv"

	"github.com/ascribner/sustainabilityapp/env"
	"github.com/golang-jwt/jwt/v5"
)

// Takes in the user email and returns a token for the user
func CreateJWT(user_email string) (string, error) {
	claims := jwt.RegisteredClaims{
		Subject: user_email,
		ID:      strconv.Itoa(rand.Int()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString(env.EnvConfig.JWTSecret)

	if err != nil {
		return "", err
	}

	return tokenStr, nil
}
