# Makefile

default: run

# Target for building the Docker images using docker-compose.
build:
	docker-compose build

# Target for starting the services defined in docker-compose.yml.
up:
	docker-compose up

run:
	docker-compose build
	docker-compose up

.PHONY: build up run
