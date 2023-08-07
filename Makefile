build:
	@echo "Building project...\n"
	@go build -o bin/main cmd/main.go
	@echo "Project built...\n"
run:
	@go run ./bin/main
dev.start:
	@docker-compose up -d
dev.stop:
	@docker-compose stop
dev.down:
	@docker-compose down
