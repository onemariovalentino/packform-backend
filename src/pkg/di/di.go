package di

import (
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

	return &Dependency{
		OrderUsecase: orderUsecase,
	}
}
