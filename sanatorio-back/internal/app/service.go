package app

import (
	"context"
	"net/http"
	v1 "sanatorioApp/internal/domain/users/http"
	"sanatorioApp/internal/domain/users/repository"
	usecase "sanatorioApp/internal/domain/users/usecases"
	postgres "sanatorioApp/internal/infraestructure/db"

	"github.com/jackc/pgx/v5/pgxpool"
)

func UserService(ctx context.Context, db *pgxpool.Pool, router *http.ServeMux) error {
	storage := postgres.NewPgxStorage(db)

	repo := repository.NewUserRepository(storage)

	uc := usecase.NewUsecase(repo)

	h := v1.NewHandler(uc)

	h.UserRoutes(router)

	return nil
}
