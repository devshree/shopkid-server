package middleware

import (
	"context"
	"database/sql"
	"net/http"
	"os"
	"strings"

	"kids-shop/internal/domain/models"
	"kids-shop/internal/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

func AuthMiddleware(db *sql.DB) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Skip auth for OPTIONS requests and login/register endpoints
			if r.Method == "OPTIONS" || 
			   r.URL.Path == "/api/auth/login" || 
			   r.URL.Path == "/api/auth/register" {
				next.ServeHTTP(w, r)
				return
			}

			// Get token from Authorization header
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Authorization header required", http.StatusUnauthorized)
				return
			}

			tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

			// Parse and validate token
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("JWT_SECRET")), nil
			})

			if err != nil || !token.Valid {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			// Extract user ID from claims
			claims := token.Claims.(jwt.MapClaims)
			userID := int(claims["user_id"].(float64))

			// Add user ID to request context
			ctx := context.WithValue(r.Context(), models.UserIDKey, userID)
			
			// Get user role using the service function
			userRole, err := service.GetUserRoleFromID(userID, db)
			if err != nil {
				http.Error(w, "Error fetching user role", http.StatusInternalServerError)
				return
			}
			ctx = context.WithValue(ctx, models.UserRoleKey, userRole)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
} 