package user

import (
	"context"
	"sanatorioApp/internal/domain/users/entities"
)

type Repository interface {
	Autenticador
	RegisterU
	GetU
	UpdateU
	DeleteU
}

type Autenticador interface {
	GetUserByIdentifier(ctx context.Context, identifier string) (entities.Account, error)
}

type RegisterU interface {
	RegisterAdminTransaction(ctx context.Context, account entities.Account, su entities.SuperAdminUser) (entities.SuperAdminUser, error)
	RegisterReceptionistTransaction(ctx context.Context, account entities.Account, ru entities.ReceptionistUser) (entities.ReceptionistUser, error)
	RegisterDoctorTransaction(ctx context.Context, account entities.Account, du entities.DoctorUser) (entities.DoctorUser, error)
	RegisterPatientTransaction(ctx context.Context, account entities.Account, pu entities.PatientUser) (entities.PatientUser, error)
	RegisterBeneficiary(ctx context.Context, request entities.BeneficiaryUser) (message string, err error)
}

type GetU interface {
	GetDoctorByID(ctx context.Context, userID int) (entities.DoctorUser, error)
	GetMedicalHistoryByID(ctx context.Context, MedicalHistoryID string) (entities.MedicalHistory, error)
}

type UpdateU interface {
	UpdateDoctor(ctx context.Context, d entities.DoctorUser) (message string, err error)
	UpdatePatient(ctx context.Context, patientAccount entities.Account, patientUpdate entities.PatientUser) (string, error)
	CompleteMedicalHistory(ctx context.Context, md entities.MedicalHistory) (bool, error)
}

type DeleteU interface {
	DeleteUser(ctx context.Context, accountID string) (string, error)
	SoftDeleteUser(ctx context.Context, accountID string) (string, error)
}
