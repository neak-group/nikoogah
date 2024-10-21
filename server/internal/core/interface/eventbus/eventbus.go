package eventbus

import (
	"context"
	"sync"

	"github.com/neak-group/nikoogah/internal/core/service/eventbus"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type eventBusImpl struct {
	handlers map[string][]eventbus.EventHandler
	mutex    sync.RWMutex

	logger *zap.Logger
}

type EventBusParams struct {
	fx.In

	Handlers []eventbus.EventHandler `group:"event-handlers"`
	Logger   *zap.Logger
}

func ProvideEventBus(params EventBusParams) eventbus.EventBus {
	eventBus := &eventBusImpl{
		handlers: make(map[string][]eventbus.EventHandler),
		logger:   params.Logger,
	}
	eventBus.logger.Info("registering event bus")
	eventBus.logger.Info("registering handlers", zap.Int("number of handlers", len(params.Handlers)))
	for _, h := range params.Handlers {
		eventBus.Register(h.GetEventTypes(), h)
	}

	return eventBus
}

func (bus *eventBusImpl) Register(eventTypes []string, handler eventbus.EventHandler) {
	bus.mutex.Lock()
	defer bus.mutex.Unlock()
	for _, eventType := range eventTypes {
		bus.logger.Info("handler registered", zap.String("type", eventType))
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
