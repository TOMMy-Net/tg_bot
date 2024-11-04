package internal

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	Categorys = map[string]string{
		"category_1": "Разработка корпоративных приложений",
		"category_2": "Системы контроля за сотрудниками",
		"category_3": "Системы лояльности и бонусов",
		"category_4": "Разработка сайтов",
		"category_5": "Приложения для карт и бонусов",
		"category_6": "Другое",
	}
)
var (
	CategoryButtons = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(Categorys["category_1"], "category_1"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(Categorys["category_2"], "category_2"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(Categorys["category_3"], "category_3"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(Categorys["category_4"], "category_4"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(Categorys["category_5"], "category_5"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(Categorys["category_6"], "category_6"),
		),
	)

	CancelApplicationButton = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton(CancelApplicationText)),
	)
)

func (s *Settings) CallbackQuery(update *tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.CallbackQuery.From.ID, "")
	text := strings.Split(update.CallbackQuery.Data, "_")
	switch text[0] {
	case "category":

		del := tgbotapi.NewDeleteMessage(update.CallbackQuery.From.ID, update.CallbackQuery.Message.MessageID)
		errR := s.Request(del)

		msg.Text = "Пожалуйста, укажите ваше имя."
		msg.ReplyMarkup = CancelApplicationButton
		errS := s.Send(msg)

		if errR == nil && errS == nil {
			s.ApplicationCache[UserId(update.CallbackQuery.From.ID)] = Application{
				UserId:   int(update.CallbackQuery.From.ID),
				Category: Categorys[update.CallbackQuery.Data],
			}
		}
	}
}
