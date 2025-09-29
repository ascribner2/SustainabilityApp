package middleware

import "net/http"

// Adds headers for common security vulnerabilities
func SecurityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		// CORS
		rw.Header().Set("Access-Control-Allow-Origin", "*")
		rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		rw.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		//preflight
		if r.Method == http.MethodOptions {
			rw.WriteHeader(http.StatusNoContent)
			return
		}

		// Continue the chain
		next.ServeHTTP(rw, r)
	})
}

// TODO: implement this to authenticate a jwt token
// Verifies that a JWT is valid
func AuthenticateRoute(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

		next(rw, r)
	}
}
