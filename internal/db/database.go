package db

import (
	"tg_bot/internal/models"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sqlx.DB
}

func Connect() (*Storage, error) {

	conn, err := sqlx.Open("sqlite3", "./internal/db/data.db")
	if err != nil {
		return &Storage{}, err
	}

	if err := migrateShema(conn); err != nil {
		return &Storage{}, err
	}

	return &Storage{db: conn}, nil
}

func migrateShema(d *sqlx.DB) error {
	driver, err := sqlite3.WithInstance(d.DB, &sqlite3.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://internal/db/migrations",
		"sqlite3", driver)
	if err != nil {
		return err
	}

	m.Up()
	return nil

}

func (s *Storage) CreateApplication(a models.Application) (int64, error) {
	res, err := s.db.NamedExec("INSERT INTO application(user_id, name, phone_number, category) VALUES(:user_id, :name, :phone_number, :category)", a)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}
