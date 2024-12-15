package v1

import (
	"net/http"
	"sanatorioApp/internal/auth"
)

func (h *ScheduleHandler) ScheduleRouter(mux *http.ServeMux) {
	mux.Handle("POST /v1/admin/schedule", (auth.AuthMiddleware(http.HandlerFunc(h.GetInfoOfficeSchedule))))

}
