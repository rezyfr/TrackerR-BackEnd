package main

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/rezyfr/Trackerr-BackEnd/util"
	"log"
	"net/http"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	var dbURI string
	dbURI = fmt.Sprintf("host=%s user=%s password=%s port=%s database=%s sslmode=disable", config.DBHost, config.DBUser, config.DBPassword, config.DBPort, config.DBName)
	log.Println("dbUri: " + dbURI)

	conn, err := sql.Open(config.DBDriver, dbURI)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	//url := fmt.Sprintf("%v://%v:%v@%v:%v/%v",
	//	config.DBDriver,
	//	config.DBUser,
	//	config.DBPassword,
	//	config.DBHost,
	//	config.DBPort,
	//	config.DBName)
	//runDbMigration(config.MigrationURL, url)

	http.HandleFunc("/", handler(conn))
	log.Fatal(http.ListenAndServe(":8080", nil))

	//store := db.NewStore(conn)
	//server, err := api.NewServer(config, store)
	//if err != nil {
	//	log.Fatal("cannot create server: ", err)
	//}
	//
	//err = server.Start(config.ServerAddress)
	//if err != nil {
	//	log.Fatal("cannot start server: ", err)
	//}
}

func handler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := db.Ping()
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("error while trying to connect to database"))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ping success to database"))
		log.Println("ping success")
	}
}

func runDbMigration(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Println(dbSource)
		log.Println("migUrl: " + migrationURL)
		log.Fatal("cannot create migration: ", err)
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("cannot run migration up: ", err)
	}

	log.Println("migration up success")
}
