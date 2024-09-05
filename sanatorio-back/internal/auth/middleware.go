package auth

import (
	"context"
	"fmt"
	"net/http"
)

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

// ExtractClaims es una función de ayuda para extraer los claims del contexto.
func ExtractClaims(ctx context.Context) *Claims {
	if claims, ok := ctx.Value(claimsKey).(*Claims); ok {
		return claims
	}
	return nil
}
