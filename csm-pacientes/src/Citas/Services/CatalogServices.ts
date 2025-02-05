import api from '../../Api/Api';
import { Response } from '../../Api/Model';
import { Appointment, Services, Shift } from '../Models/Catalogs';
import { SchedulesAppointmentRequest, OfficeScheduleResponse , PatientAndBeneficiaries, RegisterAppointmentRequest} from '../Models/Catalogs';

export const getParamsForAppointment = async (
  token?: string
): Promise<Response<{ patients: PatientAndBeneficiaries; services: Services[]; shifts: Shift[] }>> => {
  console.log('getParamsForAppointment - Starting request with token:', token ? 'Present' : 'Missing');
  try {
    const response = await api.get('/appointment/schedules', {
      headers: {
        Authorization: token ? `Bearer ${token}` : '',
      },
    });

    console.log('getParamsForAppointment - Success response:', response.data);

    return response.data as Response<{ patients: PatientAndBeneficiaries; services: Services[]; shifts: Shift[] }>;
  } catch (error: any) {
    console.error('getParamsForAppointment - Error:', error);
    console.error('Error details:', {
      message: error.message,
      response: error.response?.data,
      status: error.response?.status,
    });
    throw new Error(error.message || 'Error al obtener los parámetros');
  }
};

export const getOfficeSchedules = async (
  appointmentData: SchedulesAppointmentRequest,
  token?: string
): Promise<Response<{ data: OfficeScheduleResponse[] }>> => {
  console.log("getOfficeSchedules - Request data:", appointmentData);

  // Validar datos de la solicitud
  if (!appointmentData.service || !appointmentData.shift || !appointmentData.appointmentDate) {
    console.error("getOfficeSchedules - Datos de cita incompletos:", appointmentData);
    throw new Error("Datos de cita incompletos. Por favor, complete todos los campos.");
  }

  // Validar formato de la fecha (YYYY-MM-DD) y UTC (termina en 'Z')
  const utcDateRegex = /^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\.\d{3}Z$/;
  if (!utcDateRegex.test(appointmentData.appointmentDate)) {
    console.error("getOfficeSchedules - Formato de fecha inválido o no está en UTC:", appointmentData.appointmentDate);
    throw new Error("Formato de fecha inválido. La fecha debe estar en formato UTC (ISO 8601).");
  }

  try {
    console.log("getOfficeSchedules - Realizando solicitud a la API con:", {
      appointmentData,
      token: token ? "Present" : "Missing",
    });

    const response = await api.post("/appointment/schedules", appointmentData, {
      headers: {
        Authorization: token ? `Bearer ${token}` : "",
      },
    });

    // Verificar y retornar la respuesta
    console.log("getOfficeSchedules - Respuesta exitosa:", response.data);
    return response.data as Response<{ data: OfficeScheduleResponse[] }>;
  } catch (error: any) {
    console.error("getOfficeSchedules - Error al realizar la solicitud:", error);

    // Manejar errores HTTP
    const errorMessage =
      error.response?.data?.message || "Error al obtener los horarios disponibles.";
    console.error("Detalles del error:", {
      message: error.message,
      response: error.response?.data,
      status: error.response?.status,
    });
    throw new Error(errorMessage);
  }
};

/**
 * Registrar una nueva cita en el sistema.
 * @param appointmentData - Datos de la cita que se va a registrar.
 * @param token - Token de autorización (opcional).
 * @returns Respuesta de la API con el resultado de la operación.
 */
export const registerAppointment = async (
  appointmentData: RegisterAppointmentRequest,
  token?: string
): Promise<Response<{ success: boolean; appointmentID: string }>> => {
  console.log('registerAppointment - Datos de la cita a registrar:', appointmentData);

  if (
    !appointmentData.scheduleID ||
    !appointmentData.patientID ||
    !appointmentData.timeStart ||
    !appointmentData.timeEnd
  ) {
    console.error('registerAppointment - Datos de cita incompletos:', appointmentData);
    throw new Error('Datos de cita incompletos. Por favor, complete todos los campos obligatorios.');
  }

  const isoDateTimeRegex = /^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(?:\.\d+)?Z?$/;
  if (
    !isoDateTimeRegex.test(appointmentData.timeStart) ||
    !isoDateTimeRegex.test(appointmentData.timeEnd)
  ) {
    console.error('registerAppointment - Formato de fecha/hora inválido:', {
      timeStart: appointmentData.timeStart,
      timeEnd: appointmentData.timeEnd,
    });
    throw new Error('Formato de fecha y hora inválido. Use el formato ISO 8601.');
  }

  try {
    console.log('registerAppointment - Realizando solicitud a la API con:', {
      appointmentData,
      token: token ? 'Present' : 'Missing',
    });

    const response = await api.post('/appointment', appointmentData, {
      headers: {
        Authorization: token ? `Bearer ${token}` : '',
      },
    });

    console.log('registerAppointment - Respuesta del servidor:', response.data);

    const data = response.data;

    if (typeof data === 'string') {
      console.log("registerAppointment - La respuesta es un string:", data);
      return {
        status: 'success',
        message: data,
        data: { success: true, appointmentID: '' }
      } as Response<{ success: boolean; appointmentID: string }>;
    } else {
      console.log("registerAppointment - La respuesta es un objeto:", data);
      const res = data as Response<{ success: boolean; appointmentID: string }>;
      if (res.status !== 'success' || !res.data?.success) {
        throw new Error(res.message || 'Error al registrar la cita.');
      }
      return res;
    }
  } catch (error: any) {
    console.error('registerAppointment - Error al realizar la solicitud:', error);
    console.error('Detalles del error:', {
      message: error.message,
      response: error.response?.data,
      status: error.response?.status,
    });
    const errorMessage =
      error.response?.data?.message || error.message || 'Error al registrar la cita.';
    throw new Error(errorMessage);
  }
};

/**
 * Obtiene el historial de citas médicas del paciente.
 * @param token Token de autorización (opcional). Si se proporciona, se envía en el header.
 * @returns Una promesa que resuelve en un arreglo de citas.
 */
export const getPatientAppointments = async (token?: string): Promise<Appointment[]> => {
  console.log(
    'getPatientAppointments - Iniciando solicitud con token:',
    token ? 'Presente' : 'Missing'
  );

  try {
    // Se envía un POST a la url '/appointments/patient'. 
    // Como no se requiere enviar datos en el body, se pasa un objeto vacío {}.
    const response = await api.post('/appointments/patient', {}, {
      headers: {
        Authorization: token ? `Bearer ${token}` : "",
      },
    });

    console.log('getPatientAppointments - Respuesta exitosa:', response.data);

    // Se asume que la respuesta tiene la estructura:
    // { status: "success", message: "Citas obtenidas exitosamente.", data: Appointment[] }
    const res = response.data as Response<Appointment[]>;

    if (res.status !== 'success') {
      throw new Error(res.message || 'Error al obtener citas');
    }

    // Retorna el arreglo de citas (o un arreglo vacío si no hay data)
    return res.data || [];
  } catch (error: any) {
    console.error('getPatientAppointments - Error al obtener las citas:', error);
    console.error('Detalles del error:', {
      message: error.message,
      response: error.response?.data,
      status: error.response?.status,
    });
    throw new Error(error.message || 'Error al obtener el historial de citas.');
  }
};