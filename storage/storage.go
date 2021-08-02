package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once
)

func NewPostgresDB() {
	once.Do(func() {

		var err error

		// db, err = sql.Open("postgres", "postgres://postgres:admin@localhost:5432/apicf?sslmode=disable")
		db, err = sql.Open("postgres", "postgres://benji:benji@localhost:5432/base_test?sslmode=disable")
		if err != nil {
			log.Fatalf("Can´t open db: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("Can´t do ping: %v", err)
		}
		fmt.Println("Conectado a postgres")
	})
}

// Pool retorna una unica instancia de DB
func Pool() *sql.DB {
	return db
}
