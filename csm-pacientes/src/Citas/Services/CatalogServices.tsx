// services/RegisterService.ts
import api from '../../Api/Api';
import { Response } from '../../Api/Model';
import { Services, Shift } from '../Models/Catalogs'; // Asegúrate de que la ruta sea correcta
import { SchedulesAppointmentRequest } from '../Models/Catalogs';


export const getParamsForAppointment = async (
  token?: string
): Promise<Response<{ services: Services[]; shifts: Shift[] }>> => {
  try {
    const response = await api.get('/appointment/schedules', {
      headers: {
        Authorization: token ? `Bearer ${token}` : '',
      },
    });
    return response.data as Response<{ services: Services[]; shifts: Shift[] }>;
  } catch (error: any) {
    console.error('Error al obtener los parámetros:', error);
    throw new Error(error.message || 'Error al obtener los parámetros');
  }
};

export const createAppointment = async (
  appointmentData: SchedulesAppointmentRequest,
  token?: string
): Promise<Response<any>> => { // Reemplaza "any" con el tipo de respuesta adecuado
  try {
    const response = await api.post('/appointment/schedules', appointmentData, {
      headers: {
        Authorization: token ? `Bearer ${token}` : '',
      },
    });
    return response.data as Response<any>;
  } catch (error: any) {
    console.error('Error al crear la cita:', error);
    throw new Error(error.message || 'Error al crear la cita');
  }
};