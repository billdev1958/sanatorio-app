package user

import (
	"context"
	"sanatorioApp/internal/domain/users/http/models"
)

type Usecase interface {
	//REGISTER
	RegisterSuperAdmin(ctx context.Context, request models.RegisterSuperAdminRequest) (models.UserData, error)
	RegisterReceptionist(ctx context.Context, request models.RegisterReceptionistRequest) (models.UserData, error)
	RegisterDoctor(ctx context.Context, request models.RegisterDoctorRequest) (models.UserData, error)
	RegisterPatient(ctx context.Context, request models.RegisterPatientRequest) (models.UserData, error)

	RegisterBeneficiary(ctx context.Context, request models.RegisterBeneficiaryRequest) (message string, err error)

	LoginUser(ctx context.Context, lu models.LoginUser) (models.LoginResponse, error)
	// GET

	GetDoctorByID(ctx context.Context, doctorID int) (models.DoctorRequest, error)

	GetMedicalHistoryByID(ctx context.Context, md models.MedicalHistoryRequest) (models.MedicalHistoryResponse, error)

	// EDIT
	UpdateUser(ctx context.Context, userUpdate models.UpdateUser) (string, error)
	UpdatedDoctor(ctx context.Context, request models.DoctorUpdateRequest) (message string, err error)

	CompleteMedicalHistory(ctx context.Context, request models.CompleteMedicalHistoryRequest) (string, error)

	// deletes
	DeleteUser(ctx context.Context, accountID string) (string, error)
	SoftDeleteUser(ctx context.Context, accountID string) (string, error)
}
