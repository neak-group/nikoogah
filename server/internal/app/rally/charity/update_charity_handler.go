package charity

import (
	"context"

	"github.com/neak-group/nikoogah/internal/app/rally/charity/entity"
	"github.com/neak-group/nikoogah/internal/app/rally/charity/repository"
	"github.com/neak-group/nikoogah/internal/core/domain/base"
	"github.com/neak-group/nikoogah/internal/core/domain/events"
	"github.com/neak-group/nikoogah/internal/core/service/eventbus"
)

type UpdateCharityHandlerFxParams struct {
	base.HandlerParams

	Repo repository.CharityRepository
}

type UpdateCharityHandler struct {
	base.BaseHandler
	repo repository.CharityRepository
}

func ProvideCharityHandler(params UpdateCharityHandlerFxParams) eventbus.EventHandler {
	return &UpdateCharityHandler{
		repo: params.Repo,
		BaseHandler: base.BaseHandler{
			Logger: params.Logger,
		},
	}
}

func (h UpdateCharityHandler) GetEventTypes() []string {
	return []string{events.CharityUpdatedEvent{}.GetEventType()}
}

func (h UpdateCharityHandler) Handle(ctx context.Context, e eventbus.Event) error {
	charityEvent, ok := e.(events.CharityUpdatedEvent)
	if !ok {
		h.Logger.Error("invalid event type")
	}

	h.Logger.Info("charity handler called")

	charity := entity.UpdateCharity(charityEvent.ID, charityEvent.Name, charityEvent.Phone, charityEvent.Email)

	err := h.repo.SaveCharity(ctx, charity)
	if err != nil {
		return err
	}

	return nil
}
