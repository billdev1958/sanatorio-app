FROM node:21-alpine3.18 AS build-stage

WORKDIR /app

COPY . .

RUN npm install

RUN npm run build

FROM nginx:1.25.4 as production-stage

# Copia los archivos construidos de la aplicación React
COPY --from=build-stage  /app/dist /usr/share/nginx/html

# Copia la configuración de Nginx
COPY ./default.conf /etc/nginx/conf.d/default.conf

# Exponer el puerto 80
EXPOSE 80

# Comando para iniciar Nginx cuando el contenedor se ejecute
CMD ["nginx", "-g", "daemon off;"]