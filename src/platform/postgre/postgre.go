package postgre

import (
	"context"
	"fmt"
	"log"
	"packform-backend/src/pkg/config"

	"github.com/jackc/pgx/v5"
)

type (
	connPostgre struct {
		conn *pgx.Conn
	}
)

func NewPostgreConn(config config.AppConfig) *connPostgre {
	descriptor := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.DbConfig.DbHost, config.DbConfig.DbPort, config.DbConfig.DbUsername, config.DbConfig.DbPassword, config.DbConfig.DbName, config.DbConfig.DbSslMode)

	configPgx, err := pgx.ParseConfig(descriptor)
	if err != nil {
		log.Fatalf("error happen when parse configuration:%s\n", err)
	}
	configPgx.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

	conn, err := pgx.ConnectConfig(context.Background(), configPgx)
	if err != nil {
		log.Fatalf("error happen when connect to database:%s\n", err)
	}

	return &connPostgre{conn: conn}
}

func (p *connPostgre) GetConnection() *pgx.Conn {
	return p.conn
}
