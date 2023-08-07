build:
	@echo "Building project...\n"
	@go build -o bin/main cmd/main.go
	@echo "Project built...\n"
run:
	@go run ./bin/main
