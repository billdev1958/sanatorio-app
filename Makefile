# Makefile
CONTAINER_NAME=sanatorio-app-db-1
USER=root
DATABASE=university_db

# Función para obtener el ID del contenedor basado en el nombre
CONTAINER_ID=$(shell docker ps -qf "name=$(CONTAINER_NAME)")

# Entra a la base de datos para revisar registros
showdb:
	@docker exec -it $(CONTAINER_ID) psql -U $(USER) -d $(DATABASE)

dockeri:
	@sudo dnf install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin

cloud:
	@curl -L --output cloudflared.deb https://github.com/cloudflare/cloudflared/releases/latest/download/cloudflared-linux-amd64.deb && 

	@sudo dpkg -i cloudflared.deb && 

	@sudo cloudflared service install eyJhIjoiMTQzYmNhMmRjZTk4NzA5MDk3ZTZmNjU4MmY5YjIzZGYiLCJ0IjoiMDlhMjQxMzItOTZhMi00NjFkLTliMTctZDJkZTcyYzkxYzU4IiwicyI6IlltSmhZVGhoWlRNdE0yTmtZUzAwWmpOaUxUaGxZakV0TkdRNE9UUTNZMlEwT1dKayJ9
