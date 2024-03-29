DB_URL=postgresql://root:password@localhost:5431/go_etl?sslmode=disable

postgres:
	docker run --name postgres -p 5431:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:15.2-alpine
redis:
	docker run --name redis -p 6379:6379 -d redis:7.0.9-alpine
create_db:
	docker exec -it postgres createdb --username=root --owner=root ${DB_NAME}
migrate_create:
	migrate create -ext sql -dir db/migrations -seq ${MIGRATE_NAME}
migrate_up:
	migrate -path db/migrations -database "${DB_URL}" -verbose up
migrate_down:
	migrate -path db/migrations -database "${DB_URL}" -verbose down
sqlc:
	sqlc generate
tests:
	go test -v -cover ./...
mocks:
	mockgen -package mockdb --destination db/mocks/store.go github.com/ShadrackAdwera/go-etl/db/sqlc TxStore
start:
	go run main.go

.PHONY: create_db migrate_create migrate_up migrate_down sqlc tests mocks start postgres redis