package eventdispatcher

import (
	"github.com/neak-group/nikoogah/internal/core/service/eventbus"
	"github.com/neak-group/nikoogah/internal/core/service/eventdispatcher"
	"go.uber.org/zap"
)

type EventDispatcherParams struct {
	EventBus eventbus.EventBus
	Logger   *zap.Logger
}

func ProvideEventDispatcher(EventBus eventbus.EventBus, Logger *zap.Logger) eventdispatcher.EventDispatcher {
	return &eventDispatcherImpl{
		eventbus: EventBus,
		logger:   Logger,
	}
}

type eventDispatcherImpl struct {
	eventbus eventbus.EventBus
	logger   *zap.Logger
}

func (ed eventDispatcherImpl) Dispatch(event eventbus.Event) error {

	//TODO: Dead Letter Pattern
	return ed.eventbus.Publish(event)
}

func (ed eventDispatcherImpl) DispatchBatch(events []eventbus.Event) error {

	//TODO: Dead Letter Pattern
	for _, e := range events {
		if err := ed.eventbus.Publish(e); err != nil {
			ed.logger.Error(err.Error())
		}
	}

	return nil
}
