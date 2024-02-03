package db

import (
	"log"
	"packform-backend/src/app/orders/models"
	"packform-backend/src/pkg/config"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func BuildGormDB(conn *pgx.Conn) *gorm.DB {
	c := stdlib.OpenDB(*conn.Config())

	gorm, err := gorm.Open(postgres.New(postgres.Config{Conn: c}))
	if err != nil {
		log.Fatalf("error happen when connect to database:%s\n", err)
	}

	if config.Env.Platform == "local" {
		// use for creating  `product_kind` type for enum
		totalRow := gorm.Exec("select exists (select 1 from pg_type where typname = 'product_kind')").RowsAffected
		if totalRow == 0 {
			gorm.Exec("CREATE TYPE product_kind AS ENUM ('Corrugated Box','Hand Sanitizer')")
		}
		gorm.AutoMigrate(models.Company{}, models.Customer{}, models.Order{}, models.OrderItem{}, models.OrderItemDelivery{})
	}

	poolDB, err := gorm.DB()
	if err != nil {
		log.Fatalf("error happen when connect to database:%s\n", err)
	}
	poolDB.SetMaxOpenConns(config.Env.DbConfig.DbMaxOpenConnection)
	poolDB.SetMaxIdleConns(config.Env.DbConfig.DbMaxIdleConnection)
	return gorm
}
