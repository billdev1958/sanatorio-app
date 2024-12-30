package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sanatorioApp/internal/domain/auth/models"
)

type Middleware struct {
	authUsecases AuthUsecases
}

func NewMiddleware(authUsecases AuthUsecases) *Middleware {
	return &Middleware{
		authUsecases: authUsecases,
	}
}

// Definir un tipo específico para la clave de contexto
type contextKey string

const claimsKey contextKey = "claims"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Missing authorization header")
			return
		}

		// Remover el prefijo "Bearer " del token
		tokenString = tokenString[len("Bearer "):]

		// Validar el token y extraer los claims
		claims, err := ValidateJWT(tokenString)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Invalid token")
			return
		}

		// Agregar los claims al contexto de la solicitud
		ctx := context.WithValue(r.Context(), claimsKey, claims)

		// Continuar con la solicitud, pasando el contexto con los claims
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (m *Middleware) RequiredPermission(permissionID int) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			claims := ExtractClaims(r.Context())
			if claims == nil {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}

			checkPermission := models.CheckPermission{
				AccountID:  claims.AccountID,
				Permission: permissionID,
			}

			hasPermission, err := m.authUsecases.HasPermission(r.Context(), checkPermission)
			if err != nil {
				log.Printf("error checking permission: %v", err)
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}

			if !hasPermission {
				http.Error(w, "forbidden", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// ExtractClaims es una función de ayuda para extraer los claims del contexto.
func ExtractClaims(ctx context.Context) *Claims {
	if claims, ok := ctx.Value(claimsKey).(*Claims); ok {
		return claims
	}
	return nil
}
