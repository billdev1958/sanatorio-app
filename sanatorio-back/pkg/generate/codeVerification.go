package generate

import (
	"crypto/rand"
	"math/big"
)

func GenerateCode(length int) (string, error) {
	const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	code := make([]byte, length)
	for i := range code {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		code[i] = letters[num.Int64()]
	}
	return string(code), nil
}
