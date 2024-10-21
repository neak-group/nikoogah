package volunteer

import (
	"context"

	"github.com/neak-group/nikoogah/internal/app/rally/volunteer/entity"
	"github.com/neak-group/nikoogah/internal/app/rally/volunteer/repository"
	"github.com/neak-group/nikoogah/internal/core/domain/base"
	"github.com/neak-group/nikoogah/internal/core/domain/events"
	"github.com/neak-group/nikoogah/internal/core/service/eventbus"
)

type UpdateVolunteerHandlerFxParams struct {
	base.HandlerParams

	Repo repository.VolunteerRepository
}

type UpdateVolunteerHandler struct {
	base.BaseHandler
	repo repository.VolunteerRepository
}

func ProvideVolunteerHandler(params UpdateVolunteerHandlerFxParams) eventbus.EventHandler {
	return &UpdateVolunteerHandler{
		repo: params.Repo,
		BaseHandler: base.BaseHandler{
			Logger: params.Logger,
		},
	}
}

func (h UpdateVolunteerHandler) GetEventTypes() []string {
	return []string{events.UserJoinedEvent{}.GetEventType()}
}

func (h UpdateVolunteerHandler) Handle(ctx context.Context, e eventbus.Event) error {
	userEvent, ok := e.(events.UserJoinedEvent)
	if !ok {
		h.Logger.Error("invalid event type")
	}

	volunteer, err := entity.UpdateVolunteer(userEvent.ID, userEvent.Name)
	if err != nil {
		return err
	}

	err = h.repo.UpdateVolunteer(ctx, volunteer)
	if err != nil {
		return err
	}

	return nil
}
