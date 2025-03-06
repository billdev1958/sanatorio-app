import axios from 'axios';

// Production
const FRONTEND_HOST:string = "https://api.ax01.dev/v1"

// Dev
// const FRONTEND_HOST = "http://localhost:8080"


const api = axios.create({
    baseURL: FRONTEND_HOST,
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