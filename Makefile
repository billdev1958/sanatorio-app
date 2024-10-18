# Makefile

CONTAINER_ID=c79bd4f5d841
USER=root
DATABASE=university_db

# entra a la db
showdb:
	@docker exec -it $(CONTAINER_ID) psql -U $(USER) -d $(DATABASE)

