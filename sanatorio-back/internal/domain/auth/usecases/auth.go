package usecases

import (
	"context"
	"fmt"
	"log"
	"sanatorioApp/internal/domain/auth"
	"sanatorioApp/internal/domain/auth/models"
)

type usecase struct {
	repo auth.AuthRepository
}

func NewUSecase(repo auth.AuthRepository) auth.AuthUsecases {
	return &usecase{repo: repo}
}

func (u *usecase) HasPermission(ctx context.Context, cp models.CheckPermission) (bool, error) {
	log.Printf("checking permission for account: %s", cp.AccountID)

	checkPermission, err := u.repo.HasPermission(ctx, cp.AccountID, cp.Permission)
	if err != nil {
		log.Printf("error checking permission: %v", err)
		return false, fmt.Errorf("error checking permission: %w", err)
	}

	return checkPermission, nil
}
