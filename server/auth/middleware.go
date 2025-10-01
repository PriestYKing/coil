package auth

import (
	"net/http"
	"strings"
)

var publicRoutes = map[string]bool{
	"GET /":          true,
	"POST /login":    true,
	"POST /register": true,
	"POST /isAuth":   true,
}

func GlobalMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// --- CORS Logic ---
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		// --- End CORS Logic ---

		routeKey := r.Method + " " + r.URL.Path
		if publicRoutes[routeKey] {
			next.ServeHTTP(w, r)
			return
		}

		// Validate Protected Routes
		authHeader := r.Header.Get("Authorization")
		parts := strings.Split(authHeader, "Bearer ")
		if len(parts) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		tokenString := parts[1]
		isValid, err := ValidateToken(tokenString)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		if !isValid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
