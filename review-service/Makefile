postgres:
	docker run --name postgres17 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=mahanta -d postgres:17-alpine

createdb:
	docker exec -it postgres17 createdb --username=root --owner=root review_service

dropdb: 
	docker exec -it postgres17 dropdb review_service

migrateup:
	migrate -path db/migration -database "postgresql://root:mahanta@localhost:5432/review_service?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:mahanta@localhost:5432/review_service?sslmode=disable" -verbose down

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)


sqlc:
	sqlc generate

server:
	go run main.go

evans:
	evans --host localhost --port 9090 -r

proto:
	rm -f payment/*.go
	protoc --proto_path=proto --go_out=payment --go_opt=paths=source_relative \
    --go-grpc_out=payment --go-grpc_opt=paths=source_relative \
    proto/*.proto


.PHONY: postgres createdb dropdb migrateup new_migration sqlc server evans proto