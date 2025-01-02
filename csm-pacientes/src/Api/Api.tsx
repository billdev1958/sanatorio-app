import axios from 'axios';

const api = axios.create({
    baseURL: 'https://api.ax01.dev/v1',
    headers: {
      'Content-Type': 'application/json',
    },
    timeout: 10000,
  });

  api.interceptors.request.use(
    (config) => {
      const token = localStorage.getItem('token');
      if (token) {
        config.headers.Authorization = `Bearer ${token}`;
      }
      return config;
    },
    (error) => {
      return Promise.reject(error);
    }
  );

  api.interceptors.response.use(
    (response) => {
      return response;
    },
    (error) => {
      if (error.response.status === 401) {
        window.location.href = '/login';
      } else {
        console.error('Error en la solicitud:', error.message);
      }
      return Promise.reject(error);
    }
  );
  

  export default api;