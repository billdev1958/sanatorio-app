package v1

import (
	"encoding/json"
	"net/http"
	schedule "sanatorioApp/internal/domain/schedules"
	"sanatorioApp/internal/domain/users/http/models"
)

type ScheduleHandler struct {
	Usecase schedule.OfficeSchedule
}

func NewScheduleHandler(usecase schedule.OfficeSchedule) *ScheduleHandler {
	return &ScheduleHandler{
		Usecase: usecase,
	}
}

func (h *ScheduleHandler) GetInfoOfficeSchedule(w http.ResponseWriter, r *http.Request) {
	// Crear contexto desde la solicitud
	ctx := r.Context()

	// Llamar al caso de uso
	responseData, err := h.Usecase.GetInfoOfficeSchedule(ctx)
	if err != nil {
		// Manejar errores directamente
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(models.Response{
			Status:  "error",
			Message: "Error al obtener la información del horario",
			Errors:  err.Error(),
		})
		return
	}

	// Enviar respuesta exitosa directamente
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(models.Response{
		Status:  "success",
		Message: "Información obtenida correctamente",
		Data:    responseData,
	})
}
