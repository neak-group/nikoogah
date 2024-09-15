package eventbus

type Event interface {
	GetEventType() string
}

type EventHandler interface {
	Handle(Event) error
	GetEventType() string
}

type EventBus interface {
	Register(eventType string, handler EventHandler)
	Publish(event Event) error
}
