package db

import "github.com/jmoiron/sqlx"

type Database struct {
	*sqlx.DB
}

func Connect() (*Database, error) {

	conn, err := sqlx.Open("sqlite3", "./internal/db/data.db")
	if err != nil {
		return &Database{}, err
	}

	
	return &Database{conn}, nil
}
