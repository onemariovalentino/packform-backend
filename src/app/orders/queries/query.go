package queries

import (
	"context"
	"packform-backend/src/app/orders/models"
	"time"
)

type (
	OrderQueryInterface interface {
		GetOrderDetails(ctx context.Context, search string, startDate, endDate time.Time, page, perPage int, sortDirection string) (*models.OrderDetails, error)
		CreateCustomerCompanies(ctx context.Context, companies []*models.Company) error
		CreateCustomers(ctx context.Context, customers []*models.Customer) error
		CreateOrders(ctx context.Context, orders []*models.Order) error
		CreateOrderItems(ctx context.Context, orderItems []*models.OrderItem) error
		CreateOrderItemDeliveries(ctx context.Context, orderItemDeliveries []*models.OrderItemDelivery) error
	}
)
