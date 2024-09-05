import React, { useState } from 'react';
import { RegisterUserByAdminRequest } from '../models.tsx/users';
import { registerUserByAdmin } from '../services/userService';

function RegisterUser() {
    const [data, setData] = useState<RegisterUserByAdminRequest>({
        name: '',
        lastname1: '',
        lastname2: '',
        email: '',
        password: '',
        rol: 0, // Valor inicial como 0 para "Seleccionar"
        curp: '',
        admin_password: ''
    });

    const [errors, setErrors] = useState<{ [key: string]: boolean }>({});
    const [isSuperUserCardVisible, setSuperUserCardVisible] = useState(false);
    const [superUserPassword, setSuperUserPassword] = useState('');

    const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
        const { name, value } = e.target;

        setData({
            ...data,
            [name]: name === 'rol' ? parseInt(value) : value,
        });

        if (errors[name]) {
            setErrors({
                ...errors,
                [name]: false,
            });
        }
    };

    const handleSuperUserPasswordChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setSuperUserPassword(e.target.value);
        if (errors['admin_password']) {
            setErrors({ ...errors, admin_password: false });
        }
    };

    const validateForm = () => {
        const newErrors: { [key: string]: boolean } = {};

        if (!data.name) newErrors.name = true;
        if (!data.lastname1) newErrors.lastname1 = true;
        if (!data.lastname2) newErrors.lastname2 = true;
        if (!data.email) newErrors.email = true;
        if (!data.password) newErrors.password = true;
        if (!data.curp) newErrors.curp = true;
        if (data.rol === 0) newErrors.rol = true;

        setErrors(newErrors);
        return Object.keys(newErrors).length === 0;
    };

    const handleInitialSubmit = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        if (validateForm()) {
            setSuperUserCardVisible(true);
        } else {
            console.log("Errores en el formulario:", errors);
        }
    };

    const handleFinalSubmit = async () => {
        if (!superUserPassword) {
            setErrors({ ...errors, admin_password: true });
            return;
        }

        const updatedData = { ...data, admin_password: superUserPassword };

        const token = localStorage.getItem('token');
        if (token) {
            try {
                console.log('Datos del usuario:', updatedData);
                const response = await registerUserByAdmin(updatedData, token);
                console.log('Registro exitoso:', response);
                setSuperUserCardVisible(false);
            } catch (error) {
                console.error('Error en el registro:', error);
            }
        } else {
            console.error('No se encontró el token de autenticación.');
        }
    };

    return (
        <div className="register-page">
            <form className="register-form" onSubmit={handleInitialSubmit}>
                <h2>Registro de Usuario</h2>
                <input
                    type="text"
                    name="name"
                    placeholder="Nombre"
                    value={data.name}
                    onChange={handleChange}
                    className={errors.name ? 'input-error' : ''}
                    required
                />
                {errors.name && <span className="error-text">Este campo es obligatorio</span>}

                <input
                    type="text"
                    name="lastname1"
                    placeholder="Primer Apellido"
                    value={data.lastname1}
                    onChange={handleChange}
                    className={errors.lastname1 ? 'input-error' : ''}
                    required
                />
                {errors.lastname1 && <span className="error-text">Este campo es obligatorio</span>}

                <input
                    type="text"
                    name="lastname2"
                    placeholder="Segundo Apellido"
                    value={data.lastname2}
                    onChange={handleChange}
                    className={errors.lastname2 ? 'input-error' : ''}
                    required
                />
                {errors.lastname2 && <span className="error-text">Este campo es obligatorio</span>}

                <input
                    type="email"
                    name="email"
                    placeholder="Correo Electrónico"
                    value={data.email}
                    onChange={handleChange}
                    className={errors.email ? 'input-error' : ''}
                    required
                />
                {errors.email && <span className="error-text">Este campo es obligatorio</span>}

                <input
                    type="password"
                    name="password"
                    placeholder="Contraseña"
                    value={data.password}
                    onChange={handleChange}
                    className={errors.password ? 'input-error' : ''}
                    required
                />
                {errors.password && <span className="error-text">Este campo es obligatorio</span>}

                <input
                    type="text"
                    name="curp"
                    placeholder="CURP"
                    value={data.curp}
                    onChange={handleChange}
                    className={errors.curp ? 'input-error' : ''}
                    required
                />
                {errors.curp && <span className="error-text">Este campo es obligatorio</span>}

                <select
                    name="rol"
                    value={data.rol}
                    onChange={handleChange}
                    className={errors.rol ? 'input-error' : ''}
                    required
                >
                    <option value={0}>Seleccionar</option>
                    <option value={1}>SuperUsuario</option>
                    <option value={3}>Paciente</option>
                </select>
                {errors.rol && <span className="error-text">Este campo es obligatorio</span>}

                <button type="submit">Registrar Usuario</button>
            </form>

            {isSuperUserCardVisible && (
                <div className="superuser-card">
                    <h3>Confirmación del SuperUsuario</h3>
                    <input
                        type="password"
                        name="superuser-password"
                        placeholder="Contraseña del SuperUsuario"
                        value={superUserPassword}
                        onChange={handleSuperUserPasswordChange}
                        className={errors.admin_password ? 'input-error' : ''}
                        required
                    />
                    {errors.admin_password && <span className="error-text">Este campo es obligatorio</span>}
                    <button onClick={handleFinalSubmit}>Confirmar Registro</button>
                    <button onClick={() => setSuperUserCardVisible(false)}>Cancelar</button>
                </div>
            )}
        </div>
    );
}

export default RegisterUser;
