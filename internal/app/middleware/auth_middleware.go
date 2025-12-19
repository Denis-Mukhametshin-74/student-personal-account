package middleware

import (
	"context"
	"net/http"
	"strings"
	"student-portal/pkg/jwt"
)

type contextKey string

const (
	studentIDKey contextKey = "studentID"
	emailKey     contextKey = "email"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Требуется авторизация", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Неверный формат токена", http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]

		claims, err := jwt.ValidateToken(tokenString)
		if err != nil {
			http.Error(w, "Неверный токен", http.StatusUnauthorized)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, studentIDKey, claims.StudentID)
		ctx = context.WithValue(ctx, emailKey, claims.Email)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
