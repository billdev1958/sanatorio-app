package repository

import (
	"context"
	"fmt"
	"sanatorioApp/internal/domain/users/entities"
)

func (ur *userRepository) GetUserByIdentifier(ctx context.Context, identifier string) (entities.Account, error) {
	query := "SELECT id, email, password, role_id FROM account WHERE email = $1" // Incluye role_id

	var account entities.Account
	err := ur.storage.DbPool.QueryRow(ctx, query, identifier).Scan(&account.ID, &account.Email, &account.Password, &account.Rol)
	if err != nil {
		return account, fmt.Errorf("failed to find user: %w", err)
	}

	return account, nil
}
