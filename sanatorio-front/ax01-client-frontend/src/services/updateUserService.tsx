import { UpdateUserRequest } from '../models.tsx/users'; // Importar el tipo de datos para la actualización

export const updateUser = async (data: UpdateUserRequest, token: string) => {
    try {
        const response = await fetch(`http://localhost:8080/v1/users`, {
            method: 'PUT',
            headers: {
                Authorization: `Bearer ${token}`,
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(data),
        });

        if (!response.ok) {
            throw new Error(`Error en la solicitud: ${response.statusText}`);
        }

        const result = await response.json();
        return result;
    } catch (error) {
        throw new Error(`Error al actualizar el usuario: ${error}`);
    }
};
