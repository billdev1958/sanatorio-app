package v1

import (
	"log"
	"net/http"
	user "sanatorioApp/internal/domain/users"
)

type handler struct {
	uc user.Usecase
}

func NewHandler(uc user.Usecase) *handler {
	return &handler{uc: uc}
}

func (h *handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	// Llama al usecase para obtener el mensaje desde la base de datos
	message, err := h.uc.RegisterUser(r.Context())
	if err != nil {
		log.Printf("error getting message from usecase: %v", err)
		http.Error(w, "failed to retrieve message", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(message))
	if err != nil {
		log.Printf("error writing response: %v", err)
		http.Error(w, "failed to write response", http.StatusInternalServerError)
	}
}
