package v1

import (
	"encoding/json"
	"log"
	"net/http"
	"sanatorioApp/internal/domain/users/http/models"
)

func (h *handler) RegisterSuperUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Inicio del handler RegisterUser")

	// Decodificar el payload de la solicitud
	request := models.RegisterUserByAdminRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	log.Printf("Recibido request: %+v", request)

	// Llamar al caso de uso para manejar el registro
	userData, err := h.uc.RegisterSuperUser(r.Context(), request)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	// Preparar la respuesta formateada
	response := models.Response{
		Status:  "success",
		Message: "User registered successfully",
		Data:    userData,
	}

	// Enviar la respuesta
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *handler) RegisterDoctor(w http.ResponseWriter, r *http.Request) {
	// Decodificar el payload de la solicitud
	var request models.RegisterDoctorByAdminRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	log.Printf("Recibido request: %+v", request)

	// Llamar al caso de uso para manejar el registro
	doctorData, err := h.uc.RegisterDoctor(r.Context(), request)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	// Preparar la respuesta formateada
	response := models.Response{
		Status:  "success",
		Message: "Doctor registered successfully",
		Data:    doctorData,
	}

	// Enviar la respuesta
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

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
