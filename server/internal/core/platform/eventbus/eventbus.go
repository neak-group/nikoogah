package eventbus

import (
	"sync"

	"github.com/neak-group/nikoogah/internal/core/service/eventbus"
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

type eventBusImpl struct {
	handlers map[string][]eventbus.EventHandler
	mutex    sync.RWMutex
}

func ProvideEventBus(handlers []eventbus.EventHandler) eventbus.EventBus {
	eventBus := &eventBusImpl{
		handlers: make(map[string][]eventbus.EventHandler),
	}

	for _, h := range handlers {
		eventBus.Register(h.GetEventType(), h)
	}

	return eventBus
}

func (bus *eventBusImpl) Register(eventType string, handler eventbus.EventHandler) {
	bus.mutex.Lock()
	defer bus.mutex.Unlock()
	bus.handlers[eventType] = append(bus.handlers[eventType], handler)
}

func (bus *eventBusImpl) Publish(event eventbus.Event) error {
	bus.mutex.RLock()
	defer bus.mutex.RUnlock()
	if handlers, found := bus.handlers[event.GetEventType()]; found {
		for _, handler := range handlers {
			go handler.Handle(event)
		}
	}
	return nil
}
