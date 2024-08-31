package usecase_test

import (
	"context"
	"errors"
	"sanatorioApp/internal/domain/users/entities"
	"sanatorioApp/internal/domain/users/http/models"
	usecase "sanatorioApp/internal/domain/users/usecases"
	"testing"

	"github.com/jackc/pgx/v5"
)

// MockRepository es un mock simple del repositorio que se utilizará en las pruebas
type MockRepository struct {
	RegisterUserFunc       func(ctx context.Context, ru entities.RegisterUser) (entities.UserResponse, error)
	RegisterDoctorFunc     func(ctx context.Context, rd entities.RegisterDoctor) (entities.UserResponse, error)
	RegisterAccountFunc    func(ctx context.Context, tx pgx.Tx, ru entities.RegisterUser) (string, error)
	RegisterUserTxFunc     func(ctx context.Context, tx pgx.Tx, ru entities.RegisterUser) (int, string, error)
	RegisterTypeUserFunc   func(ctx context.Context, tx pgx.Tx, ru entities.RegisterUser) error
	RegisterTypeDoctorFunc func(ctx context.Context, tx pgx.Tx, rd entities.RegisterDoctor) error
}

func (m *MockRepository) RegisterUserTransaction(ctx context.Context, ru entities.RegisterUser) (entities.UserResponse, error) {
	if m.RegisterUserFunc != nil {
		return m.RegisterUserFunc(ctx, ru)
	}
	return entities.UserResponse{}, nil
}

func (m *MockRepository) RegisterDoctorTransaction(ctx context.Context, rd entities.RegisterDoctor) (entities.UserResponse, error) {
	if m.RegisterDoctorFunc != nil {
		return m.RegisterDoctorFunc(ctx, rd)
	}
	return entities.UserResponse{}, nil
}

func (m *MockRepository) RegisterAccount(ctx context.Context, tx pgx.Tx, ru entities.RegisterUser) (string, error) {
	if m.RegisterAccountFunc != nil {
		return m.RegisterAccountFunc(ctx, tx, ru)
	}
	return "", nil
}

func (m *MockRepository) RegisterUser(ctx context.Context, tx pgx.Tx, ru entities.RegisterUser) (int, string, error) {
	if m.RegisterUserTxFunc != nil {
		return m.RegisterUserTxFunc(ctx, tx, ru)
	}
	return 0, "", nil
}

func (m *MockRepository) RegisterTypeUser(ctx context.Context, tx pgx.Tx, ru entities.RegisterUser) error {
	if m.RegisterTypeUserFunc != nil {
		return m.RegisterTypeUserFunc(ctx, tx, ru)
	}
	return nil
}

func (m *MockRepository) RegisterTypeDoctor(ctx context.Context, tx pgx.Tx, rd entities.RegisterDoctor) error {
	if m.RegisterTypeDoctorFunc != nil {
		return m.RegisterTypeDoctorFunc(ctx, tx, rd)
	}
	return nil
}

func TestRegisterUser(t *testing.T) {
	mockRepo := &MockRepository{}
	usecase := usecase.NewUsecase(mockRepo)

	t.Run("Success", func(t *testing.T) {
		// Configurar el mock para retornar un resultado exitoso
		mockRepo.RegisterUserFunc = func(ctx context.Context, ru entities.RegisterUser) (entities.UserResponse, error) {
			return entities.UserResponse{
				Name:  ru.Name,
				Email: ru.Email,
			}, nil
		}

		mockRepo.RegisterAccountFunc = func(ctx context.Context, tx pgx.Tx, ru entities.RegisterUser) (string, error) {
			return "john.doe@example.com", nil
		}

		// Preparar los valores de prueba
		request := models.RegisterUserRequest{
			Name:      "John",
			Lastname1: "Doe",
			Lastname2: "Smith",
			Email:     "john.doe@example.com",
			Password:  "password",
			Rol:       1,
			Curp:      "CURP12345678901234", // 18 caracteres
		}

		// Llamar a la función que se está probando
		response, err := usecase.RegisterUser(context.Background(), request)

		// Validar que no hubo error
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if response.Status != "success" {
			t.Fatalf("expected success status, got %v", response.Status)
		}
		if response.Data.(models.UserData).Name != "John" {
			t.Fatalf("expected name to be John, got %v", response.Data.(models.UserData).Name)
		}
		if response.Data.(models.UserData).Email != "john.doe@example.com" {
			t.Fatalf("expected email to be john.doe@example.com, got %v", response.Data.(models.UserData).Email)
		}
	})

	t.Run("Fail_RegisterUserTransaction", func(t *testing.T) {
		// Configurar el mock para retornar un error
		mockRepo.RegisterUserFunc = func(ctx context.Context, ru entities.RegisterUser) (entities.UserResponse, error) {
			return entities.UserResponse{}, errors.New("insert user error")
		}

		// Preparar los valores de prueba
		request := models.RegisterUserRequest{
			Name:      "John",
			Lastname1: "Doe",
			Lastname2: "Smith",
			Email:     "john.doe@example.com",
			Password:  "password",
			Rol:       1,
			Curp:      "CURP12345678901234", // 18 caracteres
		}

		// Llamar a la función que se está probando
		response, err := usecase.RegisterUser(context.Background(), request)

		// Validar que hubo un error
		if err == nil {
			t.Fatalf("expected error, got none")
		}
		if response.Status != "error" {
			t.Fatalf("expected error status, got %v", response.Status)
		}
		if _, ok := response.Errors.(map[string]string)["register"]; !ok {
			t.Fatalf("expected register error, got none")
		}
	})
}

func TestRegisterDoctor(t *testing.T) {
	mockRepo := &MockRepository{}
	usecase := usecase.NewUsecase(mockRepo)

	t.Run("Success", func(t *testing.T) {
		// Configurar el mock para retornar un resultado exitoso
		mockRepo.RegisterDoctorFunc = func(ctx context.Context, rd entities.RegisterDoctor) (entities.UserResponse, error) {
			return entities.UserResponse{
				Name:  rd.Name,
				Email: rd.Email,
			}, nil
		}

		mockRepo.RegisterAccountFunc = func(ctx context.Context, tx pgx.Tx, rd entities.RegisterUser) (string, error) {
			return "jane.doe@example.com", nil
		}

		// Preparar los valores de prueba
		request := models.RegisterDoctorRequest{
			Name:           "Jane",
			Lastname1:      "Doe",
			Lastname2:      "Smith",
			Email:          "jane.doe@example.com",
			Password:       "password",
			Rol:            2,
			Specialty:      1,
			MedicalLicense: "LIC1234567", // 10 caracteres
		}

		// Llamar a la función que se está probando
		response, err := usecase.RegisterDoctor(context.Background(), request)

		// Validar que no hubo error
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if response.Status != "success" {
			t.Fatalf("expected success status, got %v", response.Status)
		}
		if response.Data.(models.UserData).Name != "Jane" {
			t.Fatalf("expected name to be Jane, got %v", response.Data.(models.UserData).Name)
		}
		if response.Data.(models.UserData).Email != "jane.doe@example.com" {
			t.Fatalf("expected email to be jane.doe@example.com, got %v", response.Data.(models.UserData).Email)
		}
	})

	t.Run("Fail_RegisterDoctorTransaction", func(t *testing.T) {
		// Configurar el mock para retornar un error
		mockRepo.RegisterDoctorFunc = func(ctx context.Context, rd entities.RegisterDoctor) (entities.UserResponse, error) {
			return entities.UserResponse{}, errors.New("insert doctor error")
		}

		// Preparar los valores de prueba
		request := models.RegisterDoctorRequest{
			Name:           "Jane",
			Lastname1:      "Doe",
			Lastname2:      "Smith",
			Email:          "jane.doe@example.com",
			Password:       "password",
			Rol:            2,
			Specialty:      1,
			MedicalLicense: "LIC1234567", // 10 caracteres
		}

		// Llamar a la función que se está probando
		response, err := usecase.RegisterDoctor(context.Background(), request)

		// Validar que hubo un error
		if err == nil {
			t.Fatalf("expected error, got none")
		}
		if response.Status != "error" {
			t.Fatalf("expected error status, got %v", response.Status)
		}
		if _, ok := response.Errors.(map[string]string)["register"]; !ok {
			t.Fatalf("expected register error, got none")
		}
	})
}
