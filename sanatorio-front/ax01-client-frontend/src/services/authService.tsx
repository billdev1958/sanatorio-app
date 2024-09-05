import axios from 'axios';
import { LoginUser } from '../models.tsx/users'; // Asegúrate de que la ruta es correcta

const API_URL = 'http://localhost:8080'; // Cambia esto a la URL de tu API

export const loginUser = async (userData: LoginUser) => {
    try {
        console.log('Datos enviados en el login:', userData);

        const response = await axios.post(`${API_URL}/v1/login`, userData);
        console.log('Respuesta del servidor:', response.data);

        const token = response.data.data.token; // Accediendo correctamente al token en la estructura de la respuesta
        if (token) {
            localStorage.setItem('token', token); // Guarda el token en localStorage
            console.log('Token guardado:', token); // Log para confirmar que el token se guardó
        } else {
            console.error('No se recibió token en la respuesta');
        }
        return response.data;
    } catch (error) {
        console.error('Error en el login:', error);
        throw error;
    }
};

// Servicio para hacer solicitudes autenticadas con JWT
export const getAuthenticatedData = async (endpoint: string) => {
    try {
        const token = localStorage.getItem('token');
        if (!token) {
            throw new Error('No se encontró el token');
        }

        const config = {
            headers: { Authorization: `Bearer ${token}` }
        };

        const response = await axios.get(`${API_URL}${endpoint}`, config);
        
        // Imprime la respuesta recibida del servidor
        console.log('Respuesta del servidor (Authenticated):', response.data);

        return response.data;
    } catch (error) {
        console.error('Error en la solicitud autenticada:', error);
        throw error;
    }
};

// Servicio para cerrar sesión y eliminar el token JWT
export const logoutUser = () => {
    localStorage.removeItem('token');
    console.log('Usuario deslogueado');
};