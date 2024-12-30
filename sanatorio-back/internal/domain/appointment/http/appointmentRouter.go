package v1

import (
	"net/http"
	"sanatorioApp/internal/domain/auth"
	user "sanatorioApp/internal/domain/users"
)

func (h *handler) AppointmentRouter(mux *http.ServeMux) {
	mux.Handle("PUT /v1/appointment/schedules",
		auth.AuthMiddleware(
			h.middleware.RequiredPermission(user.ViewAppointment)(
				http.HandlerFunc(h.GetSchedulesForAppointment))))
}
