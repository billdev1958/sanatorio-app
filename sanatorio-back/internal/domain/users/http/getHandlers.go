package v1

import (
	"encoding/json"
	"net/http"
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
	doctorData, err := h.uc.GetDoctorByID(r.Context(), userID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "Failed to retrieve doctor",
			"errors":  err.Error(),
		})
		return
	}

	// Responde con los datos del doctor
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "success",
		"message": "Doctor retrieved successfully",
		"data":    doctorData,
	})
}

func (h *handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	// Llama al caso de uso para obtener los super administradores
	usersData, err := h.uc.GetSuperAdmins(r.Context())
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "Failed to retrieve users",
			"errors":  err.Error(),
		})
		return
	}

	// Responde con los datos de los usuarios
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "success",
		"message": "Users retrieved successfully",
		"data":    usersData,
	})
}

func (h *handler) GetDoctors(w http.ResponseWriter, r *http.Request) {
	// Llama al caso de uso para obtener los doctores
	doctorsData, err := h.uc.GetDoctors(r.Context())
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "Failed to retrieve doctors",
			"errors":  err.Error(),
		})
		return
	}

	// Responde con los datos de los doctores
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "success",
		"message": "Doctors retrieved successfully",
		"data":    doctorsData,
	})
}

func (h *handler) GetUserByID(w http.ResponseWriter, r *http.Request) {
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

	// Llama al caso de uso para obtener el usuario por ID
	userData, err := h.uc.GetSuperAdminByID(r.Context(), userID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "Failed to retrieve user",
			"errors":  err.Error(),
		})
		return
	}

	// Responde con los datos del usuario
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "success",
		"message": "User retrieved successfully",
		"data":    userData,
	})
}
