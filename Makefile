# Makefile

CONTAINER_ID=167b8770477b
USER=root
DATABASE=university_db

# entra a la db
showdb:
	@docker exec -it $(CONTAINER_ID) psql -U $(USER) -d $(DATABASE)

