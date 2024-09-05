import axios from 'axios';
import { RegisterDoctorByAdminRequest } from '../models.tsx/users'; // Asegúrate de que la ruta es correcta

const API_URL = 'http://localhost:8080'; // Cambia esto a la URL de tu API

// Servicio para registrar un usuario por parte del administrador
export const registerDoctorByAdmin = async (userData: RegisterDoctorByAdminRequest, token: string) => {
    try {
        // Configurar los encabezados de la solicitud, incluyendo el token
        const config = {
            headers: {
                Authorization: `Bearer ${token}`, // Aquí se envía el token
                'Content-Type': 'application/json', // Este es el encabezado común para JSON
            }
        };

        // Enviar la solicitud con los datos del usuario y el token en los encabezados
        const response = await axios.post(`${API_URL}/v1/doctors`, userData, config);
        return response.data;
    } catch (error) {
        console.error("Error registrando al usuario", error);
        throw error;
    }
};