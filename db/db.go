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

func Get(connStr string) (*DB, error) {
	db, err := get(connStr)
	if err != nil {
		return nil, err
	}

	return &DB{
		Postgres: db,
	}, nil
}

func (d *DB) Close() error {
	return d.Postgres.Close()
}

func get(connStr string) (*sqlx.DB, error) {
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
