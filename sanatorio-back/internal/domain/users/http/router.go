package v1

import (
	"net/http"
	"sanatorioApp/internal/domain/auth"
	user "sanatorioApp/internal/domain/users"
)

func (h *handler) UserRoutes(mux *http.ServeMux) {
	// Public routes
	mux.Handle("POST /v1/login", http.HandlerFunc(h.LoginUser))

	mux.Handle("POST /v1/patients", http.HandlerFunc(h.RegisterPatient))

	mux.Handle("PUT /v1/patient",
		auth.AuthMiddleware(
			h.middleware.RequiredPermission(user.EditPatient)(
				http.HandlerFunc(h.UpdatePatient))))

	mux.Handle("POST /v1/beneficiary",
		auth.AuthMiddleware(
			h.middleware.RequiredPermission(user.CreateBeneficiary)(
				http.HandlerFunc(h.RegisterBeneficiary))))
}
