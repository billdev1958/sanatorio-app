import api from './Api';
import { RegisterBeneficiaryRequest, RegisterPatientRequest } from '../models/Login&Registers';

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

export async function registerBeneficiary(user: RegisterBeneficiaryRequest) {
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