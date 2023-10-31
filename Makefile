default: run

run:
	docker-compose build --build-arg fast=true
	docker-compose up

run-local:
	go run cmd/main.go

test:
	go test ./... -short -cover -race

prometheus:
	prometheus --config.file=prometheus.yml