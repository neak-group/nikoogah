package mongofx

import (
	"context"
	"fmt"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type MongoConfig struct {
	Username string
	Password string
	Host     string
	Port     string
}

var Username = "col_db_user_name"
var Password = "col_db_pass"
var Host = "col_db_host"
var Port = "col_db_port"

type MongoDBConn interface {
	GetDB(ctx context.Context, dbName string) (*mongo.Database, error)
}

type dbConn struct {
	MongoClient *mongo.Client
}

func (db *dbConn) GetDB(ctx context.Context, dbName string) (*mongo.Database, error) {
	if db.MongoClient == nil {
		return nil, fmt.Errorf("no database connection found")
	}
	return db.MongoClient.Database(dbName), nil
}

func ProvideMongoDBConfig(logger *zap.Logger) MongoConfig {
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

	return cfg
}

func NewMongoDBConn(client *mongo.Client, logger *zap.Logger) (MongoDBConn, error) {
	return &dbConn{
		MongoClient: client,
	}, nil
}
