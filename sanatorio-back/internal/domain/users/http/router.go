package v1

import (
	"net/http"
	"sanatorioApp/internal/auth"
)

func (h *handler) UserRoutes(mux *http.ServeMux) {

	mux.HandleFunc("POST /v1/login", h.LoginUser)
	// Registros
	mux.HandleFunc("POST /v1/patients", h.RegisterPatient)
	mux.Handle("POST /v1/users", auth.AuthMiddleware(http.HandlerFunc(h.RegisterUser)))
	mux.Handle("POST /v1/doctors", auth.AuthMiddleware(http.HandlerFunc(h.RegisterDoctor)))

	// Updates
	mux.Handle("PUT /v1/users", auth.AuthMiddleware(http.HandlerFunc(h.UpdateUser)))
	mux.Handle("PUT /v1/doctors", auth.AuthMiddleware(http.HandlerFunc(h.UpdateDoctor)))

	// Get
	mux.Handle("GET /v1/users", auth.AuthMiddleware(http.HandlerFunc(h.GetUsers)))
	mux.Handle("GET /v1/doctors", auth.AuthMiddleware(http.HandlerFunc(h.GetDoctors)))
	mux.Handle("GET /v1/user/{accountID}", auth.AuthMiddleware(http.HandlerFunc(h.GetUserByID)))
	mux.Handle("GET /v1/doctor/{accountID}", auth.AuthMiddleware(http.HandlerFunc(h.GetDoctorByID)))

	// Delete
	mux.Handle("DELETE /v1/user", auth.AuthMiddleware(http.HandlerFunc(h.DeleteUser)))
	mux.Handle("DELETE /v1/doctor", auth.AuthMiddleware(http.HandlerFunc(h.DeleteDoctor)))

	// SoftDelete
	mux.Handle("DELETE /v1/users", auth.AuthMiddleware(http.HandlerFunc(h.SoftDeleteUser)))
	mux.Handle("DELETE /v1/doctors", auth.AuthMiddleware(http.HandlerFunc(h.SoftDeleteDoctor)))
}
