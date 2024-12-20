package charityrepo

import (
	"github.com/neak-group/nikoogah/internal/app/rally/charity/repository"
	"github.com/neak-group/nikoogah/internal/infra/mongofx"
	"github.com/neak-group/nikoogah/internal/repository/rally/charityrepo/mongo"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type MongoRepositoryImplParams struct {
	fx.In

	MongoClient mongofx.MongoDBConn
	Logger      *zap.Logger
}

func ProvideMongoRepositoryImpl(params MongoRepositoryImplParams) repository.CharityRepository {
	return &mongo.RlyCharityMongoRepository{
		MongoClient: params.MongoClient,
		Logger:      params.Logger,

		RallyDatabase:       "rally_database",
		CharitiesCollection: "rly_charity",
	}
}
