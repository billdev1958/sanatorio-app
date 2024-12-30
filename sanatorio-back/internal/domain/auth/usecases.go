package auth

import (
	"context"
	"sanatorioApp/internal/domain/auth/models"
)

type AuthUsecases interface {
	HasPermission(ctx context.Context, cp models.CheckPermission) (bool, error)
}
