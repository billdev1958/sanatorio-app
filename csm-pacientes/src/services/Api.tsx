import axios from 'axios';

const api = axios.create({
    baseURL: 'VITE_BACKEND_HOST',
    headers: {
      'Content-Type': 'application/json',
    },
    timeout: 10000, // Tiempo máximo en milisegundos (10 segundos)
  });

  export default api;