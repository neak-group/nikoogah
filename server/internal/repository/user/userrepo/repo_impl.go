package userrepo

import (
	userRepo "github.com/neak-group/nikoogah/internal/app/user/repository"
	"github.com/neak-group/nikoogah/internal/infra/mongofx"
	"github.com/neak-group/nikoogah/internal/repository"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type MongoRepositoryImplParams struct {
	fx.In

	MongoClient mongofx.MongoDBConn
	Logger      *zap.Logger
}

func ProvideMongoRepositoryImpl(params MongoRepositoryImplParams) userRepo.UserRepository {
	return &userMongoRepository{
		mongoClient: params.MongoClient,
		logger:      params.Logger,

		usersCollection: "base_user",
	}
}

func init() {
	repository.RegisterRepositoryProvider(ProvideMongoRepositoryImpl)
}
