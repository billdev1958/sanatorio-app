package usecase

import (
	"context"
	"fmt"
	"sanatorioApp/internal/auth"
	"sanatorioApp/internal/domain/users/entities"
	"sanatorioApp/internal/domain/users/http/models"
)

func (u *usecase) UpdateUser(ctx context.Context, userUpdate models.UpdateUser) (string, error) {
	claims := auth.ExtractClaims(ctx)
	if claims == nil {
		return "", fmt.Errorf("unauthorized: no claims found in context")
	}

	adminAccountID := claims.AccountID
	adminRole := claims.Role

	if adminRole != 1 {
		return "", fmt.Errorf("unauthorized: insufficient permissions")
	}

	adminData := entities.AdminData{
		AccountID: adminAccountID,
		RoleAdmin: adminRole,
	}

	updateUser := entities.SuperUser{
		User: entities.User{
			Name:      userUpdate.Name,
			Lastname1: userUpdate.Lastname1,
			Lastname2: userUpdate.Lastname2,
		},
		Account: entities.Account{
			Email:    userUpdate.Email,
			Password: userUpdate.Password,
		},
		Curp: userUpdate.Curp,
	}

	message, err := u.repo.UpdateSuperUser(ctx, adminData, updateUser)
	if err != nil {
		return "", fmt.Errorf("failed to update user: %w", err)
	}

	return message, nil
}

func (u *usecase) UpdateDoctor(ctx context.Context, doctorUpdate models.UpdateDoctor) (string, error) {
	claims := auth.ExtractClaims(ctx)
	if claims == nil {
		return "", fmt.Errorf("unauthorized: no claims found in context")
	}

	adminAccountID := claims.AccountID
	adminRole := claims.Role

	if adminRole != 1 {
		return "", fmt.Errorf("unauthorized: insufficient permissions")
	}

	adminData := entities.AdminData{
		AccountID: adminAccountID,
		RoleAdmin: adminRole,
	}

	updateDoctorEntity := entities.DoctorUser{
		User: entities.User{
			Name:      doctorUpdate.Name,
			Lastname1: doctorUpdate.Lastname1,
			Lastname2: doctorUpdate.Lastname2,
		},
		Account: entities.Account{
			Email:    doctorUpdate.Email,
			Password: doctorUpdate.Password,
		},
		MedicalLicense: doctorUpdate.MedicalLicense,
		SpecialtyID:    entities.Specialties(doctorUpdate.SpecialtyID),
	}

	message, err := u.repo.UpdateDoctor(ctx, adminData, updateDoctorEntity)
	if err != nil {
		return "", fmt.Errorf("failed to update doctor: %w", err)
	}

	return message, nil
}
