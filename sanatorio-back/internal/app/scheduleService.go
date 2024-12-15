package app

import (
	"context"
	"net/http"
	"sanatorioApp/internal/domain/catalogs/repository"
	v1 "sanatorioApp/internal/domain/schedules/http"
	"sanatorioApp/internal/domain/schedules/usecases"
	postgres "sanatorioApp/internal/infraestructure/db"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ScheduleService(ctx context.Context, db *pgxpool.Pool, router *http.ServeMux) error {
	storage := postgres.NewPgxStorage(db)

	repo := repository.NewCatalogRepository(storage)

	uc := usecases.NewUsecase(repo)

	h := v1.NewScheduleHandler(uc)

	h.ScheduleRouter(router)

	return nil
}
