package services

import (
	"context"

	"github.com/neak-group/nikoogah/internal/app/rally/rally/dto"
	rallyRepo "github.com/neak-group/nikoogah/internal/app/rally/rally/repository"
	volunteerEntity "github.com/neak-group/nikoogah/internal/app/rally/volunteer/entity"
	volunteerRepo "github.com/neak-group/nikoogah/internal/app/rally/volunteer/repository"
	"github.com/neak-group/nikoogah/utils/uuid"
	"go.uber.org/fx"
)

type RallyParticipationQueryService interface {
	GetRallyHumanParticipation(ctx context.Context, rallyID uuid.UUID) ([]*dto.HumanParticipationDTO, error)
}

type RallyParticipationQueryServiceImpl struct {
	volunteerRepo volunteerRepo.VolunteerRepository
	rallyRepo     rallyRepo.RallyRepository
}

type RallyParticipationQueryServiceParams struct {
	fx.In

	RallyRepo     rallyRepo.RallyRepository
	VolunteerRepo volunteerRepo.VolunteerRepository
}

func NewRallyParticipationQueryService(params RallyParticipationQueryServiceParams) RallyParticipationQueryService {
	return &RallyParticipationQueryServiceImpl{
		volunteerRepo: params.VolunteerRepo,
		rallyRepo:     params.RallyRepo,
	}
}

func (qs *RallyParticipationQueryServiceImpl) GetRallyHumanParticipation(ctx context.Context, rallyID uuid.UUID) ([]*dto.HumanParticipationDTO, error) {
	rally, err := qs.rallyRepo.FetchRally(ctx, rallyID)
	if err != nil {
		return nil, err
	}

	participation := make([]*dto.HumanParticipationDTO, 0)

	participantIDs := make(uuid.UUIDs, 0)

	for _, participant := range rally.HumanParticipations {
		participantIDs = append(participantIDs, participant.VolunteerID)
	}

	participantData, err := qs.volunteerRepo.FetchVolunteersByBatchID(ctx, participantIDs)
	if err != nil {
		return nil, err
	}

	participants := make(map[uuid.UUID]*volunteerEntity.Volunteer)
	for _, p := range participantData {
		participants[p.UserID] = p
	}

	for _, hp := range rally.HumanParticipations {
		participant := participants[hp.VolunteerID]
		participation = append(participation, &dto.HumanParticipationDTO{
			VolunteerID:         hp.VolunteerID,
			VolunteerName:       participant.FullName,
			VolunteerReputation: participant.Reputation,
			Phone:               hp.Phone,
			ResumeFile:          hp.ResumeFile,
			Status:              string(hp.Status),
		})
	}

	return participation, nil
}
