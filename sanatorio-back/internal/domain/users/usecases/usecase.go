package usecase

import (
	"context"
	"fmt"
	"sanatorioApp/internal/auth"
	user "sanatorioApp/internal/domain/users"
	"sanatorioApp/internal/domain/users/entities"
	"sanatorioApp/internal/domain/users/http/models"
	password "sanatorioApp/pkg/pass"

	"github.com/google/uuid"
)

type usecase struct {
	repo user.Repository
}

func NewUsecase(repo user.Repository) user.Usecase {
	return &usecase{repo: repo}
}

func (u *usecase) RegisterPatient(ctx context.Context, request models.RegisterPatientRequest) (models.UserData, error) {
	// Hashear la contraseña del nuevo paciente
	hashedPassword, err := password.HashPassword(request.Password)
	if err != nil {
		return models.UserData{}, fmt.Errorf("failed to hash password: %w", err)
	}

	// Crear la entidad PatientUser con los datos de la solicitud
	registerPatient := entities.PatientUser{
		FirstName: request.Name,
		LastName1: request.Lastname1,
		LastName2: request.Lastname2,

		Account: entities.Account{
			AccountID: uuid.New(), // Asignar un nuevo UUID
			Email:     request.Email,
			Password:  hashedPassword,
			Rol:       entities.Patient,
		},
		Curp: request.Curp, // Asignar el CURP al paciente
	}

	// Intentar registrar al paciente en una transacción
	patientResponse, err := u.repo.RegisterPatientTransaction(ctx, registerPatient)
	if err != nil {
		return models.UserData{}, fmt.Errorf("failed to register patient: %w", err)
	}

	// Retornar los datos del paciente registrado
	return models.UserData{
		Name:  patientResponse.FirstName,
		Email: patientResponse.Email,
	}, nil
}

func (u *usecase) LoginUser(ctx context.Context, request models.LoginUser) (models.LoginResponse, error) {
	// Crear la entidad de login
	loginUser := entities.Account{
		Email:    request.Email,
		Password: request.Password,
	}

	// Llamar al repositorio para autenticar el usuario
	loginResponse, err := u.repo.LoginUser(ctx, loginUser)
	if err != nil {
		return models.LoginResponse{}, err
	}

	// Generar el token JWT si el login fue exitoso
	token, err := auth.GenerateJWT(loginResponse.AccountID, int(loginResponse.Rol))
	if err != nil {
		return models.LoginResponse{}, err
	}

	// Retornar los datos crudos (LoginResponse) al handler
	return models.LoginResponse{
		AccountID: loginResponse.AccountID,
		Role:      int(loginResponse.Rol),
		Token:     token,
	}, nil
}
