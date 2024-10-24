# Etapa 1: Construcción
FROM node:20 AS build

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos de configuración y otros necesarios para la construcción
COPY package*.json ./
COPY tsconfig*.json ./
COPY vite.config.ts ./
COPY index.html ./
COPY src ./src

# Instala las dependencias y construye el proyecto
RUN npm install
RUN npm run build

# Etapa 2: Producción
FROM nginx:alpine

# Copia los archivos estáticos construidos desde la etapa anterior
COPY --from=build /app/dist /usr/share/nginx/html

# Copia la configuración personalizada de Nginx, si la tienes
COPY nginx.conf /etc/nginx/nginx.conf

# Exponer el puerto 80, que es donde Nginx corre por defecto
EXPOSE 80

# El CMD predeterminado de Nginx ya está configurado para servir los archivos
