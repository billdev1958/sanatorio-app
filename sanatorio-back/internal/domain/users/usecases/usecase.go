package usecase

import (
	"context"
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

func (u *usecase) RegisterUser(ctx context.Context, request models.RegisterUserByAdminRequest) (models.Response, error) {
	// Extraer los claims del contexto para obtener AccountID y RolAdmin
	claims := auth.ExtractClaims(ctx)
	if claims == nil {
		return models.Response{
			Status:  "error",
			Message: "Unauthorized: no claims found in context",
			Errors:  map[string]string{"authorization": "No claims found in context"},
		}, nil
	}

	adminAccountID := claims.AccountID
	adminRole := claims.Role

	// Verificar que el rol es de administrador antes de proceder
	if adminRole != 1 { // Supongamos que 1 es el rol de administrador
		return models.Response{
			Status:  "error",
			Message: "Unauthorized: insufficient permissions",
			Errors:  map[string]string{"authorization": "User does not have admin privileges"},
		}, nil
	}

	// Hashear la contraseña del nuevo usuario
	hashedPassword, err := password.HashPassword(request.Password)
	if err != nil {
		return models.Response{
			Status:  "error",
			Message: "Failed to hash password",
			Errors:  map[string]string{"password": err.Error()},
		}, err
	}

	// Crear la entidad RegisterUserByAdmin con los datos del administrador y del usuario
	registerUser := entities.RegisterUserByAdmin{
		AdminData: entities.AdminData{
			AccountAdminID: adminAccountID,
			RolAdmmin:      adminRole,
			AdminPassword:  request.AdminPassword,
		},
		User: entities.User{
			Name:      request.Name,
			Lastname1: request.Lastname1,
			Lastname2: request.Lastname2,
		},
		Account: entities.Account{
			AccountID: uuid.New(),
			Email:     request.Email,
			Password:  hashedPassword,
			Rol:       request.Rol,
		},
		DocumentID: request.Curp,
	}

	// Intentar registrar el usuario en una transacción
	userResponse, err := u.repo.RegisterUserTransaction(ctx, registerUser)
	if err != nil {
		return models.Response{
			Status:  "error",
			Message: "Failed to register user",
			Errors:  map[string]string{"register": err.Error()},
		}, err
	}

	// Retornar la respuesta exitosa
	return models.Response{
		Status:  "success",
		Message: "User registered successfully",
		Data: models.UserData{
			Name:  userResponse.Name,
			Email: userResponse.Email,
		},
	}, nil
}

func (u *usecase) RegisterDoctor(ctx context.Context, request models.RegisterDoctorByAdminRequest) (models.Response, error) {
	// Extraer los claims del contexto para obtener AccountID y RolAdmin
	claims := auth.ExtractClaims(ctx)
	if claims == nil {
		return models.Response{
			Status:  "error",
			Message: "Unauthorized: no claims found in context",
			Errors:  map[string]string{"authorization": "No claims found in context"},
		}, nil
	}

	adminAccountID := claims.AccountID
	adminRole := claims.Role

	// Verificar que el rol es de administrador antes de proceder
	if adminRole != 1 {
		return models.Response{
			Status:  "error",
			Message: "Unauthorized: insufficient permissions",
			Errors:  map[string]string{"authorization": "User does not have admin privileges"},
		}, nil
	}

	if request.Rol != 2 {
		return models.Response{
			Status:  "error",
			Message: "Invalid role for doctor",
			Errors:  map[string]string{"role": "Invalid role, expected role ID 2 for doctor"},
		}, nil
	}

	hashedPassword, err := password.HashPassword(request.Password)
	if err != nil {
		return models.Response{
			Status:  "error",
			Message: "Failed to hash password",
			Errors:  map[string]string{"password": err.Error()},
		}, err
	}

	registerDoctor := entities.RegisterDoctorByAdmin{
		AdminData: entities.AdminData{
			AccountAdminID: adminAccountID,
			RolAdmmin:      adminRole,
			AdminPassword:  request.AdminPassword,
		},
		User: entities.User{
			Name:      request.Name,
			Lastname1: request.Lastname1,
			Lastname2: request.Lastname2,
		},
		Account: entities.Account{
			AccountID: uuid.New(),
			Email:     request.Email,
			Password:  hashedPassword,
			Rol:       request.Rol,
		},
		DocumentID:  request.MedicalLicense,
		SpecialtyID: request.Specialty,
	}

	doctorResponse, err := u.repo.RegisterDoctorTransaction(ctx, registerDoctor)
	if err != nil {
		return models.Response{
			Status:  "error",
			Message: "Failed to register doctor",
			Errors:  map[string]string{"register": err.Error()},
		}, err
	}

	return models.Response{
		Status:  "success",
		Message: "Doctor registered successfully",
		Data: models.UserData{
			Name:  doctorResponse.Name,
			Email: doctorResponse.Email,
		},
	}, nil
}

func (u *usecase) RegisterPatient(ctx context.Context, request models.RegisterPatient) (models.Response, error) {
	// Definir el rol para el paciente
	rol := 3

	// Hashear la contraseña del nuevo paciente
	hashedPassword, err := password.HashPassword(request.Password)
	if err != nil {
		return models.Response{
			Status:  "error",
			Message: "Failed to hash password",
			Errors:  map[string]string{"password": err.Error()},
		}, err
	}

	// Crear la entidad PatientUser con los datos de la solicitud
	registerPatient := entities.PatientUser{
		Name:      request.Name,
		Lastname1: request.Lastname1,
		Lastname2: request.Lastname2,
		AccountID: uuid.New(), // Asignar un nuevo UUID
		Email:     request.Email,
		Password:  hashedPassword,
		Rol:       rol,
		Curp:      request.Curp, // Asignar el CURP al paciente
	}

	// Intentar registrar el paciente en una transacción
	patientResponse, err := u.repo.RegisterPatientTransaction(ctx, registerPatient)
	if err != nil {
		return models.Response{
			Status:  "error",
			Message: "Failed to register patient",
			Errors:  map[string]string{"register": err.Error()},
		}, err
	}

	// Retornar la respuesta exitosa
	return models.Response{
		Status:  "success",
		Message: "Patient registered successfully",
		Data: models.UserData{
			Name:  patientResponse.Name,
			Email: patientResponse.Email,
		},
	}, nil
}

func (u *usecase) LoginUser(ctx context.Context, request models.LoginUser) (models.Response, error) {
	loginUser := entities.LoginUser{
		Email:    request.Email,
		Password: request.Password,
	}

	loginResponse, err := u.repo.LoginUser(ctx, loginUser)
	if err != nil {
		return models.Response{
			Status:  "error",
			Message: "Failed to login user",
			Errors:  map[string]string{"register": err.Error()},
		}, err
	}

	token, err := auth.GenerateJWT(loginResponse.AccountID, loginResponse.Role)
	if err != nil {
		return models.Response{
			Status:  "error",
			Message: "Failed to generate token",
			Errors:  map[string]string{"token": err.Error()},
		}, err
	}

	return models.Response{
		Status:  "success",
		Message: "User registered successfully",
		Data: models.LoginResponse{
			AccountID: loginResponse.AccountID,
			Role:      loginResponse.Role,
			Token:     token,
		},
	}, nil

}

func (u *usecase) GetUserByID(ctx context.Context, accountID string) (models.Response, error) {
	// Llama al repositorio para obtener el usuario por ID
	userEntity, err := u.repo.GetUserByID(ctx, accountID)
	if err != nil {
		return models.Response{
			Status:  "error",
			Message: "Failed to retrieve user",
			Errors:  map[string]string{"get_user": err.Error()},
		}, err
	}

	// Mapea los datos obtenidos a la estructura de respuesta del modelo HTTP
	userData := models.User{
		ID:        userEntity.ID,
		Name:      userEntity.Name,
		Lastname1: userEntity.Lastname1,
		Lastname2: userEntity.Lastname2,
		Email:     userEntity.Email,
		Curp:      userEntity.Curp,
	}

	// Construir la respuesta exitosa
	return models.Response{
		Status:  "success",
		Message: "User retrieved successfully",
		Data:    userData,
	}, nil
}

func (u *usecase) GetUsers(ctx context.Context) ([]models.Users, error) {
	// Llama al repositorio para obtener la lista de usuarios
	users, err := u.repo.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	// Mapea los datos obtenidos a la estructura de respuesta del modelo HTTP
	var response []models.Users
	for _, user := range users {
		response = append(response, models.Users{
			User: models.User{
				ID:        user.ID,
				Name:      user.Name,
				Lastname1: user.Lastname1,
				Lastname2: user.Lastname2,
			},
			Email: user.Email,
			Curp:  user.Curp,
		})
	}

	return response, nil
}

func (u *usecase) GetDoctorByID(ctx context.Context, accountID string) (models.Response, error) {
	// Llama al repositorio para obtener el doctor por ID
	doctorEntity, err := u.repo.GetDoctorByID(ctx, accountID)
	if err != nil {
		return models.Response{
			Status:  "error",
			Message: "Failed to retrieve doctor",
			Errors:  map[string]string{"get_doctor": err.Error()},
		}, err
	}

	// Mapea los datos obtenidos a la estructura de respuesta del modelo HTTP
	doctorData := models.Doctors{
		User: models.User{
			ID:        doctorEntity.ID,
			Name:      doctorEntity.Name,
			Lastname1: doctorEntity.Lastname1,
			Lastname2: doctorEntity.Lastname2,
		},
		Email:          doctorEntity.Email,
		MedicalLicense: doctorEntity.MedicalLicense,
		Specialty:      doctorEntity.Specialty,
	}

	// Construir la respuesta exitosa
	return models.Response{
		Status:  "success",
		Message: "Doctor retrieved successfully",
		Data:    doctorData,
	}, nil
}

func (u *usecase) GetDoctors(ctx context.Context) ([]models.Doctors, error) {
	// Llama al repositorio para obtener la lista de doctores
	doctors, err := u.repo.GetDoctors(ctx)
	if err != nil {
		return nil, err
	}

	// Mapea los datos obtenidos a la estructura de respuesta del modelo HTTP
	var response []models.Doctors
	for _, doctor := range doctors {
		response = append(response, models.Doctors{
			User: models.User{
				ID:        doctor.ID,
				Name:      doctor.Name,
				Lastname1: doctor.Lastname1,
				Lastname2: doctor.Lastname2,
			},
			Email:          doctor.Email,
			MedicalLicense: doctor.MedicalLicense,
			Specialty:      doctor.Specialty,
		})
	}

	return response, nil
}

// UpdateUser actualiza la información de un usuario existente.
func (u *usecase) UpdateUser(ctx context.Context, userUpdate models.UpdateUser) (models.Response, error) {
	// Convertir de models.UpdateUser a entities.UpdateUser
	updateUser := entities.UpdateUser{
		AccountID: userUpdate.AccountID,
		Name:      userUpdate.Name,
		Lastname1: userUpdate.Lastname1,
		Lastname2: userUpdate.Lastname2,
		Email:     userUpdate.Email,
		Password:  userUpdate.Password,
		Curp:      userUpdate.Curp,
	}

	// Llama al repositorio para actualizar el usuario
	message, err := u.repo.UpdateUser(ctx, updateUser)
	if err != nil {
		return models.Response{
			Status:  "error",
			Message: "Failed to update user",
			Errors:  map[string]string{"update": err.Error()},
		}, err
	}

	return models.Response{
		Status:  "success",
		Message: message,
	}, nil
}

// UpdateDoctor actualiza la información de un doctor existente.
func (u *usecase) UpdateDoctor(ctx context.Context, doctorUpdate models.UpdateDoctor) (models.Response, error) {
	// Convertir de models.UpdateDoctor a entities.UpdateDoctor
	updateDoctorEntity := entities.UpdateDoctor{
		AccountID:      doctorUpdate.AccountID,
		Name:           doctorUpdate.Name,
		Lastname1:      doctorUpdate.Lastname1,
		Lastname2:      doctorUpdate.Lastname2,
		Email:          doctorUpdate.Email,
		Password:       doctorUpdate.Password,
		MedicalLicense: doctorUpdate.MedicalLicense,
		SpecialtyID:    doctorUpdate.SpecialtyID,
	}

	// Llama al repositorio para actualizar el doctor
	message, err := u.repo.UpdateDoctor(ctx, updateDoctorEntity)
	if err != nil {
		return models.Response{
			Status:  "error",
			Message: "Failed to update doctor",
			Errors:  map[string]string{"update": err.Error()},
		}, err
	}

	return models.Response{
		Status:  "success",
		Message: message,
	}, nil
}

func (u *usecase) DeleteUser(ctx context.Context, accountID string) (models.Response, error) {
	message, err := u.repo.DeleteUser(ctx, accountID)
	if err != nil {
		return models.Response{
			Status:  "error",
			Message: "Failed to delete user",
			Errors:  map[string]string{"delete": err.Error()},
		}, err
	}

	return models.Response{
		Status:  "success",
		Message: message,
	}, nil
}

// DeleteDoctor elimina un doctor y su cuenta asociada de la base de datos.
func (u *usecase) DeleteDoctor(ctx context.Context, accountID string) (models.Response, error) {
	message, err := u.repo.DeleteDoctor(ctx, accountID)
	if err != nil {
		return models.Response{
			Status:  "error",
			Message: "Failed to delete doctor",
			Errors:  map[string]string{"delete": err.Error()},
		}, err
	}

	return models.Response{
		Status:  "success",
		Message: message,
	}, nil
}

// SoftDeleteUser marca un usuario como eliminado sin borrar su información.
func (u *usecase) SoftDeleteUser(ctx context.Context, accountID string) (models.Response, error) {
	message, err := u.repo.SoftDeleteUser(ctx, accountID)
	if err != nil {
		return models.Response{
			Status:  "error",
			Message: "Failed to soft delete user",
			Errors:  map[string]string{"soft_delete": err.Error()},
		}, err
	}

	return models.Response{
		Status:  "success",
		Message: message,
	}, nil
}

// SoftDeleteDoctor marca un doctor como eliminado sin borrar su información.
func (u *usecase) SoftDeleteDoctor(ctx context.Context, accountID string) (models.Response, error) {
	message, err := u.repo.SoftDeleteDoctor(ctx, accountID)
	if err != nil {
		return models.Response{
			Status:  "error",
			Message: "Failed to soft delete doctor",
			Errors:  map[string]string{"soft_delete": err.Error()},
		}, err
	}

	return models.Response{
		Status:  "success",
		Message: message,
	}, nil
}
