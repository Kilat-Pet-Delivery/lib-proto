// Package shop hosts the shared DTOs and enums for the Pet Shop bounded
// context. These types are consumed by service-shop, service-booking,
// service-identity, service-payment, service-notification, and the
// api-gateway response transformers.
package shop

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// ShopStatus enumerates the operational state of a shop.
const (
	ShopStatusOpen   = "open"
	ShopStatusBusy   = "busy"
	ShopStatusClosed = "closed"
	ShopStatusPaused = "paused"
)

// ShopCategory enumerates the supported shop business categories.
const (
	ShopCategoryVet      = "vet"
	ShopCategoryGrooming = "grooming"
	ShopCategorySupplies = "supplies"
)

// ShopDTO is the response representation of a shop aggregate.
type ShopDTO struct {
	ID           uuid.UUID  `json:"id"`
	OwnerUserID  uuid.UUID  `json:"owner_user_id"`
	Name         string     `json:"name"`
	Slug         string     `json:"slug"`
	Address      string     `json:"address"`
	Latitude     float64    `json:"latitude"`
	Longitude    float64    `json:"longitude"`
	Phone        string     `json:"phone"`
	Category     string     `json:"category"`
	Status       string     `json:"status"`
	AutoCloseAt  *time.Time `json:"auto_close_at,omitempty"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

// CreateShopRequest is the request body for shop self-registration by an owner.
type CreateShopRequest struct {
	Name      string  `json:"name" binding:"required"`
	Slug      string  `json:"slug" binding:"required"`
	Address   string  `json:"address" binding:"required"`
	Latitude  float64 `json:"latitude" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
	Phone     string  `json:"phone" binding:"required"`
	Category  string  `json:"category" binding:"required"`
}

// Validate returns nil if the request is valid, or a descriptive error if any
// required field is missing or invalid.
func (r CreateShopRequest) Validate() error {
	if r.Name == "" {
		return fmt.Errorf("name is required")
	}
	if r.Slug == "" {
		return fmt.Errorf("slug is required")
	}
	if r.Address == "" {
		return fmt.Errorf("address is required")
	}
	if r.Phone == "" {
		return fmt.Errorf("phone is required")
	}
	if !IsValidShopCategory(r.Category) {
		return fmt.Errorf("category must be one of: vet, grooming, supplies")
	}
	return nil
}

// UpdateShopRequest is the request body for updating shop profile fields.
type UpdateShopRequest struct {
	Name      string  `json:"name,omitempty"`
	Address   string  `json:"address,omitempty"`
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
	Phone     string  `json:"phone,omitempty"`
	Category  string  `json:"category,omitempty"`
}

// Validate returns an error if any non-empty field is invalid.
// Empty fields are skipped (partial updates).
func (r UpdateShopRequest) Validate() error {
	if r.Category != "" && !IsValidShopCategory(r.Category) {
		return fmt.Errorf("category must be one of: vet, grooming, supplies")
	}
	return nil
}

// UpdateShopStatusRequest is the request body for the merchant-facing
// PUT /shops/{id}/status endpoint. AutoCloseAt is optional; when set it
// schedules a deferred close (e.g. busy until 16:00 then closed).
type UpdateShopStatusRequest struct {
	Status      string     `json:"status" binding:"required"`
	AutoCloseAt *time.Time `json:"auto_close_at,omitempty"`
}

// Validate returns nil if the requested status is one of the four known values.
func (r UpdateShopStatusRequest) Validate() error {
	if !IsValidShopStatus(r.Status) {
		return fmt.Errorf("status must be one of: open, busy, closed, paused")
	}
	return nil
}

// IsValidShopStatus returns true when s is one of the four known shop statuses.
func IsValidShopStatus(s string) bool {
	switch s {
	case ShopStatusOpen, ShopStatusBusy, ShopStatusClosed, ShopStatusPaused:
		return true
	default:
		return false
	}
}

// IsValidShopCategory returns true when c is one of the three known shop
// categories.
func IsValidShopCategory(c string) bool {
	switch c {
	case ShopCategoryVet, ShopCategoryGrooming, ShopCategorySupplies:
		return true
	default:
		return false
	}
}
