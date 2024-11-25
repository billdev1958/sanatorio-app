package v1

import (
	"encoding/json"
	"log"
	"net/http"
	"sanatorioApp/internal/domain/users/http/models"
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
