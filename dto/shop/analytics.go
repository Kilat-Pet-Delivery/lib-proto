package shop

import (
	"time"

	"github.com/google/uuid"
)

// AnalyticsPeriod enumerates the supported aggregation windows for the
// sales-aggregate endpoints.
const (
	AnalyticsPeriodDay   = "day"
	AnalyticsPeriodWeek  = "week"
	AnalyticsPeriodMonth = "month"
)

// SalesAggregateDTO is the response representation of a sales-aggregate
// snapshot for one (shop, period_start) pair.
type SalesAggregateDTO struct {
	ShopID         uuid.UUID       `json:"shop_id"`
	Period         string          `json:"period"`
	PeriodStart    time.Time       `json:"period_start"`
	PeriodEnd      time.Time       `json:"period_end"`
	OrderCount     int64           `json:"order_count"`
	GrossSalesCents int64          `json:"gross_sales_cents"`
	NetSalesCents  int64           `json:"net_sales_cents"`
	Currency       string          `json:"currency"`
	TopProducts    []TopProductDTO `json:"top_products,omitempty"`
	RefreshedAt    time.Time       `json:"refreshed_at"`
}

// TopProductDTO summarizes one product's contribution to a sales aggregate.
type TopProductDTO struct {
	ProductID    uuid.UUID `json:"product_id"`
	SKU          string    `json:"sku"`
	Name         string    `json:"name"`
	UnitsSold    int64     `json:"units_sold"`
	RevenueCents int64     `json:"revenue_cents"`
}

// IsValidAnalyticsPeriod returns true when p is one of the three known
// analytics period values.
func IsValidAnalyticsPeriod(p string) bool {
	switch p {
	case AnalyticsPeriodDay, AnalyticsPeriodWeek, AnalyticsPeriodMonth:
		return true
	default:
		return false
	}
}
