import api from '../../Api/Api';
import { Response } from '../../Api/Model';
import { Services, Shift } from '../Models/Catalogs';
import { SchedulesAppointmentRequest, OfficeScheduleResponse , PatientAndBeneficiaries} from '../Models/Catalogs';

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
    console.log('getOfficeSchedules - Request data:', appointmentData);
  
    // Validar datos de la solicitud
    if (!appointmentData.service || !appointmentData.shift || !appointmentData.appointmentDate) {
      console.error('getOfficeSchedules - Datos de cita incompletos:', appointmentData);
      throw new Error('Datos de cita incompletos. Por favor, complete todos los campos.');
    }
  
    // Validar formato de la fecha
    const dateRegex = /^\d{4}-\d{2}-\d{2}$/;
    if (!dateRegex.test(appointmentData.appointmentDate)) {
      console.error('getOfficeSchedules - Formato de fecha inválido:', appointmentData.appointmentDate);
      throw new Error('Formato de fecha inválido. Use el formato YYYY-MM-DD.');
    }
  
    try {
      console.log('getOfficeSchedules - Realizando solicitud a la API con:', {
        appointmentData,
        token: token ? 'Present' : 'Missing',
      });
  
      const response = await api.post('/appointment/schedules', appointmentData, {
        headers: {
          Authorization: token ? `Bearer ${token}` : '',
        },
      });
  
      // Verificar y retornar la respuesta
      console.log('getOfficeSchedules - Respuesta exitosa:', response.data);
      return response.data as Response<{ data: OfficeScheduleResponse[] }>;
    } catch (error: any) {
      console.error('getOfficeSchedules - Error al realizar la solicitud:', error);
  
      // Manejar errores HTTP
      const errorMessage =
        error.response?.data?.message || 'Error al obtener los horarios disponibles.';
      console.error('Detalles del error:', {
        message: error.message,
        response: error.response?.data,
        status: error.response?.status,
      });
      throw new Error(errorMessage);
    }
  };
