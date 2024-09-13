# Sanatorio-app

## Requisitos

- [Docker](https://www.docker.com/get-started)
- Git

## Instrucciones para iniciar el proyecto

1. **Instalar Docker:**

   Si no tienes Docker instalado en tu ordenador, puedes instalarlo siguiendo las instrucciones de la [documentación oficial de Docker](https://docs.docker.com/get-docker/).

2. **Clonar la rama `dev` del repositorio:**

   Abre la terminal y clona el repositorio con el siguiente comando:

   `git clone -b dev https://github.com/usuario/sanatorio-app.git`

3. **Navegar a la carpeta del proyecto:**

   Una vez clonado el repositorio, ingresa en la carpeta del proyecto:

   `cd sanatorio-app`

4. **Iniciar los contenedores con Docker Compose:**

   En la misma carpeta del proyecto, inicia los servicios de la base de datos y el backend ejecutando:

   `docker compose up`

   Esto iniciará la base de datos en el puerto 5432 y el backend en el puerto 8080.
