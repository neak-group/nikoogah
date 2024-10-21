package volunteerrepo

import (
	"github.com/neak-group/nikoogah/internal/app/volunteer/volunteer/repository"
	"github.com/neak-group/nikoogah/internal/infra/mongofx"
	"github.com/neak-group/nikoogah/internal/repository/volunteer/volunteerrepo/mongo"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type MongoRepositoryImplParams struct {
	fx.In

	MongoClient mongofx.MongoDBConn
	Logger      *zap.Logger
}

func ProvideMongoRepositoryImpl(params MongoRepositoryImplParams) repository.VolunteerRepository {
	return &mongo.VolunteerMongoRepository{
		MongoClient: params.MongoClient,
		Logger:      params.Logger,

		VolunteerDatabase:    "volunteer_database",
		VolunteersCollection: "vol_volunteer",
	}
}
