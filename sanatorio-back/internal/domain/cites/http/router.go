package v1

import (
	"net/http"
	"sanatorioApp/internal/domain/auth"
)

func (h *handler) CitesRoutes(mux *http.ServeMux) {
	mux.Handle("POST /v1/office", CORS(auth.AuthMiddleware(http.HandlerFunc(h.RegisterOffice))))

	mux.Handle("PUT /v1/office", CORS(auth.AuthMiddleware(http.HandlerFunc(h.UpdateOffice))))

	mux.Handle("GET /v1/offices", CORS(auth.AuthMiddleware(http.HandlerFunc(h.GetAllOffices))))

	mux.Handle("POST /v1/schedule", CORS(auth.AuthMiddleware(http.HandlerFunc(h.RegisterOfficeSchedule))))

	mux.Handle("POST /v1/schedules", CORS(auth.AuthMiddleware(http.HandlerFunc(h.GetAllOfficeSchedules))))

}
