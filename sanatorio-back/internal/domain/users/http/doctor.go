package v1

import (
	"encoding/json"
	"log"
	"net/http"
	"sanatorioApp/internal/domain/users/http/models"
	"strconv"

	"github.com/google/uuid"
)

func (h *handler) RegisterDoctor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Decodificar el payload de la solicitud
	var request models.RegisterDoctorRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	log.Printf("Decoded request: %+v", request)

	// Llamar al caso de uso para manejar el registro
	patientData, err := h.uc.RegisterDoctor(r.Context(), request)
	if err != nil {
		response := models.Response{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}
	// Preparar la respuesta formateada
	response := models.Response{
		Status:  "success",
		Message: "Doctor registered successfully",
		Data:    patientData,
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *handler) UpdateDoctor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var request models.DoctorUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Printf("Failed to decode request: %v", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if request.AccountID == uuid.Nil {
		log.Printf("Missing or invalid account_id in request payload")
		http.Error(w, "Valid account_id is required", http.StatusBadRequest)
		return
	}

	log.Printf("Decoded request: %+v", request)

	message, err := h.uc.UpdatedDoctor(r.Context(), request)
	if err != nil {
		log.Printf("Failed to update doctor with account_id: %s. Error: %v", request.AccountID, err)
		response := models.Response{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := models.Response{
		Status:  "success",
		Message: message,
		Data:    nil,
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Failed to encode response: %v", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	log.Printf("Successfully updated doctor with account_id: %s", request.AccountID)
}

func (h *handler) GetDoctorByID(w http.ResponseWriter, r *http.Request) {
	// Obtén el valor del userID desde los parámetros de la ruta
	userIDStr := r.PathValue("userId")
	if userIDStr == "" {
		http.Error(w, "userID es obligatorio", http.StatusBadRequest)
		return
	}

	// Convierte el userID de string a int
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "userID debe ser un número entero", http.StatusBadRequest)
		return
	}

	// Llama al caso de uso para obtener el doctor por userID
	doctorData, err := h.uc.GetDoctorByID(r.Context(), userID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "Failed to retrieve doctor",
			"errors":  err.Error(),
		})
		return
	}

	// Responde con los datos del doctor
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "success",
		"message": "Doctor retrieved successfully",
		"data":    doctorData,
	})
}

func (h *handler) SoftDeleteDoctor(w http.ResponseWriter, r *http.Request) {
	var request struct {
		AccountID uuid.UUID `json:"account_id"`
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	message, err := h.uc.SoftDeleteDoctor(r.Context(), request.AccountID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "Failed to soft delete doctor",
			"errors":  err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"message": message,
	})
}
