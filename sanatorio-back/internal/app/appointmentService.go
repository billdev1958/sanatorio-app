package app

import (
	"context"
	"net/http"
	v1 "sanatorioApp/internal/domain/appointment/http"
	"sanatorioApp/internal/domain/appointment/repository"
	catalogRepo "sanatorioApp/internal/domain/catalogs/repository"

	"sanatorioApp/internal/domain/appointment/usecases"

	postgres "sanatorioApp/internal/infraestructure/db"

	"github.com/jackc/pgx/v5/pgxpool"
)

func AppointmentService(ctx context.Context, db *pgxpool.Pool, router *http.ServeMux) error {
	storage := postgres.NewPgxStorage(db)

	repo := repository.NewAppointmentRepository(storage)
	cat := catalogRepo.NewCatalogRepository(storage)
	uc := usecases.NewUsecase(repo, cat)
	authUc := AuthService(ctx, db)

	h := v1.NewHandler(uc, authUc)
	h.AppointmentRouter(router)

	return nil
}
