import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    port: 3000, // Cambia el puerto si lo necesitas
    open: true, // Abre automáticamente el navegador cuando el servidor se inicie
    proxy: {
      '/api': {
        target: 'http://localhost:8080', // Proxy para redirigir API durante el desarrollo
        changeOrigin: true,
        secure: false,
      },
    },
  },
  build: {
    outDir: 'dist',
    sourcemap: true, // Generar mapas de origen para depuración
  },
});
