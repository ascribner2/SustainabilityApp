package middleware

import (
	"context"
	"log"
	"net/http"

	"github.com/ascribner/sustainabilityapp/env"
	"github.com/ascribner/sustainabilityapp/internal/contextkeys"
	"github.com/golang-jwt/jwt/v5"
)

// Adds headers for common security vulnerabilities
func SecurityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		// CORS
		rw.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5173")
		rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		rw.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		rw.Header().Set("Access-Control-Allow-Credentials", "true")

		//preflight
		if r.Method == http.MethodOptions {
			rw.WriteHeader(http.StatusNoContent)
			return
		}

		// Continue the chain
		next.ServeHTTP(rw, r)
	})
}

// Verifies that a JWT is valid
func AuthenticateRoute(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		// gets the auth token cookie
		token, err := r.Cookie("auth-token")
		if err != nil {
			log.Println(err)
			rw.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Parse and validate the JWT
		parser := jwt.NewParser()
		validToken, err := parser.Parse(token.Value, func(*jwt.Token) (any, error) {
			return []byte(env.EnvConfig.JWTSecret), nil
		})
		if err != nil {
			log.Println(err)
			rw.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Get the email from the claims
		claims := validToken.Claims
		email, err := claims.GetSubject()
		if err != nil {
			log.Println(err)
			rw.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Adds the email to the context
		var emailKey contextkeys.AuthContextKey = "email"
		ctx := context.WithValue(r.Context(), emailKey, email)
		r = r.Clone(ctx)

		next(rw, r)
	}
}
