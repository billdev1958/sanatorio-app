package user

import (
	"context"
	"sanatorioApp/internal/domain/users/entities"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type Repository interface {
	// Login
	LoginUser(ctx context.Context, lu entities.LoginUser) (entities.LoginResponse, error)

	// Registros
	RegisterUserTransaction(ctx context.Context, ru entities.RegisterUserByAdmin) (entities.UserResponse, error)
	RegisterDoctorTransaction(ctx context.Context, rd entities.RegisterDoctorByAdmin) (response entities.UserResponse, err error)
	RegisterPatientTransaction(ctx context.Context, rp entities.PatientUser) (response entities.UserResponse, err error)
	RegisterUser(ctx context.Context, tx pgx.Tx, ru entities.RegisterUserByAdmin) (userID int, name string, err error)
	RegisterAccount(ctx context.Context, tx pgx.Tx, ru entities.RegisterUserByAdmin, userID int) (string, error)
	RegisterTypeUser(ctx context.Context, tx pgx.Tx, ru entities.RegisterUserByAdmin) error
	RegisterTypeDoctor(ctx context.Context, tx pgx.Tx, ru entities.RegisterDoctorByAdmin) error

	// GetUsers
	GetUsers(ctx context.Context) ([]entities.Users, error)
	GetDoctors(ctx context.Context) ([]entities.Doctors, error)
	GetDoctorByID(ctx context.Context, accountID string) (entities.Doctors, error)
	GetUserByID(ctx context.Context, accountID string) (entities.Users, error)

	// Edit Users
	UpdateUser(ctx context.Context, userUpdate entities.UpdateUser) (string, error)
	UpdateDoctor(ctx context.Context, du entities.UpdateDoctor) (string, error)

	CheckAdminPassword(ctx context.Context, tx pgx.Tx, accountID uuid.UUID, pass string) (bool, error)

	// Delete and soft delete
	DeleteUser(ctx context.Context, accountID string) (string, error)
	DeleteDoctor(ctx context.Context, accountID string) (string, error)
	SoftDeleteUser(ctx context.Context, accountID string) (string, error)
	SoftDeleteDoctor(ctx context.Context, accountID string) (string, error)
}
