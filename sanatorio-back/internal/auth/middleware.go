package auth

import (
	"context"
	"net/http"
	"strings"
)

// Definir un tipo específico para la clave de contexto
type contextKey string

const claimsKey contextKey = "claims"

// AuthMiddleware protege las rutas verificando el JWT en las peticiones.
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Obtener el token del encabezado Authorization
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		// Eliminar el prefijo "Bearer " del token si está presente
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == "" {
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		// Validar el JWT
		claims, err := ValidateJWT(tokenString)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Añadir los claims al contexto para su uso en el siguiente handler
		ctx := context.WithValue(r.Context(), claimsKey, claims)
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
