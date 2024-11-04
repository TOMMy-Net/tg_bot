package internal

type Application struct {
	Id          string
	UserId      int    `validate:"required,min=1"`
	Name        string `validate:"required,min=1"`
	PhoneNumber string `validate:"required,min=1"`
	Category    string `validate:"required,min=1"`
}
