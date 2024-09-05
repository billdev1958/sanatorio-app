import React, { useState } from 'react';
import { RegisterDoctorByAdminRequest } from '../models.tsx/users';
import { registerDoctorByAdmin } from '../services/doctorService';

function RegisterDoctor() {
    const [data, setData] = useState<RegisterDoctorByAdminRequest>({
        name: '',
        lastname1: '',
        lastname2: '',
        email: '',
        password: '',
        rol: 2, // Valor por defecto para "Doctor"
        medical_license: '',
        specialty: 0, // Valor inicial como 0 para "Seleccionar"
        admin_password: ''
    });

    const [errors, setErrors] = useState<{ [key: string]: boolean }>({});
    const [isSuperUserCardVisible, setSuperUserCardVisible] = useState(false); // Para controlar la visibilidad del formulario de superusuario
    const [superUserPassword, setSuperUserPassword] = useState(''); // Contraseña del superusuario

    const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
        const { name, value } = e.target;

        // Convertir 'specialty' o 'rol' a número
        setData({
            ...data,
            [name]: name === 'specialty' || name === 'rol' ? parseInt(value) : value,
        });

        // Limpiar el error de validación al ingresar datos
        if (errors[name]) {
            setErrors({
                ...errors,
                [name]: false,
            });
        }
    };

    const handleSuperUserPasswordChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setSuperUserPassword(e.target.value);
        // Limpiar error si la contraseña fue ingresada
        if (errors['admin_password']) {
            setErrors({ ...errors, admin_password: false });
        }
    };

    // Validar campos
    const validateForm = () => {
        const newErrors: { [key: string]: boolean } = {};

        if (!data.name) newErrors.name = true;
        if (!data.lastname1) newErrors.lastname1 = true;
        if (!data.lastname2) newErrors.lastname2 = true;
        if (!data.email) newErrors.email = true;
        if (!data.password) newErrors.password = true;
        if (!data.medical_license) newErrors.medical_license = true;
        if (data.specialty === 0) newErrors.specialty = true; // Validación para specialty

        setErrors(newErrors);
        return Object.keys(newErrors).length === 0; // Si no hay errores, retorna true
    };

    // Primer submit para validar los datos del formulario
    const handleInitialSubmit = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        if (validateForm()) {
            setSuperUserCardVisible(true); // Si es válido, muestra el campo para la contraseña de superusuario
        } else {
            console.log("Errores en el formulario:", errors);
        }
    };

    // Segundo submit para enviar los datos completos
    const handleFinalSubmit = async () => {
        // Validar la contraseña de superusuario
        if (!superUserPassword) {
            setErrors({ ...errors, admin_password: true });
            return;
        }

        // Agrega la contraseña del superusuario a los datos
        const updatedData = { ...data, admin_password: superUserPassword };

        const token = localStorage.getItem('token'); // Obtener el token desde el localStorage
        if (token) {
            try {
                console.log('Datos del usuario:', updatedData);
                const response = await registerDoctorByAdmin(updatedData, token);
                console.log('Registro exitoso:', response);
                setSuperUserCardVisible(false); // Ocultar la tarjeta después del éxito
            } catch (error) {
                console.error('Error en el registro:', error);
            }
        } else {
            console.error('No se encontró el token de autenticación.');
        }
    };

    return (
        <div className="register-page">
            {/* Formulario inicial */}
            <form className="register-form" onSubmit={handleInitialSubmit}>
                <h2>Registro de Doctor</h2>
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
                    name="medical_license"
                    placeholder="Licencia Médica"
                    value={data.medical_license}
                    onChange={handleChange}
                    className={errors.medical_license ? 'input-error' : ''}
                    required
                />
                {errors.medical_license && <span className="error-text">Este campo es obligatorio</span>}

                <select
                    name="specialty"
                    value={data.specialty}
                    onChange={handleChange}
                    className={errors.specialty ? 'input-error' : ''}
                    required
                >
                    <option value={0}>Seleccionar</option> {/* Opción por defecto */}
                    <option value={1}>Cardiologo</option>
                    <option value={2}>Dermatologo</option>
                    <option value={3}>Pediatra</option>
                    <option value={4}>Ginecologia</option>
                </select>
                {errors.specialty && <span className="error-text">Este campo es obligatorio</span>}

                <button type="submit">Registrar Usuario</button>
            </form>

            {/* Formulario de confirmación del superusuario */}
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

export default RegisterDoctor;
