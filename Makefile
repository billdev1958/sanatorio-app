# Makefile

CONTAINER_ID=3a1bd4716133
USER=root
DATABASE=university_db

# entra a la db
showdb:
	@docker exec -it $(CONTAINER_ID) psql -U $(USER) -d $(DATABASE)

