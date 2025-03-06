# Variables de entorno
#COMPOSE_DEV = docker compose --env-file .env.dev -f docker-compose.yml -f docker-compose.override.yml
COMPOSE_PROD = docker compose --env-file .env.production -f docker-compose.yml -f docker-compose-production.yml

# Construir la imagen de desarrollo y ejecutar con hot reload
dev:
	@echo "🚀 Iniciando entorno de desarrollo..."
	$(COMPOSE_DEV) up -d --build

# Detener y eliminar los contenedores de desarrollo
dev-down:
	@echo "🛑 Deteniendo entorno de desarrollo..."
	$(COMPOSE_DEV) down

# Construir la imagen de producción y ejecutar
prod:
	@echo "🚀 Iniciando entorno de producción..."
	$(COMPOSE_PROD) up -d --build

# Detener y eliminar los contenedores de producción
prod-down:
	@echo "🛑 Deteniendo entorno de producción..."
	$(COMPOSE_PROD) down

# Ver los logs del backend en desarrollo
logs-dev:
	@echo "📜 Mostrando logs del backend en desarrollo..."
	docker logs -f sanatorio-app-app-1

# Ver los logs del backend en producción
logs-prod:
	@echo "📜 Mostrando logs del backend en producción..."
	docker logs -f sanatorio-app-app-1

# Construir manualmente la imagen de producción y subirla a Docker Hub
build-prod:
	@echo "🐳 Construyendo imagen de producción..."
	docker build -t billdev1958/sanatorio-back:1.0 -f sanatorio-back/Dockerfile sanatorio-back/
	@echo "🚀 Pushing a Docker Hub..."
	docker push billdev1958/sanatorio-back:1.0

# Construir la imagen de desarrollo localmente
build-dev:
	@echo "🐳 Construyendo imagen de desarrollo..."
	docker build -t billdev1958/sanatorio-back:dev -f sanatorio-back/Dockerfile.dev sanatorio-back/

# Ver las imágenes de Docker
images:
	@echo "🔍 Listando imágenes de Docker..."
	docker images

# Limpiar imágenes y contenedores no utilizados
clean:
	@echo "🧹 Limpiando contenedores, volúmenes e imágenes sin usar..."
	docker system prune -af
