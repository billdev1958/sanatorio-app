package v1

import (
	"encoding/json"
	"log"
	"net/http"
	"sanatorioApp/internal/domain/users/http/models"

	"github.com/google/uuid"
)

func (h *handler) RegisterPatient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Decodificar el payload de la solicitud
	var request models.RegisterPatientRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// log.Printf("Decoded request: %+v", request)

	// Llamar al caso de uso para manejar el registro
	patientData, err := h.uc.RegisterPatient(r.Context(), request)
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
		Message: "Patient registered successfully",
		Data:    patientData,
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *handler) UpdatePatient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var request models.UpdateUser
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

	message, err := h.uc.UpdatedPatient(r.Context(), request)
	if err != nil {
		log.Printf("Failed to update patient with account_id: %s. Error: %v", request.AccountID, err)
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

	log.Printf("Successfully updated patient with account_id: %s", request.AccountID)
}

func (h *handler) RegisterBeneficiary(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Decodificar el payload de la solicitud
	var request models.RegisterBeneficiaryRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	log.Printf("Decoded request: %+v", request)

	// Llamar al caso de uso para manejar el registro
	patientData, err := h.uc.RegisterBeneficiary(r.Context(), request)
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
		Message: "Patient registered successfully",
		Data:    patientData,
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *handler) SoftDeletePatient(w http.ResponseWriter, r *http.Request) {
	var request struct {
		AccountID uuid.UUID `json:"account_id"`
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	message, err := h.uc.SoftDeletePatient(r.Context(), request.AccountID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "Failed to soft delete patient",
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
