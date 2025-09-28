package auth

import "github.com/golang-jwt/jwt/v5"

// Takes in the user email and returns a token for the user
func CreateJWT(user_email string) *jwt.Token {

	claims := jwt.RegisteredClaims{
		Subject: user_email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token
}
