package app

import (
	"context"
	"sanatorioApp/internal/domain/auth"
	authR "sanatorioApp/internal/domain/auth/repository"
	authU "sanatorioApp/internal/domain/auth/usecases"
	postgres "sanatorioApp/internal/infraestructure/db"

	"github.com/jackc/pgx/v5/pgxpool"
)

func AuthService(ctx context.Context, db *pgxpool.Pool) auth.AuthUsecases {
	storage := postgres.NewPgxStorage(db)
	repo := authR.NewAuthRepository(storage)
	return authU.NewUSecase(repo)
}
