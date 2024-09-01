package user

import (
	"context"
	"sanatorioApp/internal/domain/users/http/models"
)

type Usecase interface {
	//REGISTER
	RegisterUser(ctx context.Context, request models.RegisterUserByAdminRequest) (models.Response, error)
	RegisterDoctor(ctx context.Context, request models.RegisterDoctorByAdminRequest) (models.Response, error)
	RegisterPatient(ctx context.Context, request models.RegisterPatient) (models.Response, error)

	LoginUser(ctx context.Context, lu models.LoginUser) (models.Response, error)
	// GET
	GetUsers(ctx context.Context) ([]models.Users, error)
	GetDoctors(ctx context.Context) ([]models.Doctors, error)
	GetDoctorByID(ctx context.Context, accountID string) (models.Response, error)
	GetUserByID(ctx context.Context, accountID string) (models.Response, error)

	// EDIT
	UpdateUser(ctx context.Context, userUpdate models.UpdateUser) (models.Response, error)
	UpdateDoctor(ctx context.Context, du models.UpdateDoctor) (models.Response, error)

	// deletes
	DeleteUser(ctx context.Context, accountID string) (models.Response, error)
	DeleteDoctor(ctx context.Context, accountID string) (models.Response, error)
	SoftDeleteUser(ctx context.Context, accountID string) (models.Response, error)
	SoftDeleteDoctor(ctx context.Context, accountID string) (models.Response, error)
}
