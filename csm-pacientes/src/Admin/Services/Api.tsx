import axios from 'axios';

const api = axios.create({
    baseURL: 'https://api.ax01.dev/v1',
    headers: {
      'Content-Type': 'application/json',
    },
    timeout: 10000, // Tiempo máximo en milisegundos (10 segundos)
  });

  export default api;