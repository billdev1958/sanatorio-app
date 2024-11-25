package v1

import (
	"net/http"
	"sanatorioApp/internal/auth"
)

func (h *handler) UserRoutes(mux *http.ServeMux) {
	mux.Handle("POST /v1/login", http.HandlerFunc(h.LoginUser))

	// Registros
	mux.HandleFunc("POST /v1/patients", h.RegisterPatient)

	mux.Handle("POST /v1/superadmin", auth.AuthMiddleware(http.HandlerFunc(h.RegisterSuperAdmin)))
	mux.Handle("POST /v1/receptionist", auth.AuthMiddleware(http.HandlerFunc(h.RegisterReceptionist)))
	mux.Handle("POST /v1/doctor", auth.AuthMiddleware(http.HandlerFunc(h.RegisterDoctor)))

	mux.Handle("POST /v1/medicalh", auth.AuthMiddleware(http.HandlerFunc(h.GetMedicalHistoryByID)))
	mux.Handle("POST /v1/medicalhc", auth.AuthMiddleware(http.HandlerFunc(h.CompleteMedicalHistory)))

	mux.Handle("POST /v1/beneficiary", auth.AuthMiddleware(http.HandlerFunc(h.RegisterBeneficiary)))

	// Updates
	mux.Handle("PUT /v1/users", auth.AuthMiddleware(http.HandlerFunc(h.UpdateUser)))

	// Get
	mux.Handle("GET /v1/doctor/{userId}", auth.AuthMiddleware(http.HandlerFunc(h.GetDoctorByID)))

	// Delete
	mux.Handle("DELETE /v1/user", auth.AuthMiddleware(http.HandlerFunc(h.DeleteUser)))

	// SoftDelete
	mux.Handle("DELETE /v1/users", auth.AuthMiddleware(http.HandlerFunc(h.SoftDeleteUser)))
}
