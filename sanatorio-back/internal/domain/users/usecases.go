package user

import (
	"context"
	"sanatorioApp/internal/domain/users/http/models"
)

type Usecase interface {
	//REGISTER
	RegisterPatient(ctx context.Context, request models.RegisterPatientRequest) (models.UserData, error)

	LoginUser(ctx context.Context, lu models.LoginUser) (models.LoginResponse, error)
	// GET

	GetDoctorByID(ctx context.Context, doctorID int) (models.DoctorRequest, error)

	// EDIT
	UpdateUser(ctx context.Context, userUpdate models.UpdateUser) (string, error)

	// deletes
	DeleteUser(ctx context.Context, accountID string) (string, error)
	SoftDeleteUser(ctx context.Context, accountID string) (string, error)
}
