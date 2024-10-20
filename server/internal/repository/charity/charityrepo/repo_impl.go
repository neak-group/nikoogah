package charityrepo

import (
	"github.com/neak-group/nikoogah/internal/app/charity/charity/repository"
	"github.com/neak-group/nikoogah/internal/infra/mongofx"
	"github.com/neak-group/nikoogah/internal/repository/charity/charityrepo/mongo"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type MongoRepositoryImplParams struct {
	fx.In

	MongoClient mongofx.MongoDBConn
	Logger      *zap.Logger
}

func ProvideMongoRepositoryImpl(params MongoRepositoryImplParams) repository.CharityRepository {
	return &mongo.CharityMongoRepository{
		MongoClient: params.MongoClient,
		Logger:      params.Logger,

		CharityDatabase:     "charity_database",
		CharitiesCollection: "charity",
	}
}
