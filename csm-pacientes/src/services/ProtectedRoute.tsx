import { useAuth } from '../services/AuthContext';
import { Navigate } from '@solidjs/router';
import { JSX } from 'solid-js';

const ProtectedRoute = (props: { children: JSX.Element }) => {
  const auth = useAuth();

  if (!auth.isAuthenticated()) {
    // Redirigir al login si el usuario no está autenticado
    return <Navigate href="/login" />;
  }

  // Renderizar el contenido de la ruta si el usuario está autenticado
  return props.children;
};

export default ProtectedRoute;
