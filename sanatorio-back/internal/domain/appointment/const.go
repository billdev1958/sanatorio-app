package appointment

type appointmentStatus int

const (
	_ appointmentStatus = iota
	AppointmentStatusPendiente
	AppointmentStatusConfirmada
	AppointmentStatusCancelada
)
