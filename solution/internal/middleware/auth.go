package middleware

import (
	"context"
	"net/http"
	"strings"
	"solution/internal/config"
	"solution/pkg/jwt"
)

type ContextKey string

const (
	ContextUserIDKey  ContextKey = "user_id"
	ContextRoleKey    ContextKey = "user_role"
	ContextJWTDataKey ContextKey = "jwt_data"
)

func IsAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		
		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Missing token"))
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Invalid token format"))
			return
		}

		token := parts[1]
		isValid, data, err := jwt.NewJWT(config.AppConfig().RandomSecret.RANDOM_SECRET()).Parse(token)
		if err != nil || !isValid || data == nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Invalid token"))
			return
		}

		if data.UserID == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Token missing user ID"))
			return
		}

		if data.Role != "USER" && data.Role != "ADMIN" {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Invalid user role"))
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, ContextUserIDKey, data.UserID)
		ctx = context.WithValue(ctx, ContextRoleKey, data.Role)
		ctx = context.WithValue(ctx, ContextJWTDataKey, data)

		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}

func IsAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		checkRole := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			role, ok := r.Context().Value(ContextRoleKey).(string)
			if !ok {
				http.Error(w, "Unable to determine user role", http.StatusInternalServerError)
				return
			}

			if role != "ADMIN" {
				http.Error(w, "Admin access required", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})

		IsAuth(checkRole).ServeHTTP(w, r)
	})
}
