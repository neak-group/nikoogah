package mongofx

import (
	"context"
	"fmt"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type MongoConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

var Username = "COL_DB_USER_NAME"
var Password = "COL_DB_PASS"
var Host = "COL_DB_HOST"
var Port = "COL_DB_PORT"
var Database = "COL_DB_DATABASE"

type MongoDBConn interface {
	GetConn(ctx context.Context) (*mongo.Client, error)
}

type dbConn struct {
	MongoClient *mongo.Client
}

func (db *dbConn) GetConn(ctx context.Context) (*mongo.Client, error) {
	if db.MongoClient == nil {
		return nil, fmt.Errorf("no database connection found")
	}
	return db.MongoClient, nil
}

func ProvideMongoDBConfig(lc fx.Lifecycle, logger *zap.Logger) MongoConfig {
	cfg := MongoConfig{}

	cfg.Username = viper.GetString(Username)
	if cfg.Username == "" {
		logger.Panic("%v env must be specified: username", zap.String("env", Username))
	}

	cfg.Password = viper.GetString(Password)
	if cfg.Password == "" {
		logger.Panic("%v env must be specified: password", zap.String("env", Password))
	}

	cfg.Host = viper.GetString(Host)
	if cfg.Host == "" {
		logger.Panic("%v env must be specified: host", zap.String("env", Host))
	}

	cfg.Port = viper.GetString(Port)
	if cfg.Port == "" {
		logger.Panic("%v env must be specified: port", zap.String("env", Port))
	}

	cfg.Database = viper.GetString(Database)
	if cfg.Database == "" {
		logger.Panic("%v env must be specified: database", zap.String("env", Database))
	}

	return cfg
}

func NewMongoDBConn(client *mongo.Client, logger *zap.Logger) (MongoDBConn, error) {
	return &dbConn{
		MongoClient: client,
	}, nil
}
