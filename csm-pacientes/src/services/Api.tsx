import axios from 'axios';

const api = axios.create({
    baseURL: 'http://api.ax01.dev',
    headers: {
      'Content-Type': 'application/json',
    },
    timeout: 10000, // Tiempo máximo en milisegundos (10 segundos)
  });

  export default api;