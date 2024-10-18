package mongofx

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type MongoConn struct {
	conn *mongo.Client
}

func (cfg *MongoConfig) toMongoConnStr() string {
	var connectionString strings.Builder

	connectionString.WriteString("mongodb://")

	// Add username and password if provided
	if cfg.Username != "" && cfg.Password != "" {
		connectionString.WriteString(url.QueryEscape(cfg.Username))
		connectionString.WriteByte(':')
		connectionString.WriteString(url.QueryEscape(cfg.Password))
		connectionString.WriteByte('@')
	}

	// Add host and port
	connectionString.WriteString(cfg.Host)
	if cfg.Port != "" {
		connectionString.WriteByte(':')
		connectionString.WriteString(cfg.Port)
	}

	authSource := viper.GetString("mongo_auth_source")

	if authSource != "" {
		connectionString.WriteByte('/')
		connectionString.WriteByte('?') // Add auth source
		connectionString.WriteString("authSource=")
		connectionString.WriteString(authSource)
	}

	return connectionString.String()
}

type MongoParams struct {
	fx.In

	Config MongoConfig
	Logger *zap.Logger
}

func NewMongoConn(params MongoParams) (*mongo.Client, error) {
	cfg := params.Config
	connStr := cfg.toMongoConnStr()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connStr))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("could not connect to mongodb, %w", err)
	}

	return client, nil
}

func (c MongoConn) GetConn() (*mongo.Client, error) {
	if c.conn == nil {
		return nil, fmt.Errorf("no collection database connection found")
	}
	return c.conn, nil
}
