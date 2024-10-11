package user

import (
	"context"
	"sanatorioApp/internal/domain/users/entities"
)

type Repository interface {
	Autenticador
	RegisterU
	//GetU
	//UpdateU
	DeleteU
}

type Autenticador interface {
	LoginUser(ctx context.Context, lu entities.Account) (entities.Account, error)
}

type RegisterU interface {
	RegisterPatientTransaction(ctx context.Context, pu entities.PatientUser) (entities.PatientUser, error)
}

/*type GetU interface {
	GetSuperAdmins(ctx context.Context) ([]entities.SuperUser, error)
	GetSuperUserByID(ctx context.Context, userID int) (entities.SuperUser, error)
	GetDoctors(ctx context.Context) ([]entities.DoctorUser, error)
	GetDoctorByID(ctx context.Context, userID int) (entities.DoctorUser, error)
	GetPatients(ctx context.Context) ([]entities.PatientUser, error)
}*/

/*type UpdateU interface {
	UpdateSuperUser(ctx context.Context, ad entities.AdminData, userUpdate entities.SuperUser) (string, error)
	UpdateDoctor(ctx context.Context, ad entities.AdminData, doctorUpdate entities.DoctorUser) (string, error)
	UpdatePatient(ctx context.Context, ad entities.AdminData, patientUpdate entities.PatientUser) (string, error)
}*/

type DeleteU interface {
	DeleteUser(ctx context.Context, accountID string) (string, error)
	DeleteDoctor(ctx context.Context, accountID string) (string, error)
	SoftDeleteUser(ctx context.Context, accountID string) (string, error)
	SoftDeleteDoctor(ctx context.Context, accountID string) (string, error)
}
