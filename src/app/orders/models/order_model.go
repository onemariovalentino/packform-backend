package models

import (
	"database/sql/driver"
	"time"
)

type ProductKind string

func (k ProductKind) Value() (driver.Value, error) {
	return string(k), nil
}

func (k *ProductKind) Scan(value interface{}) error {
	stringValue, ok := value.([]uint8)
	if ok {
		*k = ProductKind(stringValue)
	} else {
		*k = ProductKind(value.(string))
	}
	return nil
}

const (
	CorrugatedBox ProductKind = "Corrugated Box"
	HandSanitizer ProductKind = "Hand Sanitizer"
)

type (
	Company struct {
		CompanyID   int      `gorm:"column:company_id;type:int;primaryKey;autoIncrement"`
		CompanyName string   `gorm:"column:company_name;type:varchar(100);not null"`
		Customer    Customer `gorm:"foreignKey:CompanyID"`
	}

	Customer struct {
		UserID      string  `gorm:"column:user_id;type:varchar(30);primaryKey"`
		Login       string  `gorm:"column:login;type:varchar(30);not null"`
		Password    string  `gorm:"column:password;type:varchar(100);not null"`
		Name        string  `gorm:"column:name;type:varchar(60);not null"`
		CompanyID   int     `gorm:"column:company_id;type:int;not null"`
		CreditCards string  `gorm:"column:credit_cards;type:varchar(60);not null"`
		Order       []Order `gorm:"foreignKey:CustomerID"`
	}

	Order struct {
		ID         int         `gorm:"column:id;primaryKey;autoIncrement"`
		CreatedAt  time.Time   `gorm:"column:created_at;not null"`
		OrderName  string      `gorm:"column:order_name;type:varchar(100);not null"`
		CustomerID string      `gorm:"column:customer_id;type:varchar(30);not null"`
		OrderItem  []OrderItem `gorm:"foreignKey:OrderID"`
	}

	OrderItem struct {
		ID                int                 `gorm:"column:id;primaryKey;autoIncrement"`
		OrderID           int                 `gorm:"column:order_id"`
		PricePerUnit      float64             `gorm:"column:price_per_unit;type:decimal(12,4)"`
		Qty               int                 `gorm:"column:quantity"`
		Product           string              `gorm:"column:product;type:product_kind"`
		OrderItemDelivery []OrderItemDelivery `gorm:"foreignKey:OrderItemID"`
	}

	OrderItemDelivery struct {
		ID           int `gorm:"column:id;primaryKey;autoIncrement"`
		OrderItemID  int `gorm:"column:order_item_id;type:int"`
		DeliveredQty int `gorm:"column:delivered_quantity;type:int"`
	}

	OrderDetail struct {
		OrderID         int64   `json:"order_id" gorm:"column:order_id"`
		OrderName       string  `json:"order_name" gorm:"column:order_name"`
		Product         string  `json:"product" gorm:"column:product_name"`
		CompanyName     string  `json:"company_name" gorm:"column:company_name"`
		CustomerName    string  `json:"customer_name" gorm:"column:customer_name"`
		OrderDate       string  `json:"order_date" gorm:"column:order_date"`
		DeliveredAmount float64 `json:"delivered_amount" gorm:"column:delivered_amount"`
		TotalAmount     float64 `json:"total_amount" gorm:"column:total_amount"`
	}

	OrderDetails struct {
		TotalData   int            `json:"total_data"`
		TotalAmount float64        `json:"total_amount"`
		Orders      []*OrderDetail `json:"orders"`
	}
)

// Change table name to `tbl_companies`
func (c Company) TableName() string {
	return "tbl_companies"
}

// Change table name to `tbl_customers`
func (c Customer) TableName() string {
	return "tbl_customers"
}

// Change table name to `tbl_orders`
func (c Order) TableName() string {
	return "tbl_orders"
}

// Change table name to `tbl_order_items`
func (c OrderItem) TableName() string {
	return "tbl_order_items"
}

// Change table name to `tbl_order_item_deliveries`
func (c OrderItemDelivery) TableName() string {
	return "tbl_order_item_deliveries"
}
