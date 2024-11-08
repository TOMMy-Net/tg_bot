package internal

import (
	"tg_bot/internal/models"
	"tg_bot/internal/tools"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
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

// all messages from users
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
		s.Cache.DeleteApplication(update.Message.From.ID)
		msg.Text = s.MsgTexts.HelloText
		msg.ReplyMarkup = MenuButtons
		s.Send(msg)
	case HelpButtonText:
		s.HelpHandlerMessage(update)
	default:
		cache := s.Cache.ReadSupport(update.Message.From.ID)
		if cache != (models.Support{}) {
			cache.Problem = update.Message.Text
			if err := tools.Validate(cache); err != nil {
				s.Send(tgbotapi.NewMessage(update.Message.From.ID, "Ошибка сервера"))
				return
			}
			cache.Id = uuid.NewString()
			if err := s.Storage.CreateSupport(cache); err != nil {
				s.Send(tgbotapi.NewMessage(update.Message.From.ID, "Ошибка сервера"))
				return
			}
			s.Cache.DeleteSupport(update.Message.From.ID)
			s.Send(tgbotapi.NewMessage(update.Message.From.ID, "Ваш запрос был отправлен в техническую поддержку. Мы свяжемся с вами в ближайшее время."))
			return
		}
		msg.Text = "Не знаю такой команды"
		s.Send(msg)
	}

}

func (s *Settings) HelpHandlerMessage(update *tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.From.ID, HelpText)

	if err := s.Send(msg); err == nil {
		s.Cache.StoreSupport(update.Message.From.ID, models.Support{
			UserId: int(update.Message.From.ID),
		})
	}
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
