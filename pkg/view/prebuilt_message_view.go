package view

type PrebuiltMessageView struct {
	EventID     uint      `json:"singleEventId"`
	Status      string    `json:"status"`
	Message     string    `json:"message"`
}