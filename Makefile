default: run

run:
	docker-compose build --build-arg fast=true
	docker-compose up

run-local:
	go run cmd/main.go

.PHONY: run run-local
