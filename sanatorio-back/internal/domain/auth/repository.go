package auth

import (
	"context"

	"github.com/google/uuid"
)

type AuthRepository interface {
	HasPermission(ctx context.Context, accountID uuid.UUID, permission int) (bool, error)
}
