package eventbus

import (
	"sync"

	"go.uber.org/fx"
)

var Module = fx.Module("event-bus",
	fx.Provide(
		fx.Annotate(
			ProvideEventBus,
			fx.ParamTags(`group:"eventHandlers`),
		),
	),
)

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

type eventBusImpl struct {
	handlers map[string][]EventHandler
	mutex    sync.RWMutex
}

func ProvideEventBus(handlers []EventHandler) EventBus {
	eventBus := &eventBusImpl{
		handlers: make(map[string][]EventHandler),
	}

	for _, h := range handlers {
		eventBus.Register(h.GetEventType(), h)
	}

	return eventBus
}

func (bus *eventBusImpl) Register(eventType string, handler EventHandler) {
	bus.mutex.Lock()
	defer bus.mutex.Unlock()
	bus.handlers[eventType] = append(bus.handlers[eventType], handler)
}

func (bus *eventBusImpl) Publish(event Event) error {
	bus.mutex.RLock()
	defer bus.mutex.RUnlock()
	if handlers, found := bus.handlers[event.GetEventType()]; found {
		for _, handler := range handlers {
			go handler.Handle(event)
		}
	}
	return nil
}
