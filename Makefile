postgres:
	docker run --name postgres12 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=postgres -d -p 5432:5432 postgres:12-alpine
createdb:
	docker exec -it postgres12 createdb --username=root --owner=root trackerr

dropdb:
	docker exec -it postgres12 dropdb trackerr

migrateup:
	migrate -path db/migration -database "postgresql://root:postgres@localhost:5432/trackerr?sslmode=disable" -verbose up

migratedown: 
	migrate -path db/migration -database "postgresql://root:postgres@localhost:5432/trackerr?sslmode=disable" -verbose down
.PHONY: postgres createdb dropdb migrateup migratedown