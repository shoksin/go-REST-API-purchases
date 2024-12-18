version := latest

docker-build:
	docker build -t go-purchases-rest-api:$(version) .

docker-run:
	docker run --name $(container_name) go-purchases-rest-api:$(version)

find-volumes:
	cd \\wsl$\docker-desktop\mnt\docker-desktop-disk\data\docker\volumes

open-db:
	docker exec -it purchases-db psql -U $(DB_USER) -d $(DB_NAME)
# make open-db DB_USER=username DB_NAME=dbname