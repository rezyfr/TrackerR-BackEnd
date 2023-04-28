DB_URL=postgresql://root:postgres@localhost:5432/trackerr?sslmode=disable

postgres:
	docker run --name postgres14 --network trackerr-network -e POSTGRES_USER=root -e POSTGRES_PASSWORD=postgres -d -p 5432:5432 postgres:14-alpine
createdb:
	docker exec -it postgres14 createdb --username=root --owner=root trackerr

dropdb:
	docker exec -it postgres14 dropdb trackerr

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown: 
	migrate -path db/migration -database "$(DB_URL)" -verbose down

sqlc:
	sqlc generate	

test:
	go test -v -cover ./...
	
server: 
	go run main.go

mock: 
	mockgen -package mockdb -destination db/mock/store.go github.com/rezyfr/Trackerr-BackEnd/db/sqlc Store
.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server mock