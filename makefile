DB_CONTAINER_NAME=inventory_system
DB_USER_ROOT=poncho
DB_ROOT_PASSSWORD=SecurePassword
DB_NAME=inventory_system
CONTAINER_INTERNAL_PORT=5432
CONTAINER_EXTERNAL_PORT=3000


Database:
	docker run --name $(DB_CONTAINER_NAME) -e POSTGRES_USER=$(DB_USER_ROOT) -e POSTGRES_PASSWORD=$(DB_ROOT_PASSSWORD) -e POSTGRES_DB=$(DB_NAME) -p $(CONTAINER_EXTERNAL_PORT):$(CONTAINER_INTERNAL_PORT) -d postgres:latest
	docker ps

Stop:
	docker stop $(DB_CONTAINER_NAME)

Start:
	docker start $(DB_CONTAINER_NAME)

Exec:
	docker exec -it $(DB_CONTAINER_NAME) psql -U $(DB_USER_ROOT) -d $(DB_NAME)

Remove:
	docker rm $(DB_CONTAINER_NAME)