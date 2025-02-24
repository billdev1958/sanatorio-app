package v1

import (
	"encoding/json"
	"net/http"
	"sanatorioApp/internal/domain/email"
	"sanatorioApp/internal/domain/email/http/models"
)

type handler struct {
	e email.EmailS
}

func NewHandler(e email.EmailS) *handler {
	return &handler{e: e}
}

func (h *handler) ConfirmAccount(w http.ResponseWriter, r *http.Request) {
	var request models.ConfirmRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, `{"error": "Invalid request payload"}`, http.StatusBadRequest)
		return
	}

	confirm, err := h.e.ConfirmAccount(r.Context(), request)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	response := models.Response{
		Status:  "success",
		Message: "Account confirmed successfully",
		Data:    confirm,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, `{"error": "Failed to encode response"}`, http.StatusInternalServerError)
	}
}
