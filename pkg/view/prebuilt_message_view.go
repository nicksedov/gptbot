package view

type PrebuiltMessageView struct {
	ID          uint      `json:"id"`
	EventID     uint      `json:"singleEventId"`
	Status      string    `json:"status"`
	Message     string    `json:"message"`
}