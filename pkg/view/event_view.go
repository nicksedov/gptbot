package view

type EventsTabView struct {
	EventViews   []EventView        `json:"events"`
	Prompts      []DropdownItem     `json:"prompts"`
	PromptParams []PromptParamView  `json:"promptParams"`
	Chats        []DropdownItem     `json:"chats"`
}

type EventView struct {
	ID             uint
	Date           string
	Time           string
	TZOffset       int
	PromptTitle    string
	Prompt         string
	TelegramChatID uint
}

type PromptParamView struct {
	ID             uint
	PromptID       uint
	Title          string
}

type DropdownItem struct {
	ID    uint
	Value string
}
