// services/RegisterService.ts
import api from '../../Api/Api';
import { Response } from '../../Api/Model';
import { Services, Shift } from '../Models/Catalogs'; // Asegúrate de que la ruta sea correcta

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