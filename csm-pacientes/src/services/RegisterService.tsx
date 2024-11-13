import api from './Api';
import { RegisterBeneficiaryRequest, RegisterPatientRequest } from '../models/Login&Registers';
// import { useAuth } from './AuthContext';


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
    // Configurar los encabezados de autorización solo si hay un token
    const headers = token ? { Authorization: `Bearer ${token}` } : {};

    // Enviar la solicitud con el encabezado Authorization si el token está disponible
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