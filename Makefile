# Makefile

CONTAINER_ID=5cc5250ed367
USER=root
DATABASE=university_db

# entra a la db
showdb:
	@docker exec -it $(CONTAINER_ID) psql -U $(USER) -d $(DATABASE)

