package app

import (
	"context"
	"net/http"
	v1 "sanatorioApp/internal/domain/cites/http"
	"sanatorioApp/internal/domain/cites/repository"
	usecase "sanatorioApp/internal/domain/cites/usecases"
	postgres "sanatorioApp/internal/infraestructure/db"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CitesService(ctx context.Context, db *pgxpool.Pool, router *http.ServeMux) error {
	storage := postgres.NewPgxStorage(db)

	repo := repository.NewCitesRepository(storage)

	uc := usecase.NewUsecase(repo)

	h := v1.NewHandler(uc)

	h.CitesRoutes(router)

	return nil
}
