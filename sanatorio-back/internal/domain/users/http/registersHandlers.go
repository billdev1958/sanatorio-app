package v1

import (
	"encoding/json"
	"net/http"
	"sanatorioApp/internal/domain/users/http/models"
)

func (h *handler) RegisterPatient(w http.ResponseWriter, r *http.Request) {
	// Decodificar el payload de la solicitud
	var request models.RegisterPatientRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Llamar al caso de uso para manejar el registro
	patientData, err := h.uc.RegisterPatient(r.Context(), request)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	// Preparar la respuesta formateada
	response := models.Response{
		Status:  "success",
		Message: "Patient registered successfully",
		Data:    patientData,
	}

	// Enviar la respuesta
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
