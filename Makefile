DB_URL=postgresql://root:password@localhost:5431/go_etl_test?sslmode=disable

create_db:
	docker exec -it postgres15 createdb --username=root --owner=root go_etl
migrate_init:
	migrate create -ext sql -dir db/migrations -seq init_schema
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

.PHONY: migrate_init migrate_up migrate_down sqlc tests mocks start