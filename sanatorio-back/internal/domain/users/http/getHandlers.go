package v1

import (
	"encoding/json"
	"net/http"
	"sanatorioApp/internal/domain/users/http/models"
	"strconv"
)

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
	response, err := h.uc.GetDoctorByID(r.Context(), userID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.Response{
			Status:  "error",
			Message: "Failed to retrieve doctor",
			Errors:  err.Error(),
		})
		return
	}

	// Responde con los datos del doctor
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.uc.GetUsers(r.Context())
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.Response{
			Status:  "error",
			Message: "Failed to retrieve users",
			Errors:  err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.Response{
		Status:  "success",
		Message: "Users retrieved successfully",
		Data:    users,
	})
}

func (h *handler) GetDoctors(w http.ResponseWriter, r *http.Request) {
	doctors, err := h.uc.GetDoctors(r.Context())
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.Response{
			Status:  "error",
			Message: "Failed to retrieve doctors",
			Errors:  err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.Response{
		Status:  "success",
		Message: "Doctors retrieved successfully",
		Data:    doctors,
	})
}

func (h *handler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.PathValue("userId")
	if userIDStr == "" {
		http.Error(w, "userID es obligatorio", http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "userID debe ser un número entero", http.StatusBadRequest)
		return
	}

	response, err := h.uc.GetUserByID(r.Context(), userID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.Response{
			Status:  "error",
			Message: "Failed to retrieve user",
			Errors:  err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *handler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	// Llama al caso de uso para obtener la lista combinada de todos los usuarios
	users, err := h.uc.GetAllUsers(r.Context())
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.Response{
			Status:  "error",
			Message: "Failed to retrieve users",
			Errors:  err.Error(),
		})
		return
	}

	// Responde con la lista de todos los usuarios
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.Response{
		Status:  "success",
		Message: "Users, Doctors, and SuperUsers retrieved successfully",
		Data:    users,
	})
}
