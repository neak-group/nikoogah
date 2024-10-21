package app

import (
	"github.com/neak-group/nikoogah/internal/app/charity"
	"github.com/neak-group/nikoogah/internal/app/rally"
	"github.com/neak-group/nikoogah/internal/app/volunteer"
	"github.com/neak-group/nikoogah/internal/core/service/eventbus"
	"go.uber.org/fx"
)

func GetHandlerProviders() []interface{} {
	domainHandlerProviders := make([]interface{}, 0)

	domainHandlerProviders = append(domainHandlerProviders, rally.GetHandlerProviders()...)
	domainHandlerProviders = append(domainHandlerProviders, charity.GetHandlerProviders()...)
	domainHandlerProviders = append(domainHandlerProviders, volunteer.GetHandlerProviders()...)

	domainHandlerProvidersAnnotated := make([]interface{}, 0)

	for _, hp := range domainHandlerProviders {
		domainHandlerProvidersAnnotated = append(domainHandlerProvidersAnnotated, fx.Annotate(
			hp,
			fx.As(new(eventbus.EventHandler)),
			fx.ResultTags(`group:"event-handlers"`),
		))
	}

	return domainHandlerProvidersAnnotated
}
