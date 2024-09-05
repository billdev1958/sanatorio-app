import axios from 'axios';
import { Users, Doctor } from '../models.tsx/users';

const API_URL = 'http://localhost:8080'; // Cambia esto a la URL de tu API

export const getUsers = async (token: string): Promise<Users[]> => {
  try {
    const config = {
      headers: {
        Authorization: `Bearer ${token}`,
        'Content-Type': 'application/json',
      },
    };

    const response = await axios.get(`${API_URL}/v1/users`, config);
    
    // Asegúrate de acceder a `data.data` para obtener los usuarios.
    return response.data.data; // Aquí accedes a la lista de usuarios que está dentro de `data`
  } catch (error) {
    console.error('Error fetching users:', error);
    throw error;
  }
};

export const getUserById = async (userId: string, token: string): Promise<Users> => {
  try {
    const config = {
      headers: {
        Authorization: `Bearer ${token}`,
        'Content-Type': 'application/json',
      },
    };

    const response = await axios.get(`${API_URL}/v1/user/${userId}`, config);
    
    // Asegúrate de acceder a `data.data` para obtener el usuario.
    return response.data.data; // Aquí accedes al usuario individual que está dentro de `data`
  } catch (error) {
    console.error(`Error fetching user with ID ${userId}:`, error);
    throw error;
  }
};

export const getAllUsers = async (token: string): Promise<(Users | Doctor)[]> => {
  try {
    const config = {
      headers: {
        Authorization: `Bearer ${token}`,
        'Content-Type': 'application/json',
      },
    };

    const response = await axios.get(`${API_URL}/v1/users/all`, config);

    // Aquí accedemos a `response.data.data`, ya que es donde están los usuarios.
    return response.data.data; 
  } catch (error) {
    console.error('Error fetching all users:', error);
    throw error;
  }
};