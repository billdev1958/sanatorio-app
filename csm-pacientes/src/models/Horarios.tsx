export interface Horarios{
   scheduleID: number
   officeID: number
   dayOfWeek: number
   timeStart: string
   timeEnd: string 
}

export interface services{
    serviceID: number
    name: string
}

export interface ScheduleAppointment {
    shift: number; // 1 para Matutino, 2 para Vespertino
    service: number; // ID del servicio seleccionado (Ej: 25 para Medicina General)
    day: number; // Día de la semana (0=Domingo, 1=Lunes, ..., 6=Sábado)
  }

  export interface ApiResponse<T> {
    status: string; // Ejemplo: "success", "error"
    message: string; // Mensaje descriptivo de la respuesta
    data: T; // Datos específicos del servicio (tipado genérico)
  }
  
  export interface ScheduleData {
    ID: number; // Identificador único del horario
    OfficeID: number; // ID de la oficina
    ShiftID: number; // Turno (1: Matutino, 2: Vespertino)
    ServiceID: number; // ID del servicio
    DoctorID: string; // ID del doctor asociado
    StatusID: number; // Estado del horario (ej. 1: Activo)
    DayOfWeek: number; // Día de la semana (0=Domingo, 1=Lunes, ..., 6=Sábado)
    time_start: string; // Hora de inicio (formato HH:mm)
    time_end: string; // Hora de finalización (formato HH:mm)
    time_duration: string; // Duración del horario (ej. "1h0m0s")
  }
  