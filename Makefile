main:
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
migration:
ifndef NAME
	$(error NAME is undefined. Pass it like this: make migration NAME=my_migration)
endif
	@migrate create -ext sql -dir migrations/ -seq $(NAME)
migrate-up:
	@migrate -path db/migrations/ -database "postgresql://user:password@localhost:5432/prom-ler-db?sslmode=disable" -verbose up
migrate-down:
	@migrate -path db/migrations/ -database "postgresql://user:password@localhost:5432/prom-ler-db?sslmode=disable" -verbose down
migration-fix:
ifndef VERSION
	$(error VERSION is undefined. Pass it like this: make migration-fix VERSION=00001)
endif
	@migrate -path migrations/ -database "postgresql://postgres:password@localhost:5432/prom-ler-db?sslmode=disable" force $(VERSION)
