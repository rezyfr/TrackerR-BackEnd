postgres:
	docker run --name postgres15 --network trackerr-network -e POSTGRES_USER=root -e POSTGRES_PASSWORD=postgres -d -p 5432:5432 postgres:15
createdb:
	docker exec -it postgres15 createdb --username=root --owner=root trackerr

dropdb:
	docker exec -it postgres15 dropdb trackerr

migrateup:
	migrate -path db/migration -database "postgresql://postgres:dx1%3F%281%5CMN5KY%3D%3ADm@34.101.74.76:5432/trackerr-dev" -verbose up

migratedown: 
	migrate -path db/migration -database "postgresql://postgres:dx1%3F%281%5CMN5KY%3D%3ADm@34.101.74.76:5432/trackerr-dev" -verbose down

sqlc:
	sqlc generate	

test:
	go test -v -cover ./...
	
server: 
	go run main.go

mock: 
	mockgen -package mockdb -destination db/mock/store.go github.com/rezyfr/Trackerr-BackEnd/db/sqlc Store
.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server mock