package user

import (
	"context"
	"sanatorioApp/internal/domain/users/entities"

	"github.com/jackc/pgx/v5"
)

type Repository interface {
	RegisterUserTransaction(ctx context.Context, ru entities.RegisterUser) (entities.UserResponse, error)
	RegisterDoctorTransaction(ctx context.Context, rd entities.RegisterDoctor) (response entities.UserResponse, err error)

	RegisterUser(ctx context.Context, tx pgx.Tx, ru entities.RegisterUser) (userID int, name string, err error)
	RegisterAccount(ctx context.Context, tx pgx.Tx, ru entities.RegisterUser) (email string, err error)
	RegisterTypeUser(ctx context.Context, tx pgx.Tx, ru entities.RegisterUser) error
	RegisterTypeDoctor(ctx context.Context, tx pgx.Tx, ru entities.RegisterDoctor) error
}
