package v1

import (
	"encoding/json"
	"net/http"
	"sanatorioApp/internal/domain/cites"
	"sanatorioApp/internal/domain/cites/http/models"
)

type handler struct {
	uc cites.Usecase
}

func NewHandler(uc cites.Usecase) *handler {
	return &handler{uc: uc}
}

func (h *handler) RegisterSpecialty(w http.ResponseWriter, r *http.Request) {
	var request models.RegisterSpecialtyRequest

	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	specialty, err := h.uc.RegisterSpecialty(r.Context(), request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	response := models.Response{
		Status:  "success",
		Message: "Specialty registered successfully",
		Data:    specialty,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *handler) RegisterOffice(w http.ResponseWriter, r *http.Request) {
	var request models.RegisterOfficeRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	office, err := h.uc.RegisterOffice(r.Context(), request)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	response := models.Response{
		Status:  "success",
		Message: "office registered successfully",
		Data:    office,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *handler) RegisterOfficeSchedule(w http.ResponseWriter, r *http.Request) {
	var request models.RegisterOfficeScheduleRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	schedule, err := h.uc.RegisterOfficeSchedule(r.Context(), request)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	response := models.Response{
		Status:  "success",
		Message: "office registered successfully",
		Data:    schedule,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

}

func (h *handler) GetAllOfficeSchedules(w http.ResponseWriter, r *http.Request) {
	// Llamar al caso de uso para obtener todos los horarios
	schedules, err := h.uc.GetSchedules(r.Context())
	if err != nil {
		// Manejar errores del caso de uso
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.Response{
			Status:  "error",
			Message: "Failed to retrieve office schedules",
			Errors:  err.Error(),
		})
		return
	}

	// Construir respuesta de éxito
	response := models.Response{
		Status:  "success",
		Message: "Office schedules retrieved successfully",
		Data:    schedules,
	}

	// Responder con éxito
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
