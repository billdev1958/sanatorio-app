import { ConfirmationData, RegisterBeneficiaryRequest, RegisterPatientRequest } from '../models/Login&Registers';
import { ScheduleAppointment } from '../models/Horarios';
import api from '../Api/Api';


export async function registerUser(user: RegisterPatientRequest) {
  try {
    const response = await api.post('/patients', user);
    console.log('Registration successful:', response.data);
    return response.data;
  } catch (error: any) {
    if (error.response) {
      console.error('Error during registration:', error.response.data);
      throw new Error(error.response.data.message || 'Registration failed');
    } else {
      console.error('Network error:', error);
      throw new Error('Network error, please try again later.');
    }
  }
}

export async function registerBeneficiary(user: RegisterBeneficiaryRequest, token?: string) {
  try {
    const headers = token ? { Authorization: `Bearer ${token}` } : {};

    const response = await api.post('/beneficiary', user, { headers });

    console.log('Registro exitoso:', response.data);
    return response.data;
  } catch (error: any) {
    if (error.response) {
      console.error('Error durante el registro:', error.response.data);
      throw new Error(error.response.data.message || 'Falló el registro');
    } else {
      console.error('Error de red:', error);
      throw new Error('Error de red, intenta nuevamente más tarde.');
    }
  }
}

export async function getScheduleAppointment(appointment: ScheduleAppointment, token?: string) {
  try {
    const headers = token ? { Authorization: `Bearer ${token}` } : {};

    const response = await api.post('/appointment/schedules', appointment, { headers });

    console.log('Horarios obtenidos exitosamente:', response.data);
    return response.data; // Retornamos los datos recibidos
  } catch (error: any) {
    if (error.response) {
      console.error('Error al obtener los horarios:', error.response.data);
      throw new Error(error.response.data.message || 'Error al obtener los horarios');
    } else {
      console.error('Error de red:', error);
      throw new Error('Error de red, intenta nuevamente más tarde.');
    }
  }
}

/**
 * Enviar el código de verificación al correo electrónico
 * @param email Email del usuario
 */
export async function sendVerificationEmail(email: string) {
  try {
    const response = await api.post('/confirmation/forward', { email });
    console.log('✅ Código de verificación enviado:', response.data);
    return response.data;
  } catch (error: any) {
    if (error.response) {
      console.error('❌ Error al reenviar código:', error.response.data);
      throw new Error(error.response.data.message || 'No se pudo reenviar el código.');
    } else {
      console.error('❌ Error de red:', error);
      throw new Error('Error de red, inténtalo de nuevo más tarde.');
    }
  }
}

/**
 * Verificar el código ingresado por el usuario
 * @param data Objeto con email y código de verificación
 */
export async function verifyCode(data: ConfirmationData) {
  try {
    const response = await api.post('/confirmation/verify', data);
    console.log('✅ Código verificado correctamente:', response.data);
    return response.data;
  } catch (error: any) {
    if (error.response) {
      console.error('❌ Error al verificar código:', error.response.data);
      throw new Error(error.response.data.message || 'El código es inválido o ha expirado.');
    } else {
      console.error('❌ Error de red:', error);
      throw new Error('Error de red, inténtalo de nuevo más tarde.');
    }
  }
}