package v1

import (
	"encoding/json"
	"net/http"
	"sanatorioApp/internal/domain/appointment"
	"sanatorioApp/internal/domain/appointment/http/models"
	"sanatorioApp/internal/domain/auth"
)

type handler struct {
	uc         appointment.Usecase
	authUc     auth.AuthUsecases
	middleware *auth.Middleware
}

func NewHandler(uc appointment.Usecase, authUc auth.AuthUsecases) *handler {
	return &handler{
		uc:         uc,
		authUc:     authUc,
		middleware: auth.NewMiddleware(authUc),
	}
}

func (h *handler) GetSchedulesForAppointment(w http.ResponseWriter, r *http.Request) {
	var filtersRequest models.SchedulesAppointmentRequest
	if err := json.NewDecoder(r.Body).Decode(&filtersRequest); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	schedules, err := h.uc.GetAvaliableSchedules(r.Context(), filtersRequest)
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

func (h *handler) GetParamsForAppointments(w http.ResponseWriter, r *http.Request) {
	claims := auth.ExtractClaims(r.Context())
	if claims == nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	params, err := h.uc.GetParamsForAppointments(r.Context(), claims.AccountID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.Response{
			Status:  "error",
			Message: "Failed to retrieve params for register appointment",
			Errors:  err.Error(),
		})
		return
	}

	response := params

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
