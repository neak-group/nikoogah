package eventbus

import (
	"context"
	"sync"

	"github.com/neak-group/nikoogah/internal/core/service/eventbus"
)

type eventBusImpl struct {
	handlers map[string][]eventbus.EventHandler
	mutex    sync.RWMutex
}

func ProvideEventBus(handlers []eventbus.EventHandler) eventbus.EventBus {
	eventBus := &eventBusImpl{
		handlers: make(map[string][]eventbus.EventHandler),
	}

	for _, h := range handlers {
		eventBus.Register(h.GetEventTypes(), h)
	}

	return eventBus
}

func (bus *eventBusImpl) Register(eventTypes []string, handler eventbus.EventHandler) {
	bus.mutex.Lock()
	defer bus.mutex.Unlock()
	for _, eventType := range eventTypes {
		
		bus.handlers[eventType] = append(bus.handlers[eventType], handler)
	}
}

func (bus *eventBusImpl) Publish(event eventbus.Event) error {
	bus.mutex.RLock()
	defer bus.mutex.RUnlock()
	if handlers, found := bus.handlers[event.GetEventType()]; found {
		for _, handler := range handlers {
			go handler.Handle(context.TODO(), event)
		}
	}
	return nil
}
