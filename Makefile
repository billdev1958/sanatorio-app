name: CI/CD Pipeline

on:
  push:
    branches:
      - main

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout del código
        uses: actions/checkout@v4

      - name: Mostrar build number
        run: echo "Build number: ${{ github.run_number }}"

      - name: Login en Docker Hub
        run: echo "${{ secrets.DOCKER_PASS }}" | docker login -u "${{ secrets.DOCKER_USER }}" --password-stdin

      - name: Construir y subir imágenes a Docker Hub
        env:
          DOCKER_USER: ${{ secrets.DOCKER_USER }}
          BUILD_TAG: "1.${{ github.run_number }}"
        run: |
          echo "Tag para build: $BUILD_TAG"
          
          # Construir y subir imagen del Backend
          docker build --no-cache -t $DOCKER_USER/sanatorio-back:$BUILD_TAG ./sanatorio-back
          docker push $DOCKER_USER/sanatorio-back:$BUILD_TAG

          # Construir y subir imagen del Frontend
          docker build --no-cache -t $DOCKER_USER/csm-pacientes:$BUILD_TAG ./csm-pacientes
          docker push $DOCKER_USER/csm-pacientes:$BUILD_TAG

  deploy:
    needs: build
    runs-on: ubuntu-latest
    env:
      BUILD_TAG: "1.${{ github.run_number }}"
    steps:
      - name: Desplegar en el servidor vía SSH
        uses: appleboy/ssh-action@v1.0.3
        with:
          host: ${{ secrets.SERVER_HOST }}
          username: ${{ secrets.SERVER_USER }}
          key: ${{ secrets.SERVER_SSH_KEY }}
          envs: BUILD_TAG
          script: |
            set -e  # Detener el script en caso de error

            # Moverse al directorio del proyecto
            cd /home/bill/sanatorio-app
            
            # Actualizar repositorio con los últimos cambios
            git reset --hard HEAD
            git pull origin main

            # Verificar el tag antes de usarlo
            echo "Usando tag: ${BUILD_TAG}"

            # Actualizar las imágenes en docker-compose-production.yml usando el tag serial
            sed -i "s|image: billdev1958/sanatorio-back:.*|image: billdev1958/sanatorio-back:${BUILD_TAG}|g" docker-compose-production.yml
            sed -i "s|image: billdev1958/csm-pacientes:.*|image: billdev1958/csm-pacientes:${BUILD_TAG}|g" docker-compose-production.yml

            # Verificar que el archivo se actualizó correctamente
            cat docker-compose-production.yml | grep "image: billdev1958"

            # Reiniciar entorno de producción
            make prod-reset

            # Descargar nuevas imágenes y levantar los contenedores con Makefile
            make prod BUILD_TAG=${BUILD_TAG}

            # Limpiar imágenes y archivos innecesarios
            docker system prune -f
