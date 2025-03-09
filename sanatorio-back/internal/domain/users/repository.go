package user

import (
	"context"
	"sanatorioApp/internal/domain/users/entities"
	"time"

	"github.com/google/uuid"
)

type Repository interface {
	Autenticador
	RegisterU
	GetU
	UpdateU
	DeleteU
	Verification
}

type Autenticador interface {
	GetUserByIdentifier(ctx context.Context, identifier string) (entities.Account, error)
	AccountVerification(ctx context.Context, accountID uuid.UUID, isValid bool) (bool, error)
}

type RegisterU interface {
	RegisterSuperAdminTransaction(ctx context.Context, account entities.Account, su entities.SuperAdminUser) (entities.SuperAdminUser, error)
	RegisterAdminTransaction(ctx context.Context, account entities.Account, au entities.AdminUser) (entities.AdminUser, error)
	RegisterReceptionistTransaction(ctx context.Context, account entities.Account, ru entities.ReceptionistUser) (entities.ReceptionistUser, error)
	RegisterDoctorTransaction(ctx context.Context, account entities.Account, du entities.DoctorUser) (entities.DoctorUser, error)
	RegisterPatientTransaction(ctx context.Context, account entities.Account, pu entities.PatientUser) (entities.PatientUser, error)
	RegisterBeneficiary(ctx context.Context, request entities.BeneficiaryUser) (message string, err error)
}

type Verification interface {
	SaveCodeVerification(ctx context.Context, email, code string, expired_at time.Time) error
	VerifyCode(ctx context.Context, code, email string) (bool, error)
}

type GetU interface {
	GetDoctorByID(ctx context.Context, userID int) (entities.DoctorUser, error)
	GetMedicalHistoryByID(ctx context.Context, MedicalHistoryID string) (entities.MedicalHistory, error)
}

type UpdateU interface {
	UpdateSuperAdmin(ctx context.Context, d entities.SuperAdminUser) (message string, err error)
	UpdateAdmin(ctx context.Context, d entities.AdminUser) (message string, err error)
	UpdateDoctor(ctx context.Context, d entities.DoctorUser) (message string, err error)
	UpdateReceptionist(ctx context.Context, d entities.ReceptionistUser) (message string, err error)
	UpdatePatient(ctx context.Context, d entities.PatientUser) (message string, err error)
	CompleteMedicalHistory(ctx context.Context, md entities.MedicalHistory) (bool, error)
}

type DeleteU interface {
	DeleteUser(ctx context.Context, accountID string) (string, error)
	SoftDeleteUserSuperAdmin(ctx context.Context, a entities.Account) (bool, error)
	SoftDeleteUserAdmin(ctx context.Context, au entities.Account) (bool, error)
	SoftDeleteUserDoctor(ctx context.Context, a entities.Account) (bool, error)
	SoftDeleteUserReceptionist(ctx context.Context, a entities.Account) (bool, error)
	SoftDeleteUserPatient(ctx context.Context, a entities.Account) (bool, error)
}
