package usecases

import (
	"context"
	"packform-backend/src/app/orders/models"
	"time"
)

type (
	OrderUsecaseInterface interface {
		GetOrderDetails(ctx context.Context, search string, startDate, endDate time.Time, page, perPage int) ([]*models.OrderDetails, error)
		FeedingDataFromCSV(ctx context.Context, destination string, files []string) error
	}
)
