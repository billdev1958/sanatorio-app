package user

import (
	"context"
	"sanatorioApp/internal/domain/users/http/models"

	"github.com/google/uuid"
)

type Usecase interface {
	//REGISTER
	RegisterSuperAdmin(ctx context.Context, request models.RegisterSuperAdminRequest) (models.UserData, error)
	RegisterAdmin(ctx context.Context, request models.RegisterAdminRequest) (models.UserData, error)
	RegisterReceptionist(ctx context.Context, request models.RegisterReceptionistRequest) (models.UserData, error)
	RegisterDoctor(ctx context.Context, request models.RegisterDoctorRequest) (models.UserData, error)
	RegisterPatient(ctx context.Context, request models.RegisterPatientRequest) (models.UserData, error)
	AccountVerification(ctx context.Context, token string) (models.Response, error)
	RegisterBeneficiary(ctx context.Context, request models.RegisterBeneficiaryRequest) (message string, err error)
	SendEmailVerification(ctx context.Context, email string) (string, error)
	CodeVerification(ctx context.Context, cd models.ConfirmationData) (string, error)
	LoginUser(ctx context.Context, lu models.LoginUser) (models.Response, error)
	// GET

	GetDoctorByID(ctx context.Context, doctorID int) (models.DoctorRequest, error)

	GetMedicalHistoryByID(ctx context.Context, md models.MedicalHistoryRequest) (models.MedicalHistoryResponse, error)

	// EDIT
	UpdatedSuperAdmin(ctx context.Context, request models.UpdateUser) (message string, err error)
	UpdatedAdmin(ctx context.Context, request models.UpdateUser) (message string, err error)
	UpdatedDoctor(ctx context.Context, request models.DoctorUpdateRequest) (message string, err error)
	UpdatedReceptionist(ctx context.Context, request models.UpdateUser) (message string, err error)
	UpdatedPatient(ctx context.Context, request models.UpdateUser) (message string, err error)

	CompleteMedicalHistory(ctx context.Context, request models.CompleteMedicalHistoryRequest) (string, error)

	// deletes
	SoftDeleteSuperAdmin(ctx context.Context, accountID uuid.UUID) (message string, err error)
	SoftDeleteAdmin(ctx context.Context, accountID uuid.UUID) (message string, err error)
	SoftDeleteDoctor(ctx context.Context, accountID uuid.UUID) (message string, err error)
	SoftDeleteReceptionist(ctx context.Context, accountID uuid.UUID) (message string, err error)
	SoftDeletePatient(ctx context.Context, accountID uuid.UUID) (message string, err error)
}
