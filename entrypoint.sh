#!/bin/sh

# Cargar las variables de entorno desde el archivo .env
export $(grep -v '^#' /etc/nginx/.env | xargs)

# Reemplazar las variables en nginx.conf.template y generar nginx.conf final
envsubst '$VITE_FRONTEND_HOST $VITE_BACKEND_HOST' < /etc/nginx/nginx.conf.template > /etc/nginx/nginx.conf

# Ejecutar el comando original de Nginx
exec "$@"

