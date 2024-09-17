package eventbus

import "context"

type Event interface {
	GetEventType() string
}

type EventHandler interface {
	Handle(context.Context, Event) error
	GetEventType() string
}

type EventBus interface {
	Register(eventType string, handler EventHandler)
	Publish(event Event) error
}
