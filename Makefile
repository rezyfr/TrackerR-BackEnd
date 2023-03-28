postgres:
	docker run --name postgres15 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=postgres -d -p 5432:5432 postgres:15
createdb:
	docker exec -it postgres15 createdb --username=root --owner=root trackerr

dropdb:
	docker exec -it postgres15 dropdb trackerr

migrateup:
	migrate -path db/migration -database "postgresql://root:postgres@localhost:5432/trackerr?sslmode=disable" -verbose up

migratedown: 
	migrate -path db/migration -database "postgresql://root:postgres@localhost:5432/trackerr?sslmode=disable" -verbose down

sqlc:
	sqlc generate	

test:
	go test -v -cover ./...
	
.PHONY: postgres createdb dropdb migrateup migratedown sqlc test