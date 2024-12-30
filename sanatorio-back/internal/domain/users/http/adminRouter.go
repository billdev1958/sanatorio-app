package v1

import (
	"net/http"
	"sanatorioApp/internal/domain/auth"
	user "sanatorioApp/internal/domain/users"
)

func (h *handler) AdminRoutes(mux *http.ServeMux) {
	// Registros
	mux.Handle("POST /v1/admin/superadmin",
		auth.AuthMiddleware(
			h.middleware.RequiredPermission(user.CreateSuperAdmin)(
				http.HandlerFunc(h.RegisterSuperAdmin))))

	mux.Handle("POST /v1/admin/receptionist",
		auth.AuthMiddleware(
			h.middleware.RequiredPermission(user.CreateReceptionist)(
				http.HandlerFunc(h.RegisterReceptionist))))

	mux.Handle("POST /v1/admin/doctor",
		auth.AuthMiddleware(
			h.middleware.RequiredPermission(user.CreateDoctor)(
				http.HandlerFunc(h.RegisterDoctor))))

	mux.Handle("PUT /v1/admin/doctor",
		auth.AuthMiddleware(
			h.middleware.RequiredPermission(user.EditDoctor)(
				http.HandlerFunc(h.UpdateDoctor))))

	mux.Handle("POST /v1/admin/medicalh",
		auth.AuthMiddleware(
			h.middleware.RequiredPermission(user.ViewMedicalHistory)(
				http.HandlerFunc(h.GetMedicalHistoryByID))))

	mux.Handle("POST /v1/admin/medicalhc",
		auth.AuthMiddleware(
			h.middleware.RequiredPermission(user.EditMedicalHistory)(
				http.HandlerFunc(h.CompleteMedicalHistory))))

	// Updates
	mux.Handle("PUT /v1/admin/supera",
		auth.AuthMiddleware(
			h.middleware.RequiredPermission(user.EditSuperAdmin)(
				http.HandlerFunc(h.UpdateSuperAdmin))))

	mux.Handle("PUT /v1/admin/admin",
		auth.AuthMiddleware(
			h.middleware.RequiredPermission(user.EditAdmin)(
				http.HandlerFunc(h.UpdateAdmin))))

	mux.Handle("PUT /v1/admin/receptionist",
		auth.AuthMiddleware(
			h.middleware.RequiredPermission(user.EditReceptionist)(
				http.HandlerFunc(h.UpdateReceptionist))))

	mux.Handle("PUT /v1/admin/patient",
		auth.AuthMiddleware(
			h.middleware.RequiredPermission(user.EditPatient)(
				http.HandlerFunc(h.UpdatePatient))))

	// Get
	mux.Handle("GET /v1/doctor/{userId}", auth.AuthMiddleware(http.HandlerFunc(h.GetDoctorByID)))

	// Delete

	// SoftDelete
	mux.Handle("DELETE /v1/admin/supera",
		auth.AuthMiddleware(
			h.middleware.RequiredPermission(user.DeleteSuperAdmin)(
				http.HandlerFunc(h.SoftDeleteSuperAdmin))))

	mux.Handle("DELETE /v1/admin/admin",
		auth.AuthMiddleware(
			h.middleware.RequiredPermission(user.DeleteAdmin)(
				http.HandlerFunc(h.SoftDeleteAdmin))))

	mux.Handle("DELETE /v1/admin/doctor",
		auth.AuthMiddleware(
			h.middleware.RequiredPermission(user.DeleteDoctor)(
				http.HandlerFunc(h.SoftDeleteDoctor))))

	mux.Handle("DELETE /v1/admin/receptionist",
		auth.AuthMiddleware(
			h.middleware.RequiredPermission(user.DeleteReceptionist)(
				http.HandlerFunc(h.SoftDeleteReceptionist))))

	mux.Handle("DELETE /v1/admin/patient",
		auth.AuthMiddleware(
			h.middleware.RequiredPermission(user.DeletePatient)(
				http.HandlerFunc(h.SoftDeletePatient))))
}
