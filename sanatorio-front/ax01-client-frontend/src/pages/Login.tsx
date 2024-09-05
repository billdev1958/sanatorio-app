import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { loginUser } from '../services/authService';
import { LoginUser } from '../models.tsx/users';

function Login() {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const navigate = useNavigate(); // Inicializa useNavigate

    const handleLogin = async () => {
        const userData: LoginUser = {
            email,
            password,
        };

        try {
            const response = await loginUser(userData); // Aquí usamos el método login
            if (response && response.data && response.data.token) {
                const token = response.data.token;
                console.log('Login exitoso:', token);

                // Guarda el token en localStorage
                localStorage.setItem('token', token);

                // Redirige a la ruta raíz después de un login exitoso
                navigate('/');
            } else {
                console.error('Error en el login: Token no recibido');
            }
        } catch (error) {
            console.error('Error en el login:', error);
        }
    };

    return (
        <div className="login-page">
            <div className="login-container">
                <h2>Iniciar Sesión</h2>
                <input
                    type="email"
                    placeholder="Correo Electrónico"
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                />
                <input
                    type="password"
                    placeholder="Contraseña"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                />
                <button onClick={handleLogin}>Iniciar Sesión</button>
            </div>
        </div>
    );
}

export default Login;
