# Makefile

CONTAINER_ID=f4f4bc3b635c
USER=root
DATABASE=university_db

# entra a la db
showdb:
	@docker exec -it $(CONTAINER_ID) psql -U $(USER) -d $(DATABASE)

