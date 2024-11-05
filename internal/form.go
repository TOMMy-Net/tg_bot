package internal

import (
	"fmt"
	"tg_bot/internal/tools"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	ConfirmApplication    = "Подтвердить"
	CancelApplicationText = "Отменить ввод"
	EditApplication       = "Изменить"
	EditAll               = "Полностью"
	EditContacts          = "Имя и контакты"
)
var (
	QuesEdit = "Хотите ли вы редактировать только имя и контакты или полностью заявку?"
)

var (
	ApplicationButtons = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton(ConfirmApplication)),
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton(EditApplication)),
	)

	EditButtons = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton(EditAll)),
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton(EditContacts)),
	)
)

func (s *Settings) ApplicationEnter(update *tgbotapi.Update) {

	msg := tgbotapi.NewMessage(update.Message.From.ID, "")
	switch update.Message.Text {
	case CancelApplicationText:
		delete(s.ApplicationCache, UserId(update.Message.From.ID))
		msg.Text = s.MsgTexts.HelloText
		msg.ReplyMarkup = MenuButtons
		s.Send(msg)
	case ConfirmApplication:
		application := s.ApplicationCache[UserId(update.Message.From.ID)]
		err := tools.Validate(application)
		if err != nil {
			msg.Text = "Ошибка ввода, попробуйте заново."
			s.Send(msg)
			return
		}
		id, err := s.Storage.CreateApplication(application)
		if err != nil {
			msg.Text = "Ошибка при сохранении заявки, попробуйте заново."
			s.Send(msg)
		} else {
			delete(s.ApplicationCache, UserId(update.Message.From.ID))
			msg.Text = fmt.Sprintf(`Спасибо! Ваша заявка успешно отправлена. Мы свяжемся с вами в ближайшее время. Ваш идентификационный номер заявки: %d.`, id)
			msg.ReplyMarkup = MenuButtons
			s.Send(msg)
		}
	case EditApplication:
		msg.Text = QuesEdit
		msg.ReplyMarkup = EditButtons
		s.Send(msg)
	case EditAll:
		msg1 := tgbotapi.NewMessage(update.Message.From.ID, "Изменение анкеты")
		msg1.ReplyMarkup = CancelApplicationButton
		s.Send(msg1)
		msg.Text = s.MsgTexts.CategoryText
		msg.ReplyMarkup = CategoryButtons
		s.Send(msg)
	case EditContacts:
		c := s.ApplicationCache[UserId(update.Message.From.ID)]
		c.Name = ""
		c.PhoneNumber = ""
		s.ApplicationCache[UserId(update.Message.From.ID)] = c
		msg.ReplyMarkup = CancelApplicationButton
		msg.Text = "Пожалуйста, укажите ваше имя."
		s.Send(msg)
	default:
		c := s.ApplicationCache[UserId(update.Message.From.ID)]
		if c.Name == "" {
			c.Name = update.Message.Text
			s.ApplicationCache[UserId(update.Message.From.ID)] = c
			msg.Text = "Пожалуйста, укажите ваш номер телефона."
			s.Send(msg)
		} else if c.PhoneNumber == "" {
			c.PhoneNumber = update.Message.Text
			s.ApplicationCache[UserId(update.Message.From.ID)] = c
			msg.Text = fmt.Sprintf(`Ваша заявка готова. Пожалуйста, проверьте информацию:
				Категория: %s
				Имя: %s
				Контакты: %s`, c.Category, c.Name, c.PhoneNumber)
			msg.ReplyMarkup = ApplicationButtons
			s.Send(msg)

		}
	}

}

func (s *Settings) CheckApplication(id int64) (b bool) {
	_, b = s.ApplicationCache[UserId(id)]
	return
}
