package v1

import (
	"net/http"
	"sanatorioApp/internal/auth"
)

func (h *ScheduleHandler) ScheduleRouter(mux *http.ServeMux) {
	mux.Handle("GET /v1/admin/schedule", CORS(auth.AuthMiddleware(http.HandlerFunc(h.GetInfoOfficeSchedule))))

}
