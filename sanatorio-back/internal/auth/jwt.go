package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// TODO AGREGAR A ENV
var secretKey []byte

func init() {
	secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))
	if len(secretKey) == 0 {
		panic("JWT_SECRET_KEY is not set in environment variables")
	}
}

type Claims struct {
	AccountID uuid.UUID
	Role      int
	jwt.RegisteredClaims
}

func GenerateJWT(accountID uuid.UUID, role int) (string, error) {
	claims := Claims{
		AccountID: accountID,
		Role:      role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(8 * time.Hour)),
			Issuer:    "sanatorio-app",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	fmt.Println("Token being generated:", tokenString)
	return tokenString, nil
}

func ValidateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Validar que el método de firma sea HMAC.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		fmt.Println("Error parsing token:", err)
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if !token.Valid {
		fmt.Println("Token is invalid:", tokenString)
		return nil, fmt.Errorf("invalid token: %w", jwt.ErrSignatureInvalid)
	}

	fmt.Println("Token is valid, claims:", claims)
	return claims, nil
}
