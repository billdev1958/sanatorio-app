package v1

import (
	"encoding/json"
	"net/http"
	user "sanatorioApp/internal/domain/users"
	"sanatorioApp/internal/domain/users/http/models"
)

type handler struct {
	uc user.Usecase
}

func NewHandler(uc user.Usecase) *handler {
	return &handler{uc: uc}
}

func (h *handler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var request models.LoginUser

	// Establecer Content-Type para la respuesta
	w.Header().Set("Content-Type", "application/json")

	// Decodificar la solicitud del cuerpo
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Llamar al caso de uso para manejar el login
	response, err := h.uc.LoginUser(r.Context(), request)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Authentication failed"})
		return
	}

	// Si el login fue exitoso, devolver el token y la información del usuario
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
