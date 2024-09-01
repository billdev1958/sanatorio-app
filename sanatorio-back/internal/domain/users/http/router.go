package v1

import (
	"net/http"
	"sanatorioApp/internal/auth"
)

func (h *handler) UserRoutes(mux *http.ServeMux) {

	mux.HandleFunc("POST /v1/login", h.LoginUser)
	mux.Handle("POST /v1/users", auth.AuthMiddleware(http.HandlerFunc(h.RegisterUser)))

	mux.Handle("PUT /v1/users", auth.AuthMiddleware(http.HandlerFunc(h.UpdateUser)))
	mux.Handle("PUT /v1/doctors", auth.AuthMiddleware(http.HandlerFunc(h.UpdateDoctor)))

	mux.Handle("GET /v1/users", auth.AuthMiddleware(http.HandlerFunc(h.GetUsers)))
	mux.Handle("GET /v1/doctors", auth.AuthMiddleware(http.HandlerFunc(h.GetDoctors)))
	mux.Handle("GET /v1/user/{accountID}", auth.AuthMiddleware(http.HandlerFunc(h.GetUserByID)))
	mux.Handle("GET /v1/doctor/{accountID}", auth.AuthMiddleware(http.HandlerFunc(h.GetDoctorByID)))

	mux.Handle("DELETE /v1/user", auth.AuthMiddleware(http.HandlerFunc(h.DeleteUser)))
	mux.Handle("DELETE /v1/doctor", auth.AuthMiddleware(http.HandlerFunc(h.DeleteDoctor)))

	mux.Handle("DELETE /v1/users", auth.AuthMiddleware(http.HandlerFunc(h.SoftDeleteUser)))
	mux.Handle("DELETE /v1/doctors", auth.AuthMiddleware(http.HandlerFunc(h.SoftDeleteDoctor)))
}
