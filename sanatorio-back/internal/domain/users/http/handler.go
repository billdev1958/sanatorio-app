package v1

import (
	"encoding/json"
	"net/http"
	"sanatorioApp/internal/domain/auth"
	user "sanatorioApp/internal/domain/users"
	"sanatorioApp/internal/domain/users/http/models"
)

type handler struct {
	uc         user.Usecase
	authUc     auth.AuthUsecases
	middleware *auth.Middleware
}

func NewHandler(uc user.Usecase, authUc auth.AuthUsecases) *handler {
	return &handler{
		uc:         uc,
		authUc:     authUc,
		middleware: auth.NewMiddleware(authUc),
	}
}

func (h *handler) LoginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req models.LoginUser
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Status:  "error",
			Message: "Invalid request payload",
			Errors:  err.Error(),
		})
		return
	}

	res, err := h.uc.LoginUser(r.Context(), req)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(res)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
