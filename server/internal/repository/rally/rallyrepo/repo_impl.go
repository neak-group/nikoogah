package rallyrepo

import (
	"github.com/neak-group/nikoogah/internal/app/rally/rally/repository"
	"github.com/neak-group/nikoogah/internal/infra/mongofx"
	"github.com/neak-group/nikoogah/internal/repository/rally/rallyrepo/mongo"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type MongoRepositoryImplParams struct {
	fx.In

	MongoClient mongofx.MongoDBConn
	Logger      *zap.Logger
}

func ProvideMongoRepositoryImpl(params MongoRepositoryImplParams) repository.RallyRepository {
	return &mongo.RallyMongoRepository{
		MongoClient: params.MongoClient,
		Logger:      params.Logger,

		RallyDatabase:   "rally_database",
		RallyCollection: "rly_rally",
	}
}
