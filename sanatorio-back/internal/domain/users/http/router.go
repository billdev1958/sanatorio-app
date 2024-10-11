package v1

import (
	"net/http"
	"sanatorioApp/internal/auth"
)

func (h *handler) UserRoutes(mux *http.ServeMux) {

	mux.Handle("/v1/login", CORS(http.HandlerFunc(h.LoginUser)))
	// Registros
	mux.HandleFunc("POST /v1/patients", h.RegisterPatient)

	// Updates
	mux.Handle("PUT /v1/users", CORS(auth.AuthMiddleware(http.HandlerFunc(h.UpdateUser))))

	// Get
	mux.Handle("GET /v1/doctor/{userId}", CORS(auth.AuthMiddleware(http.HandlerFunc(h.GetDoctorByID))))

	// Delete
	mux.Handle("DELETE /v1/user", CORS(auth.AuthMiddleware(http.HandlerFunc(h.DeleteUser))))

	// SoftDelete
	mux.Handle("DELETE /v1/users", CORS(auth.AuthMiddleware(http.HandlerFunc(h.SoftDeleteUser))))
}
