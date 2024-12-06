include .env 

DSN=postgres://$(DATABASE_USER):$(DATABASE_PASSWORD)@$(DATABASE_HOST):$(DATABASE_PORT)/$(DATABASE_NAME)?sslmode=disable

migrateup:
	migrate -path ./database/migration -database $(DSN) -verbose up

migratedown:
	migrate -path ./database/migration -database $(DSN) -verbose down

migratefix:
	migrate -path ./database/migration -database $(DSN) force 1 

runapp:
	go run ./cmd/app/main.go
