package keystorefx

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

const (
	KeyStoreHost = "key_store_host"
	KeyStorePort = "key_store_port"
	KeyStoreUser = "key_store_user"
	KeyStorePass = "key_store_pass"
)

type KeyStoreConfig struct {
	Address  string
	Port     uint16
	User     string
	Password string
}

type KeyStoreConn interface {
	KSClient(context.Context) (*redis.Client, error)
}

type ksConn struct {
	client *redis.Client
}

func (ks *ksConn) KSClient(ctx context.Context) (*redis.Client, error) {
	if ks.client == nil {
		return nil, fmt.Errorf("no database connection found")
	}
	return ks.client, nil
}

func ProvideKeyStoreConfig(lc fx.Lifecycle, logger *zap.Logger) KeyStoreConfig {
	cfg := KeyStoreConfig{}
	cfg.Address = viper.GetString(KeyStoreHost)
	if cfg.Address == "" {
		logger.Panic("failed to read db env", zap.String("env", KeyStoreHost))
	}

	cfg.Port = viper.GetUint16(KeyStorePort)
	if cfg.Port == 0 {
		logger.Panic("failed to read db env", zap.String("env", KeyStorePort))
	}

	cfg.User = viper.GetString(KeyStoreUser)
	if cfg.User == "" {
		logger.Panic("failed to read db env", zap.String("env", KeyStoreUser))
	}

	cfg.Password = viper.GetString(KeyStorePass)
	if cfg.Password == "" {
		logger.Panic("failed to read db env", zap.String("env", KeyStorePass))
	}

	return cfg
}

func ProvideKeyStoreConn(client *redis.Client, logger *zap.Logger) KeyStoreConn {
	logger.Info("keystore connection established")

	return &ksConn{
		client: client,
	}
}
