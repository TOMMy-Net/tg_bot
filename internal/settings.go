package internal

import (
	"encoding/json"
	"log"
	"os"
	"tg_bot/internal/db"
	"tg_bot/internal/tools"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Settings struct {
	Bot      *tgbotapi.BotAPI
	Logger   *log.Logger
	MsgTexts Texts
	Cache    *tools.Cache
	Storage  *db.Storage
}

type Texts struct {
	HelloText    string `json:"hello_text"`
	Information  string `json:"company_inforamation"`
	CategoryText string `json:"category_text"`
}

func NewSettings() *Settings {
	s := new(Settings)
	s.Logger = newLogger()
	s.MsgTexts = loadTexts()
	s.Cache = tools.NewCache()
	return s
}

func newLogger() *log.Logger {
	if os.Getenv("LOG_PATH") == "" {
		os.Setenv("LOG_PATH", "/")
	}
	file, err := os.Open(os.Getenv("LOG_PATH"))
	if err != nil {
		log.Fatal(err)
	}
	l := log.New(file, "\n", log.Ldate|log.Ltime)
	return l
}

func loadTexts() Texts {
	file, err := os.Open("texts.json")
	if err != nil {
		log.Fatal(err)
	}

	t := Texts{}
	err = json.NewDecoder(file).Decode(&t)
	if err != nil {
		log.Fatal(err)
	}
	return t
}
