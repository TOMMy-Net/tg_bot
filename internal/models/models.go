package models

type Application struct {
	Id          int    `db:"id"`
	UserId      int    `validate:"required,min=1" db:"user_id"`
	Name        string `validate:"required,min=1" db:"name"`
	PhoneNumber string `validate:"required,min=1" db:"phone_number"`
	Category    string `validate:"required,min=1" db:"category"`
}

type Support struct {
	Id          string `db:"id"` // uuid
	UserId      int    `validate:"required,min=1" db:"user_id"`
	Name        string `validate:"required,min=1" db:"name"`
	PhoneNumber string `validate:"required,min=1" db:"phone_number"`
	Problem     string `validate:"required,min=1" db:"problem"`
}
