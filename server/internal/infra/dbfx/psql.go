package dbfx

import (
	"database/sql"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var GormDebugLogger logger.Interface

func (cfg DBConfig) toPSQLConnStr() string {
	var connectionString strings.Builder

	connectionType := viper.GetString("psql_conn_type")

	connectionString.WriteString("postgres://")

	// Add username and password if provided
	if cfg.User != "" && cfg.Password != "" {
		connectionString.WriteString(cfg.User)
		connectionString.WriteByte(':')
		connectionString.WriteString(cfg.Password)
		connectionString.WriteByte('@')
	}

	// Add host and port
	connectionString.WriteString(cfg.Host)

	if cfg.Port != 0 && connectionType != "SRV" {
		connectionString.WriteByte(':')
		connectionString.WriteString(strconv.Itoa(int(cfg.Port)))
	}

	// Add database name
	connectionString.WriteByte('/')
	connectionString.WriteString(cfg.Database)

	// Add SSL mode if provided
	if connectionType != "SRV" {
		connectionString.WriteString("?sslmode=disable")
	}
	return connectionString.String()
}

type PSQLParams struct {
	fx.In

	Config DBConfig
	Logger *zap.Logger
}

func NewPSQLConn(p PSQLParams) *gorm.DB {

	if err := viper.BindEnv(LogGormEnv); err != nil {
		log.Panic("failed to load env", zap.String("env", LogGormEnv))
	}

	GormLog := viper.GetBool(LogGormEnv)

	connStr := p.Config.toPSQLConnStr()

	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		p.Logger.Error(err.Error())
		log.Panic("failed to establish database connection", zap.Error(err))
	}

	conn.SetMaxIdleConns(int(p.Config.MaxIdleConnection))
	conn.SetMaxOpenConns(int(p.Config.MaxOpenConnection))

	if err := conn.Ping(); err != nil {
		log.Panic("failed to verify database connection", zap.Error(err))
	}

	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}

	if GormLog {
		gormConfig.Logger = logger.New(log.Default(), logger.Config{
			SlowThreshold:             365 * 24 * time.Hour,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		}).LogMode(logger.LogLevel(p.Config.LogLevel))
	}

	GormConnectionPool, err := gorm.Open(postgres.New(postgres.Config{Conn: conn}), gormConfig)
	if err != nil {
		log.Panic("failed to establish gorm connection", zap.Error(err))
	}

	GormDebugLogger = logger.New(log.Default(), logger.Config{
		SlowThreshold:             365 * 24 * time.Hour,
		LogLevel:                  logger.Info,
		IgnoreRecordNotFoundError: true,
		Colorful:                  true,
	})

	return GormConnectionPool
}
