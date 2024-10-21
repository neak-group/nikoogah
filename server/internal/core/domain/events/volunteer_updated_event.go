package events



import "github.com/neak-group/nikoogah/utils/uuid"

type VolunteerUpdatedEvent struct {
	ID   uuid.UUID
	Name string
}

func (VolunteerUpdatedEvent) GetEventType() string {
	return "event_volunteer_updated"
}