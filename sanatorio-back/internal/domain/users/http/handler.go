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

func (h *handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	// Decodificar el cuerpo de la solicitud en el modelo RegisterUserRequest
	request := models.RegisterUserRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		// Respuesta de error si el payload no es válido
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Llamar al caso de uso para registrar el usuario
	response, err := h.uc.RegisterUser(r.Context(), request)
	if err != nil {
		// Manejar el error y enviar una respuesta adecuada usando response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Configurar el tipo de contenido de la respuesta y el estado HTTP
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Enviar la respuesta exitosa usando response
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
