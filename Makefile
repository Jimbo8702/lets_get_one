build:
	@go build -o bin/app ./cmd

run: build
	@./bin/app
test: 
	@go test -v ./...

postgres: 
	docker run --name postgres20 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

db: 
	docker exec -it postgres20 createdb --username=root --owner=root temp_simple

drop: 
	docker exec -it postgres20 dropdb temp_simple

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/temp_simple?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/temp_simple?sslmode=disable" -verbose down

.PHONY: postgres db drop migrateup migratedown 