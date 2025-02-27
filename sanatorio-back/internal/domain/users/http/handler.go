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

func (h *handler) AccountVerification(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	type tokenRequest struct {
		Token string `json:"token"`
	}

	var req tokenRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Status:  "error",
			Message: "Solicitud inválida, formato incorrecto",
			Errors:  err.Error(),
		})
		return
	}

	res, err := h.uc.AccountVerification(r.Context(), req.Token)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.Response{
			Status:  "error",
			Message: "Error interno al verificar la cuenta",
			Errors:  err.Error(),
		})
		return
	}

	if res.Status == "error" {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	json.NewEncoder(w).Encode(res)
}
