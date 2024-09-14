package dbfx

import (
	"context"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/spf13/viper"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type DBConfig struct {
	Host              string
	Port              uint16
	Database          string
	User              string
	Password          string
	LogLevel          uint16
	MaxOpenConnection uint16
	MaxIdleConnection uint16
}

type DBConn interface {
	GormDB(context.Context) (*gorm.DB, error)
}

type dbConn struct {
	GormConnectionPool *gorm.DB
}

func (db *dbConn) GormDB(ctx context.Context) (*gorm.DB, error) {
	if db.GormConnectionPool == nil {
		return nil, fmt.Errorf("no database connection found")
	}
	return db.GormConnectionPool.WithContext(ctx), nil
}

const (
	DbHost          = "db_host"
	DbDatabase      = "db_database"
	DbPort          = "db_port"
	DbUser          = "db_user"
	DbPass          = "db_pass"
	GormLogMODE     = "gorm_log_mode"
	LogGormEnv      = "log_gorm_env"
	GormMaxOpenConn = "max_open_conn"
	GormMaxIdleConn = "max_idle_conn"
)

func ProvideDBConfig(lc fx.Lifecycle, logger *zap.Logger) DBConfig {
	cfg := DBConfig{}
	cfg.Host = viper.GetString(DbHost)
	if cfg.Host == "" {
		logger.Panic("failed to read db env", zap.String("env", DbHost))
	}

	cfg.Database = viper.GetString(DbDatabase)
	if cfg.Database == "" {
		logger.Panic("failed to read db env", zap.String("env", DbDatabase))
	}

	cfg.Port = viper.GetUint16(DbPort)
	if cfg.Port == 0 {
		logger.Panic("failed to read db env", zap.String("env", DbPort))
	}

	cfg.User = viper.GetString(DbUser)
	if cfg.User == "" {
		logger.Panic("failed to read db env", zap.String("env", DbUser))
	}

	cfg.Password = viper.GetString(DbPass)
	if cfg.Password == "" {
		logger.Panic("failed to read db env", zap.String("env", DbPass))
	}

	cfg.LogLevel = viper.GetUint16(GormLogMODE)
	if cfg.LogLevel == 0 {
		logger.Panic("failed to read db env", zap.String("env", GormLogMODE))
	}

	cfg.MaxOpenConnection = viper.GetUint16(GormMaxOpenConn)
	if cfg.MaxOpenConnection == 0 {
		logger.Panic("failed to read db env", zap.String("env", GormMaxOpenConn))
	}

	cfg.MaxIdleConnection = viper.GetUint16(GormMaxIdleConn)
	if cfg.MaxIdleConnection == 0 {
		logger.Panic("failed to read db env", zap.String("env", GormMaxIdleConn))
	}

	return cfg
}

func ProvideDBConn(db *gorm.DB, logger *zap.Logger) DBConn {
	logger.Info("database connection established")

	return &dbConn{
		GormConnectionPool: db,
	}
}
