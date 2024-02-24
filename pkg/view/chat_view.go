package view

type ChatsTabView struct {
	ChatViews   []ChatView        `json:"chats"`
} 

type ChatView struct {
	ID        uint    `json:"id"`
	ChatID    uint    `json:"chatId"`
	ChatName  string  `json:"chatName"`
}