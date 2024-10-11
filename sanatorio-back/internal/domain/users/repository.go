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
	LoginUser(ctx context.Context, lu entities.Account) (entities.Account, error)
}

type RegisterU interface {
	RegisterPatientTransaction(ctx context.Context, pu entities.PatientUser) (entities.PatientUser, error)
}

type GetU interface {
	GetDoctorByID(ctx context.Context, userID int) (entities.DoctorUser, error)
}

type UpdateU interface {
	UpdatePatient(ctx context.Context, patientUpdate entities.PatientUser) (string, error)
}

type DeleteU interface {
	DeleteUser(ctx context.Context, accountID string) (string, error)
	SoftDeleteUser(ctx context.Context, accountID string) (string, error)
}
