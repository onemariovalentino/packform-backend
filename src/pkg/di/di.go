package di

import (
	"context"
	"packform-backend/src/app/orders/queries"
	"packform-backend/src/app/orders/usecases"
	"packform-backend/src/pkg/config"
	"packform-backend/src/platform/db"
	"packform-backend/src/platform/postgre"
)

type (
	Dependency struct {
		OrderUsecase usecases.OrderUsecaseInterface
	}
)

func NewDependency() *Dependency {
	postgresConnection := postgre.NewPostgreConn(config.Env).GetConnection()
	gormDB := db.BuildGormDB(postgresConnection)

	orderQuery := queries.New(gormDB)
	orderUsecase := usecases.New(orderQuery)
	if config.Env.PopulateDataFrom == "api" {
		ctx := context.Background()
		orderUsecase.FeedingDataFromCSV(ctx, "companies", []string{`files/csv/Test task - Postgres - customer_companies.csv`})
		orderUsecase.FeedingDataFromCSV(ctx, "customers", []string{`files/csv/Test task - Postgres - customers.csv`})
		orderUsecase.FeedingDataFromCSV(ctx, "orders", []string{`files/csv/Test task - Postgres - orders.csv`})
		orderUsecase.FeedingDataFromCSV(ctx, "order_items", []string{`files/csv/Test task - Postgres - order_items.csv`})
		orderUsecase.FeedingDataFromCSV(ctx, "order_item_deliveries", []string{`files/csv/Test task - Postgres - deliveries.csv`})
	}

	return &Dependency{
		OrderUsecase: orderUsecase,
	}
}
