// Package booking hosts the shared booking-status enum plus the cross-service
// BookingDTO that the api-gateway, service-shop, service-notification, and
// future analytics consumers project booking data into.
//
// The authoritative state machine lives in service-booking; this package
// exposes only the closed enum of recognized status strings so that downstream
// services agree on the wire format.
package booking

import "fmt"

// BookingStatus is the wire-format string for a booking's current state.
type BookingStatus string

// The eight recognized booking statuses. The first six are runner-flow
// states inherited from Plan A; the final three substates (AcceptedByShop,
// Preparing, ReadyForPickup) are introduced by Plan C for shop-fulfilled
// orders that route through service-shop before runner pickup.
const (
	BookingStatusRequested      BookingStatus = "requested"
	BookingStatusAccepted       BookingStatus = "accepted"
	BookingStatusInProgress     BookingStatus = "in_progress"
	BookingStatusDelivered      BookingStatus = "delivered"
	BookingStatusCompleted      BookingStatus = "completed"
	BookingStatusCancelled      BookingStatus = "cancelled"
	BookingStatusAcceptedByShop BookingStatus = "accepted_by_shop"
	BookingStatusPreparing      BookingStatus = "preparing"
	BookingStatusReadyForPickup BookingStatus = "ready_for_pickup"
)

// allBookingStatuses is the closed set used by IsValid. Listed once so adding
// a new status only requires one edit.
var allBookingStatuses = []BookingStatus{
	BookingStatusRequested,
	BookingStatusAccepted,
	BookingStatusInProgress,
	BookingStatusDelivered,
	BookingStatusCompleted,
	BookingStatusCancelled,
	BookingStatusAcceptedByShop,
	BookingStatusPreparing,
	BookingStatusReadyForPickup,
}

// IsValid returns true when s is one of the recognized booking statuses.
func (s BookingStatus) IsValid() bool {
	for _, known := range allBookingStatuses {
		if s == known {
			return true
		}
	}
	return false
}

// String returns the string representation of the status.
func (s BookingStatus) String() string {
	return string(s)
}

// ParseBookingStatus converts a string to a BookingStatus, returning an
// error if the input is not one of the recognized values.
func ParseBookingStatus(s string) (BookingStatus, error) {
	status := BookingStatus(s)
	if !status.IsValid() {
		return "", fmt.Errorf("invalid booking status: %s", s)
	}
	return status, nil
}

// BookingDTO is the cross-service projection of a booking record. The
// canonical fields and aggregate behavior live in service-booking; this DTO
// captures only what other services and the api-gateway need to wire
// shop-side flows.
//
// QRPickupToken is a pointer so that the field is omitted from JSON when
// the caller lacks the merchant/runner scope. The api-gateway's response
// transformer strips this field for non-privileged scopes; the pointer
// shape lets consumers also nil it out before re-serialization without
// having to know about a sentinel value.
type BookingDTO struct {
	ID            string        `json:"id"`
	BookingNumber string        `json:"booking_number"`
	OwnerID       string        `json:"owner_id"`
	RunnerID      *string       `json:"runner_id,omitempty"`
	ShopID        *string       `json:"shop_id,omitempty"`
	Status        BookingStatus `json:"status"`
	QRPickupToken *string       `json:"qr_pickup_token,omitempty"`
	Notes         string        `json:"notes,omitempty"`
}
