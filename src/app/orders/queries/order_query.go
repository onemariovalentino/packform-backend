package queries

import (
	"context"
	"database/sql"
	"fmt"
	"packform-backend/src/app/orders/models"
	"time"

	"gorm.io/gorm"
)

type (
	OrderQuery struct {
		db *gorm.DB
	}
)

func New(db *gorm.DB) OrderQueryInterface {
	return &OrderQuery{db: db}
}

func (q *OrderQuery) GetOrderDetails(ctx context.Context, search string, startDate, endDate time.Time, page, perPage int) ([]*models.OrderDetails, error) {
	offset := (page * perPage) - perPage

	query := `SELECT o.order_name as order_name,
				   co.company_name as company_name,
				   c.name as customer_name,
				   o.created_at as order_date,
    			   SUM(oid.delivered_quantity * oi.price_per_unit) as delivered_amount,
    			   SUM(oi.price_per_unit * oi.quantity) as total_amount
			FROM tbl_orders o
			LEFT JOIN tbl_customers c ON c.user_id = o.customer_id
			LEFT JOIN tbl_companies co ON co.company_id = c.company_id
			LEFT JOIN tbl_order_items oi ON oi.order_id = o.id
			LEFT JOIN tbl_order_item_deliveries oid ON oid.order_item_id=oi.id
			WHERE 1=1`

	if search != "" {
		query += fmt.Sprintf(` AND (o.order_name LIKE '%%%s%%' OR oi.product::text LIKE '%%%s%%')`, search, search)
	}

	if !startDate.IsZero() && !endDate.IsZero() {
		query += fmt.Sprintf(` AND (DATE(o.created_at) BETWEEN '%s' AND '%s')`, startDate.Format("2006-01-02"), endDate.Format("2006-01-02"))
	}
	query += ` GROUP BY o.id, co.company_name, c.name`
	query += ` ORDER BY o.created_at DESC`

	query += fmt.Sprintf(` LIMIT %d OFFSET %d`, perPage, offset)

	rs, err := q.db.Raw(query).Rows()
	if err != nil {
		return []*models.OrderDetails{}, err
	}

	orderDetails := []*models.OrderDetails{}
	var orderName, companyName, customerName string
	var orderTime time.Time
	var deliveredAmount sql.NullFloat64
	var totalAmount float64
	for rs.Next() {
		err := rs.Scan(&orderName, &companyName, &customerName, &orderTime, &deliveredAmount, &totalAmount)
		if err != nil {
			return []*models.OrderDetails{}, err
		}

		tz, _ := time.LoadLocation("Australia/Melbourne")
		dt, _ := time.Parse(time.RFC3339, orderTime.Format(time.RFC3339))

		orderDetail := &models.OrderDetails{
			OrderName:       orderName,
			CompanyName:     companyName,
			CustomerName:    customerName,
			OrderDate:       dt.In(tz),
			DeliveredAmount: deliveredAmount,
			TotalAmount:     totalAmount,
		}
		orderDetails = append(orderDetails, orderDetail)
	}

	return orderDetails, nil
}

func (q *OrderQuery) CreateCustomerCompanies(ctx context.Context, companies []*models.Company) error {

	/*
		// You can use this syntax as alternative
		sqlStr := "INSERT INTO tbl_companies(company_id, company_name) VALUES "
		vals := []interface{}{}
		for _, row := range companies {
			sqlStr += "(?, ?),"
			vals = append(vals, row.CompanyID, row.CompanyName)
		}
		sqlStr = sqlStr[0 : len(sqlStr)-1]
		if err := q.db.Exec(sqlStr, vals...).Error; err != nil {
			return err
		}
	*/

	if err := q.db.Create(companies).Error; err != nil {
		return err
	}

	return nil
}

func (q *OrderQuery) CreateCustomers(ctx context.Context, customers []*models.Customer) error {

	/*
		// You can use this syntax as alternative
		sqlStr := "INSERT INTO tbl_customers(user_id, login, password, name, company_id, credit_cards) VALUES "
		vals := []interface{}{}
		for _, row := range customers {
			sqlStr += "(?, ?, ?, ?, ?, ?),"
			vals = append(vals, row.UserID, row.Login, row.Password, row.Name, row.CompanyID, row.CreditCards)
		}
		sqlStr = sqlStr[0 : len(sqlStr)-1]
		if err := q.db.Exec(sqlStr, vals...).Error; err != nil {
			return err
		}
	*/

	if err := q.db.Create(customers).Error; err != nil {
		return err
	}

	return nil
}

func (q *OrderQuery) CreateOrders(ctx context.Context, orders []*models.Order) error {

	/*
		// You can use this syntax as alternative
		sqlStr := "INSERT INTO tbl_orders(id, created_at, order_name, customer_id) VALUES "
		vals := []interface{}{}
		for _, row := range orders {
			sqlStr += "(?, ?, ?, ?),"
			vals = append(vals, row.ID, row.CreatedAt, row.OrderName, row.CustomerID)
		}
		sqlStr = sqlStr[0 : len(sqlStr)-1]
		if err := q.db.Exec(sqlStr, vals...).Error; err != nil {
			return err
		}
	*/

	if err := q.db.Create(orders).Error; err != nil {
		return err
	}

	return nil
}

func (q *OrderQuery) CreateOrderItems(ctx context.Context, orderItems []*models.OrderItem) error {

	// You can use this syntax as alternative
	sqlStr := "INSERT INTO tbl_order_items(id, order_id, price_per_unit, quantity, product) VALUES "
	vals := []interface{}{}
	for _, row := range orderItems {
		sqlStr += "(?, ?, ?, ?, ?),"
		vals = append(vals, row.ID, row.OrderID, row.PricePerUnit, row.Qty, row.Product)
	}
	sqlStr = sqlStr[0 : len(sqlStr)-1]
	if err := q.db.Exec(sqlStr, vals...).Error; err != nil {
		return err
	}

	/*
		if err := q.db.Create(orderItems).Error; err != nil {
			return err
		}
	*/

	return nil
}

func (q *OrderQuery) CreateOrderItemDeliveries(ctx context.Context, orderItemDeliveries []*models.OrderItemDelivery) error {

	// You can use this syntax as alternative
	sqlStr := "INSERT INTO tbl_order_item_deliveries(id, order_item_id, delivered_quantity) VALUES "
	vals := []interface{}{}
	for _, row := range orderItemDeliveries {
		sqlStr += "(?, ?, ?),"
		vals = append(vals, row.ID, row.OrderItemID, row.DeliveredQty)
	}
	sqlStr = sqlStr[0 : len(sqlStr)-1]
	if err := q.db.Exec(sqlStr, vals...).Error; err != nil {
		return err
	}

	/*
		if err := q.db.Create(orderItemDeliveries).Error; err != nil {
			return err
		}
	*/

	return nil
}
