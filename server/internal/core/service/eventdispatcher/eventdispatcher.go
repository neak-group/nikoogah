package eventdispatcher

import "github.com/neak-group/nikoogah/internal/core/service/eventbus"

type EventDispatcher interface {
	Dispatch(event eventbus.Event) error
}