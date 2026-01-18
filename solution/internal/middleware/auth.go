package middleware

import (
	"context"
	"net/http"
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

		isValid, data, err := jwt.NewJWT(config.AppConfig().RandomSecret.RANDOM_SECRET()).Parse(authHeader)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Invalid token"))
			return
		}

		if !isValid || data == nil {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		if data.UserID == "" {
			http.Error(w, "Token missing user ID", http.StatusUnauthorized)
			return
		}

		if data.Role != "USER" && data.Role != "ADMIN" {
			http.Error(w, "Invalid user role", http.StatusForbidden)
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
