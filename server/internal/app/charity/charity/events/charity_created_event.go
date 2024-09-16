package events

type CharityAddedEvent struct{
	Name string
	Phone string
	Email string
}

func (CharityAddedEvent) GetEventType() string{
	return "event_charity_added"
}