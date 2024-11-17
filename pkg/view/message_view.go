package view

type InstantMessageFormView struct {
	MessageText    string `json:"messageText"`
	TelegramChatID uint   `json:"telegramChatId"`
}