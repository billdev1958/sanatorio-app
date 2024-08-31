package v1_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	v1 "sanatorioApp/internal/domain/users/http"
	"sanatorioApp/internal/domain/users/http/models"
	"strings"
	"testing"
)

// MockUsecase es un mock del caso de uso que se utilizará en las pruebas
type MockUsecase struct {
	RegisterUserFunc   func(ctx context.Context, request models.RegisterUserRequest) (models.Response, error)
	RegisterDoctorFunc func(ctx context.Context, request models.RegisterDoctorRequest) (models.Response, error)
}

func (m *MockUsecase) RegisterUser(ctx context.Context, request models.RegisterUserRequest) (models.Response, error) {
	if m.RegisterUserFunc != nil {
		return m.RegisterUserFunc(ctx, request)
	}
	return models.Response{}, nil
}

func (m *MockUsecase) RegisterDoctor(ctx context.Context, request models.RegisterDoctorRequest) (models.Response, error) {
	if m.RegisterDoctorFunc != nil {
		return m.RegisterDoctorFunc(ctx, request)
	}
	return models.Response{}, nil
}

func TestRegisterUser(t *testing.T) {
	mockUC := &MockUsecase{}
	handler := v1.NewHandler(mockUC)

	t.Run("Success", func(t *testing.T) {
		// Configurar el mock para retornar un resultado exitoso
		mockUC.RegisterUserFunc = func(ctx context.Context, request models.RegisterUserRequest) (models.Response, error) {
			return models.Response{
				Status:  "success",
				Message: "User registered successfully",
				Data: map[string]interface{}{
					"name":  request.Name,
					"email": request.Email,
				},
			}, nil
		}

		// Preparar el request body
		reqBody := `{"name":"John","lastname1":"Doe","lastname2":"Smith","email":"john.doe@example.com","password":"password","rol":1,"curp":"CURP12345678901234"}`
		req := httptest.NewRequest(http.MethodPost, "/v1/users", strings.NewReader(reqBody))
		rec := httptest.NewRecorder()

		// Llamar al handler
		handler.RegisterUser(rec, req)

		// Validar la respuesta
		if rec.Code != http.StatusOK {
			t.Fatalf("expected status 200, got %v", rec.Code)
		}

		var response models.Response
		if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
			t.Fatalf("failed to decode response: %v", err)
		}

		// Decodificar manualmente response.Data a un map
		dataMap := response.Data.(map[string]interface{})

		// Convertir el map a UserData
		userData := models.UserData{
			Name:  dataMap["name"].(string),
			Email: dataMap["email"].(string),
		}

		if userData.Name != "John" {
			t.Fatalf("expected name to be John, got %v", userData.Name)
		}
		if userData.Email != "john.doe@example.com" {
			t.Fatalf("expected email to be john.doe@example.com, got %v", userData.Email)
		}
	})

	t.Run("InvalidRequest", func(t *testing.T) {
		// Preparar un request body inválido
		reqBody := `{"name":"John"`
		req := httptest.NewRequest(http.MethodPost, "/v1/users", strings.NewReader(reqBody))
		rec := httptest.NewRecorder()

		// Llamar al handler
		handler.RegisterUser(rec, req)

		// Validar la respuesta
		if rec.Code != http.StatusBadRequest {
			t.Fatalf("expected status 400, got %v", rec.Code)
		}

		expected := "Invalid request payload\n"
		if rec.Body.String() != expected {
			t.Fatalf("expected body %v, got %v", expected, rec.Body.String())
		}
	})

	t.Run("UsecaseError", func(t *testing.T) {
		// Configurar el mock para retornar un error
		mockUC.RegisterUserFunc = func(ctx context.Context, request models.RegisterUserRequest) (models.Response, error) {
			return models.Response{
				Status:  "error",
				Message: "Failed to register user",
			}, errors.New("insert user error")
		}

		// Preparar el request body
		reqBody := `{"name":"John","lastname1":"Doe","lastname2":"Smith","email":"john.doe@example.com","password":"password","rol":1,"curp":"CURP12345678901234"}`
		req := httptest.NewRequest(http.MethodPost, "/v1/users", strings.NewReader(reqBody))
		rec := httptest.NewRecorder()

		// Llamar al handler
		handler.RegisterUser(rec, req)

		// Validar la respuesta
		if rec.Code != http.StatusInternalServerError {
			t.Fatalf("expected status 500, got %v", rec.Code)
		}

		var response models.Response
		if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
			t.Fatalf("failed to decode response: %v", err)
		}

		if response.Status != "error" {
			t.Fatalf("expected error status, got %v", response.Status)
		}
	})
}
