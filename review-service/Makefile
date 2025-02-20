postgres:
	docker run --name postgres17 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=mahanta -d postgres:17-alpine

createdb:
	docker exec -it postgres17 createdb --username=root --owner=root review_service

dropdb: 
	docker exec -it postgres17 dropdb review_service

migrateup:
	migrate -path db/migration -database "postgresql://root:mahanta@localhost:5432/review_service?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:2DHKOABiwRKoKVgsYnLHa0yoOOyHiHoR@dpg-cure72qj1k6c73b1gscg-a.oregon-postgres.render.com/review_service_ggxp" -verbose down

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

mock:
	mockgen -package mockdb -destination db/mock/store.go review-service/db/sqlc Store
	mockgen -package mockhlp -destination helpers/mock/helpers.go review-service/helpers Helpers
	mockgen -package mockcl -destination clients/mock/client.go review-service/clients Client
	mockgen -package mocklb -destination leaderboard/mock/leaderboard.go review-service/leaderboard Leaderboard

test:
	go test -v -cover ./...

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