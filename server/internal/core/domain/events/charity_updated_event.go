package events

import "github.com/neak-group/nikoogah/utils/uuid"

type CharityUpdatedEvent struct {
	ID            uuid.UUID
	Name          string
	Phone         string
	Email         string
	MaxRallyLimit int
}

func (CharityUpdatedEvent) GetEventType() string {
	return "event_charity_updated"
}
