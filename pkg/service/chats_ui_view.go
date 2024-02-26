package service

import (
	"github.com/nicksedov/gptbot/pkg/model"
	"github.com/nicksedov/gptbot/pkg/view"
)

func GetChatsTabView() (*view.ChatsTabView, error) {
	chats, dbErr := model.GetAll[model.TelegramChat]()
	if dbErr != nil {
		return nil, dbErr
	}
	chatViews := make([]view.ChatView, 0, len(*chats))
	for _, chat := range *chats {
		chatViews = append(chatViews, view.ChatView{ID: chat.ID, ChatID: chat.ChatID, ChatName: chat.ChatName})
	}
	return &view.ChatsTabView{ChatViews: chatViews}, nil
}