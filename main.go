package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/rezyfr/Trackerr-BackEnd/api"
	db "github.com/rezyfr/Trackerr-BackEnd/db/sqlc"
	"github.com/rezyfr/Trackerr-BackEnd/util"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}
	var dbURI string
	dbURI = fmt.Sprintf("host=%s user=%s password=%s port=%s database=%s", config.DBHost, config.DBUser, config.DBPassword, config.DBPort, config.DBName)
	conn, err := sql.Open(config.DBDriver, dbURI)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
