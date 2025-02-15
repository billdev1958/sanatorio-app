package v1

import (
	"net/http"
	"sanatorioApp/internal/domain/auth"
	user "sanatorioApp/internal/domain/users"
)

func (h *handler) AppointmentRouter(mux *http.ServeMux) {

	mux.Handle("POST /v1/appointment/schedules",
		auth.AuthMiddleware(
			h.middleware.RequiredPermission(user.ViewAppointment)(
				http.HandlerFunc(h.GetSchedulesForAppointment))))

	mux.Handle("GET /v1/appointment/schedules",
		auth.AuthMiddleware(
			h.middleware.RequiredPermission(user.ViewAppointment)(
				http.HandlerFunc(h.GetParamsForAppointments))))

	mux.Handle("POST /v1/appointment",
		auth.AuthMiddleware(
			h.middleware.RequiredPermission(user.CreateAppointment)(
				http.HandlerFunc(h.RegisterAppointment))))

	mux.Handle("GET /v1/appointment/{appointmentID}",
		auth.AuthMiddleware(
			h.middleware.RequiredPermission(user.ViewAppointment)(
				http.HandlerFunc(h.GetAppointmentByID))))

	mux.Handle("POST /v1/appointments/patient",
		auth.AuthMiddleware(
			h.middleware.RequiredPermission(user.ViewAppointment)(
				http.HandlerFunc(h.GetAppointmentsForPatient))))
}
