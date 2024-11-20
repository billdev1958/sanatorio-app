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
	// Decodificar los filtros desde el cuerpo de la solicitud
	var filtersRequest models.OfficeSCheduleFiltersRequest
	if err := json.NewDecoder(r.Body).Decode(&filtersRequest); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Llamar al caso de uso con los filtros
	schedules, err := h.uc.GetSchedules(r.Context(), filtersRequest)
	if err != nil {
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

func (h *handler) GetAllOffices(w http.ResponseWriter, r *http.Request) {
	// Llamar al caso de uso con los filtros
	schedules, err := h.uc.GetOffices(r.Context())
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.Response{
			Status:  "error",
			Message: "Failed to retrieve offices",
			Errors:  err.Error(),
		})
		return
	}

	// Construir respuesta de éxito
	response := models.Response{
		Status:  "success",
		Message: "Office offices retrieved successfully",
		Data:    schedules,
	}

	// Responder con éxito
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *handler) UpdateOffice(w http.ResponseWriter, r *http.Request) {
	var request models.UpdateOfficeRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if request.OfficeID == 0 || request.OfficeName == "" {
		http.Error(w, "Missing required fields: office_id or office_name", http.StatusBadRequest)
		return
	}

	office, err := h.uc.UpdateOffice(r.Context(), request)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	response := models.Response{
		Status:  "success",
		Message: "office updated successfully",
		Data:    office,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

}
