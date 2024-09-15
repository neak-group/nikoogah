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

func ProvideEventDispatcher(params EventDispatcherParams) eventdispatcher.EventDispatcher {
	return &EventDispatcherImpl{
		eventbus: params.EventBus,
		logger:   params.Logger,
	}
}

type EventDispatcherImpl struct {
	eventbus eventbus.EventBus
	logger   *zap.Logger
}

func (EventDispatcherImpl) Dispatch(event eventbus.Event) error {
	return nil
}
