import { defineConfig } from 'vite';
import solid from 'vite-plugin-solid';

export default defineConfig({
  plugins: [solid()],
  build: {
    outDir: 'dist',  // Asegúrate de que el directorio de salida sea 'dist'
    target: 'esnext',  // Opcional: especifica el objetivo de ES para el bundle
  },
});
