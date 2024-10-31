package db

import "github.com/jmoiron/sqlx"

type Database struct {
	*sqlx.DB
}

func Connect() *Database {
	db := new(Database)

	return db
}
