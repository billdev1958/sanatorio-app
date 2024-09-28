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

func (u *usecase) RegisterSuperUser(ctx context.Context, request models.RegisterUserByAdminRequest) (models.UserData, error) {
	// Extraer los claims del contexto para obtener AccountID y RolAdmin
	claims := auth.ExtractClaims(ctx)
	if claims == nil {
		return models.UserData{}, fmt.Errorf("unauthorized: no claims found in context")
	}

	adminAccountID := claims.AccountID
	adminRole := claims.Role

	fmt.Println(adminAccountID)
	fmt.Println(adminRole)

	// Verificar que el rol es de administrador antes de proceder
	if adminRole != entities.SuperUsuario {
		return models.UserData{}, fmt.Errorf("unauthorized: insufficient permissions")
	}

	// Hashear la contraseña del nuevo usuario
	hashedPassword, err := password.HashPassword(request.Password)
	if err != nil {
		return models.UserData{}, fmt.Errorf("failed to hash password: %w", err)
	}

	// Datos del administrador
	adminData := entities.AdminData{
		AccountID:     adminAccountID,
		RoleAdmin:     adminRole,
		PasswordAdmin: request.AdminPassword,
	}

	// Crear la estructura del usuario a registrar
	registerUser := entities.SuperUser{
		User: entities.User{
			Name:      request.Name,
			Lastname1: request.Lastname1,
			Lastname2: request.Lastname2,
		},
		Account: entities.Account{
			AccountID: uuid.New(),
			Email:     request.Email,
			Password:  hashedPassword,
			Rol:       entities.SuperUsuario,
		},
		Curp: request.Curp,
	}

	// Intentar registrar el usuario en una transacción
	userResponse, err := u.repo.RegisterSuperUserTransaction(ctx, adminData, registerUser)
	if err != nil {
		return models.UserData{}, fmt.Errorf("failed to register user: %w", err)
	}

	// Retornar los datos del usuario registrado
	return models.UserData{
		Name:  userResponse.Name,
		Email: userResponse.Email,
	}, nil
}

func (u *usecase) RegisterDoctor(ctx context.Context, request models.RegisterDoctorByAdminRequest) (models.UserData, error) {
	// Extraer los claims del contexto para obtener AccountID y RolAdmin
	claims := auth.ExtractClaims(ctx)
	if claims == nil {
		return models.UserData{}, fmt.Errorf("unauthorized: no claims found in context")
	}

	adminAccountID := claims.AccountID
	adminRole := claims.Role

	// Verificar que el rol es de administrador antes de proceder
	if adminRole != 1 { // Supongamos que 1 es el rol de administrador
		return models.UserData{}, fmt.Errorf("unauthorized: insufficient permissions")
	}

	// Verificar que el rol proporcionado es de doctor (Rol = 2)
	if request.Rol != 2 {
		return models.UserData{}, fmt.Errorf("invalid role for doctor, expected role ID 2 for doctor")
	}

	// Hashear la contraseña del nuevo doctor
	hashedPassword, err := password.HashPassword(request.Password)
	if err != nil {
		return models.UserData{}, fmt.Errorf("failed to hash password: %w", err)
	}

	adminData := entities.AdminData{
		AccountID: adminAccountID,
		RoleAdmin: adminRole,
	}

	// Crear la estructura del doctor a registrar
	registerDoctor := entities.DoctorUser{
		User: entities.User{
			Name:      request.Name,
			Lastname1: request.Lastname1,
			Lastname2: request.Lastname2,
		},
		Account: entities.Account{
			AccountID: uuid.New(), // Generar un nuevo UUID para el doctor
			Email:     request.Email,
			Password:  hashedPassword,
			Rol:       entities.Doctor,
		},
		MedicalLicense: request.MedicalLicense,                  // Asignar el número de licencia médica
		SpecialtyID:    entities.Specialties(request.Specialty), // Asignar el ID de la especialidad médica
	}

	// Intentar registrar al doctor en una transacción
	doctorResponse, err := u.repo.RegisterDoctorTransaction(ctx, adminData, registerDoctor)
	if err != nil {
		return models.UserData{}, fmt.Errorf("failed to register doctor: %w", err)
	}

	// Retornar los datos del doctor registrado
	return models.UserData{
		Name:  doctorResponse.Name,
		Email: doctorResponse.Email,
	}, nil
}

func (u *usecase) RegisterPatient(ctx context.Context, request models.RegisterPatientRequest) (models.UserData, error) {
	// Hashear la contraseña del nuevo paciente
	hashedPassword, err := password.HashPassword(request.Password)
	if err != nil {
		return models.UserData{}, fmt.Errorf("failed to hash password: %w", err)
	}

	// Crear la entidad PatientUser con los datos de la solicitud
	registerPatient := entities.PatientUser{
		User: entities.User{
			Name:      request.Name,
			Lastname1: request.Lastname1,
			Lastname2: request.Lastname2,
		},
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
		Name:  patientResponse.Name,
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
