postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e  POSTGRES_PASSWORD=secret -d postgres:latest

createdb:
	docker exec -it postgres createdb --username=root --owner=root auto_care

dropdb:
	docker exec -it postgres dropdb auto_care

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/auto_care?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/auto_care?sslmode=disable" -verbose down

migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/auto_care?sslmode=disable" -verbose up 1


migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/auto_care?sslmode=disable" -verbose down 1


sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/IvanRoussev/autocare/db/sqlc Store
.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc test server mock