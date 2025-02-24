package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type ClaimsConfirmation struct {
	AccountID uuid.UUID
	jwt.RegisteredClaims
}

func GenerateJWTConfirmation(accountID uuid.UUID) (string, error) {
	claims := ClaimsConfirmation{
		AccountID: accountID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			Issuer:    "sanatorio-app",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	fmt.Println("Token being generated for confirmation:")
	return tokenString, nil
}

func ValidateJWTConfirmation(tokenString string) (*ClaimsConfirmation, error) {
	claims := &ClaimsConfirmation{}

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
