package user

import (
	"context"
	"sanatorioApp/internal/domain/users/http/models"
)

type Usecase interface {
	//REGISTER
	RegisterSuperUser(ctx context.Context, request models.RegisterUserByAdminRequest) (models.UserData, error)
	RegisterDoctor(ctx context.Context, request models.RegisterDoctorByAdminRequest) (models.UserData, error)
	RegisterPatient(ctx context.Context, request models.RegisterPatientRequest) (models.UserData, error)

	LoginUser(ctx context.Context, lu models.LoginUser) (models.LoginResponse, error)
	// GET
	GetSuperAdmins(ctx context.Context) ([]models.UserRequest, error)
	GetSuperAdminByID(ctx context.Context, superUserID int) (models.UserRequest, error)

	GetDoctors(ctx context.Context) ([]models.DoctorRequest, error)
	GetDoctorByID(ctx context.Context, doctorID int) (models.DoctorRequest, error)

	// EDIT
	UpdateUser(ctx context.Context, userUpdate models.UpdateUser) (string, error)
	UpdateDoctor(ctx context.Context, du models.UpdateDoctor) (string, error)

	// deletes
	DeleteUser(ctx context.Context, accountID string) (string, error)
	DeleteDoctor(ctx context.Context, accountID string) (string, error)
	SoftDeleteUser(ctx context.Context, accountID string) (string, error)
	SoftDeleteDoctor(ctx context.Context, accountID string) (string, error)
}
