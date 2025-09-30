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
		routeKey := r.Method + " " + r.URL.Path
		if publicRoutes[routeKey] {
			// Allow public routes
			next.ServeHTTP(w, r)
			return
		}

		// Validate Protected Routes
		tokenString := strings.Split(r.Header.Get("Authorization"), "Bearer ")[1]
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
