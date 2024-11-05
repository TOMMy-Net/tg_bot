package models

type Application struct {
	Id          string `db:"id"`
	UserId      int    `validate:"required,min=1" db:"user_id"`
	Name        string `validate:"required,min=1" db:"name"`
	PhoneNumber string `validate:"required,min=1" db:"phone_number"`
	Category    string `validate:"required,min=1" db:"category"`
}
