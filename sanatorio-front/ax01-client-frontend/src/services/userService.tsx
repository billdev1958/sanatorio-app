import axios from 'axios';
import { RegisterUserByAdminRequest } from '../models.tsx/users'; // Asegúrate de que la ruta es correcta

const API_URL = 'http://localhost:8080'; // Cambia esto a la URL de tu API

// Servicio para registrar un usuario por parte del administrador
export const registerUserByAdmin = async (userData: RegisterUserByAdminRequest, token: string) => {
    try {
        const config = {
            headers: {
                Authorization: `Bearer ${token}`, // Aquí se envía el token
                'Content-Type': 'application/json', // Encabezado para JSON
            }
        };

        // Envía la solicitud con los datos del usuario y el token en los encabezados
        const response = await axios.post(`${API_URL}/v1/users`, userData, config);
        return response.data;
    } catch (error: any) {
        console.error("Error registrando al usuario", error.response?.data || error.message);
        throw error;
    }
};
