package database

import (
	"context"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DB struct {
	DB *sqlx.DB
}

func ConnectDB() (*DB, error) {
	db, err := sqlx.Connect("postgres", "host=db port=5432 user=postgres password=postgres dbname=movie_wishlist sslmode=disable")
	if err != nil {
		return &DB{DB: nil}, err
	}

	return &DB{DB: db}, nil
}

func (db *DB) Ping(ctx context.Context) error {
	return db.DB.PingContext(ctx)
}
