package repository

import (
	"context"
	"fmt"
	"sanatorioApp/internal/domain/users/entities"
)

func (ur *userRepository) GetUserByIdentifier(ctx context.Context, identifier string) (entities.Account, error) {
	query := "SELECT email, password FROM account WHERE email = $1"

	var account entities.Account
	err := ur.storage.DbPool.QueryRow(ctx, query, identifier).Scan(&account.Email, &account.Password)
	if err != nil {
		return account, fmt.Errorf("failed to find user: %w", err)
	}

	return account, nil
}
