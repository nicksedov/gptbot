package scheduler

import "gptbot/pkg/model"

type EventHandler interface {
	handle(t *model.SingleEvent) error

	onError(err error)
}
