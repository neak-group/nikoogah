package volunteerrepo

import (
	"github.com/neak-group/nikoogah/internal/app/rally/volunteer/repository"
	"github.com/neak-group/nikoogah/internal/infra/mongofx"
	"github.com/neak-group/nikoogah/internal/repository/rally/volunteerrepo/mongo"
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

		RallyDatabase:        "rally_database",
		VolunteersCollection: "rly_volunteer",
	}
}
