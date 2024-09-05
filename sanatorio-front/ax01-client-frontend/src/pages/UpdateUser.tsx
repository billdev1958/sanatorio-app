import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom'; 
import { UpdateUserRequest } from '../models.tsx/users'; 
import { updateUser } from '../services/updateUserService'; 
import { getUserById } from '../services/getUsers'; 

function UpdateUser() {
    const { userId } = useParams<{ userId: string }>(); 

    const [data, setData] = useState<UpdateUserRequest>({
        account_id: '', 
        name: '',
        lastname1: '',
        lastname2: '',
        email: '',
        password: '',
        curp: '',
        admin_password: '', 
    });

    const [originalData, setOriginalData] = useState<UpdateUserRequest>({
        account_id: '',
        name: '',
        lastname1: '',
        lastname2: '',
        email: '',
        password: '',
        curp: '',
        admin_password: '',
    });

    const [isAdminConfirmationCardVisible, setAdminConfirmationCardVisible] = useState(false); 
    const [superUserPassword, setSuperUserPassword] = useState(''); 

    useEffect(() => {
        const fetchUserData = async () => {
            const token = localStorage.getItem('token');
            console.log('Token:', token); 
            console.log('UserID:', userId); 

            if (token && userId) { 
                try {
                    const userData = await getUserById(userId, token); 
                    setData({ 
                        ...userData, 
                        account_id: userData.account_id || '', 
                        admin_password: '' 
                    });
                    setOriginalData({
                        ...userData,
                        account_id: userData.account_id || '', 
                        admin_password: '',
                    }); 
                } catch (error) {
                    console.error('Error al obtener los datos del usuario:', error);
                }
            } else {
                console.error('No se encontró el token de autenticación o userId.');
            }
        };

        fetchUserData();
    }, [userId]);

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const value = e.target.value || ''; 
        setData({
            ...data,
            [e.target.name]: value,
        });
    };

    const handleSuperUserPasswordChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setSuperUserPassword(e.target.value || ''); 
    };

    const handleInitialSubmit = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();

        const noChanges = Object.keys(data).every(
            key => data[key as keyof UpdateUserRequest] === originalData[key as keyof UpdateUserRequest]
        );

        if (noChanges) {
            alert('No has modificado nada.');
        } else {
            setAdminConfirmationCardVisible(true);
        }
    };

    const handleFinalSubmit = async () => {
        const updatedData = { ...data, admin_password: superUserPassword };

        const finalData = Object.keys(updatedData).reduce((acc, key) => {
            if (updatedData[key as keyof UpdateUserRequest] === originalData[key as keyof UpdateUserRequest]) {
                acc[key as keyof UpdateUserRequest] = "";
            } else {
                acc[key as keyof UpdateUserRequest] = updatedData[key as keyof UpdateUserRequest] || ''; 
            }
            return acc;
        }, {} as UpdateUserRequest);

        finalData.account_id = data.account_id || ''; 

        const token = localStorage.getItem('token'); 
        if (token) {
            try {
                console.log('Datos finales para la actualización:', finalData);

                const response = await updateUser(finalData, token); 
                console.log('Actualización exitosa:', response);
                setAdminConfirmationCardVisible(false); 
            } catch (error) {
                console.error('Error en la actualización:', error);
            }
        } else {
            console.error('No se encontró el token de autenticación.');
        }
    };

    return (
        <div className="update-page">
            <h2>Actualización de Usuario</h2>
            <h4>Número de cuenta del usuario: {data.account_id}</h4>
            <form className="update-form" onSubmit={handleInitialSubmit}>
                <input
                    type="text"
                    name="name"
                    placeholder="Nombre"
                    value={data.name || ''} 
                    onChange={handleChange}
                />
                <input
                    type="text"
                    name="lastname1"
                    placeholder="Primer Apellido"
                    value={data.lastname1 || ''} 
                    onChange={handleChange}
                />
                <input
                    type="text"
                    name="lastname2"
                    placeholder="Segundo Apellido"
                    value={data.lastname2 || ''} 
                    onChange={handleChange}
                />
                <input
                    type="email"
                    name="email"
                    placeholder="Correo Electrónico"
                    value={data.email || ''} 
                    onChange={handleChange}
                />
                <input
                    type="password"
                    name="password"
                    placeholder="Contraseña"
                    value={data.password || ''} 
                    onChange={handleChange}
                />
                <input
                    type="text"
                    name="curp"
                    placeholder="CURP"
                    value={data.curp || ''} 
                    onChange={handleChange}
                />
                <button type="submit">Actualizar Usuario</button>
            </form>

            {isAdminConfirmationCardVisible && (
                <div className="admin-confirmation-card">
                    <h3>Confirmación del SuperUsuario</h3>
                    <input
                        type="password"
                        name="superuser-password"
                        placeholder="Contraseña del SuperUsuario"
                        value={superUserPassword || ''} 
                        onChange={handleSuperUserPasswordChange}
                        required
                    />
                    <button onClick={handleFinalSubmit}>Confirmar Actualización</button>
                    <button onClick={() => setAdminConfirmationCardVisible(false)}>Cancelar</button>
                </div>
            )}
        </div>
    );
}

export default UpdateUser;
