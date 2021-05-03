// pkg/db/db.go

package db

import (
	"log"

	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

type DB struct {
	Postgres *sqlx.DB
}

var dbInstance *DB

func Connect(connStr string) (*DB, error) {
	if dbInstance == nil {
		db, err := connect(connStr)
		if err != nil {
			return nil, err
		}

		dbInstance = &DB{
			Postgres: db,
		}
	}
	return dbInstance, nil
}

func GetPostgresDB() *sqlx.DB {
	return dbInstance.Postgres
}

func (d *DB) Close() error {
	return d.Postgres.Close()
}

func connect(connStr string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Println("DB Error", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Println("DB Ping Error", err)
		return nil, err
	}
	log.Println("DB connection successful")

	return db, nil
}
