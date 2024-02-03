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

	// if you set `POPULATE_DATA_FROM=api` it means directly insert data when api server start
	// which can make error every time you rebuild golang because it will try insert existing data
	// so set `POPULATE_DATA_FROM=cli` after you run `POPULATE_DATA_FROM=api`
	// it will skip insert data every server run
	// populate can only run one once whether api or cli, you should drop all tables if you test both
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
