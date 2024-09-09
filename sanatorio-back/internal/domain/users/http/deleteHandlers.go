package v1

import (
	"encoding/json"
	"net/http"
)

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var request struct {
		AccountID string `json:"account_id"`
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Llamar al caso de uso para eliminar el usuario
	message, err := h.uc.DeleteUser(r.Context(), request.AccountID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "Failed to delete user",
			"errors":  err.Error(),
		})
		return
	}

	// Responder exitosamente con el mensaje retornado
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"message": message,
	})
}

func (h *handler) DeleteDoctor(w http.ResponseWriter, r *http.Request) {
	var request struct {
		AccountID string `json:"account_id"`
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Llamar al caso de uso para eliminar al doctor
	message, err := h.uc.DeleteDoctor(r.Context(), request.AccountID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "Failed to delete doctor",
			"errors":  err.Error(),
		})
		return
	}

	// Responder exitosamente con el mensaje retornado
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"message": message,
	})
}

func (h *handler) SoftDeleteUser(w http.ResponseWriter, r *http.Request) {
	var request struct {
		AccountID string `json:"account_id"`
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Llamar al caso de uso para soft delete del usuario
	message, err := h.uc.SoftDeleteUser(r.Context(), request.AccountID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "Failed to soft delete user",
			"errors":  err.Error(),
		})
		return
	}

	// Responder exitosamente con el mensaje retornado
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"message": message,
	})
}

func (h *handler) SoftDeleteDoctor(w http.ResponseWriter, r *http.Request) {
	var request struct {
		AccountID string `json:"account_id"`
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Llamar al caso de uso para soft delete del doctor
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

	// Responder exitosamente con el mensaje retornado
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"message": message,
	})
}
