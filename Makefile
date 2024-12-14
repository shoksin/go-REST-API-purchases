version := latest

docker-build:
	docker build -t go-purchases-rest-api:$(version) .

docker-run:
	docker run --name $(container_name) go-purchases-rest-api:$(version)