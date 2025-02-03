package appointment

import (
	"context"
	"sanatorioApp/internal/domain/appointment/entities"
)

type AppointmentRepository interface {
	GetAvaliableSchedules(ctx context.Context, date string, dayOfWeek int, serviceID int, shiftID int) ([]entities.OfficeSchedule, error)

	RegisterAppointment(ctx context.Context, a entities.Appointment, c entities.Consultation) (bool, error)

	//GetAppointmentForPatient(ctx context.Context) ([])
}
