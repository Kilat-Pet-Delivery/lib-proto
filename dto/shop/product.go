package shop

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// ProductDTO is the response representation of a catalog product.
type ProductDTO struct {
	ID            uuid.UUID `json:"id"`
	ShopID        uuid.UUID `json:"shop_id"`
	SKU           string    `json:"sku"`
	Name          string    `json:"name"`
	Description   string    `json:"description,omitempty"`
	PriceCents    int64     `json:"price_cents"`
	Currency      string    `json:"currency"`
	ImageURL      string    `json:"image_url,omitempty"`
	Category      string    `json:"category,omitempty"`
	Active        bool      `json:"active"`
	StockQuantity int64     `json:"stock_quantity"`
	Version       int64     `json:"version"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// CreateProductRequest is the request body for adding a new product to a
// shop's catalog. The caller's shop is resolved from the auth context.
type CreateProductRequest struct {
	SKU         string `json:"sku" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description,omitempty"`
	PriceCents  int64  `json:"price_cents" binding:"required"`
	Currency    string `json:"currency" binding:"required"`
	ImageURL    string `json:"image_url,omitempty"`
	Category    string `json:"category,omitempty"`
}

// Validate returns nil if the request is valid, or a descriptive error
// otherwise.
func (r CreateProductRequest) Validate() error {
	if r.SKU == "" {
		return fmt.Errorf("sku is required")
	}
	if r.Name == "" {
		return fmt.Errorf("name is required")
	}
	if r.PriceCents <= 0 {
		return fmt.Errorf("price_cents must be greater than 0")
	}
	if r.Currency == "" {
		return fmt.Errorf("currency is required")
	}
	return nil
}

// UpdateProductRequest is the request body for partial product updates.
// Pointer fields distinguish "not provided" from "set to zero value".
type UpdateProductRequest struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	PriceCents  *int64  `json:"price_cents,omitempty"`
	Currency    *string `json:"currency,omitempty"`
	ImageURL    *string `json:"image_url,omitempty"`
	Category    *string `json:"category,omitempty"`
	Active      *bool   `json:"active,omitempty"`
}

// Validate returns nil if every supplied field is valid. Unset fields are
// skipped.
func (r UpdateProductRequest) Validate() error {
	if r.PriceCents != nil && *r.PriceCents <= 0 {
		return fmt.Errorf("price_cents must be greater than 0")
	}
	return nil
}

// InventoryLevelDTO is the response representation of a product's on-hand
// stock together with its optimistic-concurrency version token.
type InventoryLevelDTO struct {
	ProductID     uuid.UUID `json:"product_id"`
	ShopID        uuid.UUID `json:"shop_id"`
	StockQuantity int64     `json:"stock_quantity"`
	Version       int64     `json:"version"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// UpdateInventoryRequest is the request body for adjusting on-hand stock.
// ExpectedVersion is the optimistic-lock token returned by the most recent
// read; a mismatch yields HTTP 409 from the server.
type UpdateInventoryRequest struct {
	Quantity        int64  `json:"quantity" binding:"required"`
	Reason          string `json:"reason" binding:"required"`
	ExpectedVersion int64  `json:"expected_version" binding:"required"`
	Note            string `json:"note,omitempty"`
}

// Validate returns nil if quantity is non-negative, reason is a known value,
// and expected_version is non-negative.
//
// Quantity is the absolute target stock level after the adjustment, not a
// delta; therefore negative values are rejected.
func (r UpdateInventoryRequest) Validate() error {
	if r.Quantity < 0 {
		return fmt.Errorf("quantity must be non-negative")
	}
	if r.ExpectedVersion < 0 {
		return fmt.Errorf("expected_version must be non-negative")
	}
	if !IsValidInventoryMovementReason(r.Reason) {
		return fmt.Errorf("reason must be one of: sale, restock, adjustment, return, damage")
	}
	return nil
}
