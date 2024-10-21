package events

import "github.com/neak-group/nikoogah/utils/uuid"

type UserJoinedEvent struct {
	ID   uuid.UUID
	Name string
}

func (UserJoinedEvent) GetEventType() string {
	return "event_user_joined"
}
