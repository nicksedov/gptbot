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

type NewEventFormView struct {
	PromptID        uint       `json:"promptId"`
	Date            string     `json:"date"`
	Time            string     `json:"time"`
	TZOffset        int        `json:"tzOffset"` 
	ParamID0        string     `json:"param_id_0"`
	Param0          string     `json:"param_0"`
	ParamID1        string     `json:"param_id_1"`
	Param1          string     `json:"param_1"`
	ParamID2        string     `json:"param_id_2"`
	Param2          string     `json:"param_2"`
	TelegramChatID  uint       `json:"telegramChatId"`
}