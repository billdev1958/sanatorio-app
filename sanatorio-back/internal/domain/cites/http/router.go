package v1

import (
	"net/http"
	"sanatorioApp/internal/auth"
)

func (h *handler) CitesRoutes(mux *http.ServeMux) {
	mux.Handle("POST /v1/office", CORS(auth.AuthMiddleware(http.HandlerFunc(h.RegisterOffice))))

	mux.Handle("POST /v1/specialty", CORS(auth.AuthMiddleware(http.HandlerFunc(h.RegisterSpecialty))))

	mux.Handle("POST /v1/schedule", CORS(auth.AuthMiddleware(http.HandlerFunc(h.RegisterOfficeSchedule))))

	mux.Handle("GET /v1/schedule", CORS(auth.AuthMiddleware(http.HandlerFunc(h.GetAllOfficeSchedules))))

}
