package internal

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	CategoryButtonText    = "Меню категорий"
	InformationButtonText = "Информация о компании"
	HelpButtonText        = "Поддержка"
)

const (
	HelpText = "Опишите ваш вопрос, и мы передадим его в техническую поддержку. Вам ответят в ближайшее время."
)

var (
	MenuButtons = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton(CategoryButtonText)),
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton(InformationButtonText)),
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton(HelpButtonText)),
	)
)

func (s *Settings) Commands(update *tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.From.ID, "")
	switch update.Message.Command() {
	case "start":
		msg.Text = s.MsgTexts.HelloText
		msg.ReplyMarkup = MenuButtons
	case "about":
		msg.Text = s.MsgTexts.Information
		msg.ReplyMarkup = MenuButtons
	}

	s.Send(msg)
}

func (s *Settings) Messages(update *tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.From.ID, "")
	switch update.Message.Text {
	case InformationButtonText:
		msg.Text = s.MsgTexts.Information
		s.Send(msg)
	case CategoryButtonText:
		msg.Text = s.MsgTexts.CategoryText
		msg.ReplyMarkup = CategoryButtons
		s.Send(msg)
	case CancelApplicationText:
		delete(s.ApplicationCache, UserId(update.Message.From.ID))
		msg.Text = s.MsgTexts.HelloText
		msg.ReplyMarkup = MenuButtons
		s.Send(msg)
	case HelpButtonText:
		s.HelpHandlerMessage(update)
	default:
		msg.Text = "Не знаю такой команды"
		s.Send(msg)
	}
	
}


func (s *Settings) HelpHandlerMessage(update *tgbotapi.Update)  {
	msg := tgbotapi.NewMessage(update.Message.From.ID, HelpText)
	s.Send(msg)
}


func (s *Settings) Send(msg tgbotapi.Chattable) error {
	if _, err := s.Bot.Send(msg); err != nil {
		s.Logger.Println(err)
		return err
	}
	return nil
}

func (s *Settings) Request(msg tgbotapi.Chattable) error {
	if _, err := s.Bot.Request(msg); err != nil {
		s.Logger.Println(err)
		return err
	}
	return nil
}


