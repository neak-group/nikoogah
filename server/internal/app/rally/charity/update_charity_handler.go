package charity

import (
	"context"

	"github.com/neak-group/nikoogah/internal/app"
	"github.com/neak-group/nikoogah/internal/app/rally/charity/entity"
	"github.com/neak-group/nikoogah/internal/core/domain/events"
	"github.com/neak-group/nikoogah/internal/core/service/eventbus"
)

type UpdateCharityHandlerFxParams struct {
	app.HandlerParams

	Repo CharityRepository
}

type UpdateCharityHandler struct {
	app.BaseHandler
	repo CharityRepository
}

func init() {
	app.RegisterHandlerProvider(ProvideCharityHandler)
}

func ProvideCharityHandler(params UpdateCharityHandlerFxParams) *UpdateCharityHandler {
	return &UpdateCharityHandler{
		repo: params.Repo,
		BaseHandler: app.BaseHandler{
			Logger: params.Logger,
		},
	}
}

func (h UpdateCharityHandler) Handle(ctx context.Context, e eventbus.Event) error {
	charityEvent, ok := e.(events.CharityUpdatedEvent)
	if !ok {
		h.Logger.Error("invalid event type")
	}

	charity := entity.UpdateCharity(charityEvent.ID, charityEvent.Name, charityEvent.Phone, charityEvent.Email)

	err := h.repo.SaveCharity(charity)
	if err != nil {
		return err
	}

	return nil
}
