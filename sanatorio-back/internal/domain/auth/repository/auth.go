package repository

import (
	"context"
	"fmt"
	"log"
	"sanatorioApp/internal/domain/auth"
	postgres "sanatorioApp/internal/infraestructure/db"

	"github.com/google/uuid"
)

type authRepository struct {
	storage *postgres.PgxStorage
}

func NewAuthRepository(storage *postgres.PgxStorage) auth.AuthRepository {
	return &authRepository{storage: storage}
}

func (ar *authRepository) HasPermission(ctx context.Context, accountID uuid.UUID, permission int) (bool, error) {

	query := `SELECT EXISTS(
		SELECT 1
		FROM account a
		JOIN role_permission rp ON a.role_id = rp.role_id
		WHERE a.id = $1
		AND rp.permission_id = $2
		AND a.deleted_at IS NULL
		AND rp.deleted_at IS NULL
	)`

	var hasPermission bool

	err := ar.storage.DbPool.QueryRow(ctx, query, accountID, permission).Scan(&hasPermission)
	if err != nil {
		log.Printf("error checking permission: account_id=%s, permission=%d, error=%v",
			accountID, permission, err)
		return false, fmt.Errorf("error checking permission: %v", err)
	}

	log.Printf("permission check result: account_id=%s, permission=%d, has_permission=%v",
		accountID, permission, hasPermission)

	return hasPermission, nil

}
