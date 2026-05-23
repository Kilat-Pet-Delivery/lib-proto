package events

import (
	"time"

	"github.com/google/uuid"
)

// Kafka topic for booking events.
const TopicBookingEvents = "booking.events"

// Booking event types.
const (
	BookingRequested      = "booking.requested"
	BookingAccepted       = "booking.accepted"
	BookingRunnerMatched  = "booking.runner_matched"
	BookingPetPickedUp    = "booking.pet_picked_up"
	BookingDeliveryInProg = "booking.delivery_in_progress"
	BookingDeliveryConfirmed = "booking.delivery_confirmed"
	BookingCompleted      = "booking.completed"
	BookingCancelled      = "booking.cancelled"

	// Shop-side substates introduced by Plan C. Routed on the existing
	// TopicBookingEvents topic; consumed by service-shop, service-tracking,
	// service-notification, and the api-gateway.
	BookingAcceptedByShop = "booking.accepted_by_shop"
	BookingPreparing      = "booking.preparing"
	BookingReadyForPickup = "booking.ready_for_pickup"
)

// BookingRequestedEvent is published when an owner creates a new booking.
type BookingRequestedEvent struct {
	BookingID      uuid.UUID `json:"booking_id"`
	BookingNumber  string    `json:"booking_number"`
	OwnerID        uuid.UUID `json:"owner_id"`
	PetType        string    `json:"pet_type"`
	PetName        string    `json:"pet_name"`
	PickupLat      float64   `json:"pickup_lat"`
	PickupLng      float64   `json:"pickup_lng"`
	DropoffLat     float64   `json:"dropoff_lat"`
	DropoffLng     float64   `json:"dropoff_lng"`
	EstimatedPrice int64     `json:"estimated_price_cents"`
	Currency       string    `json:"currency"`
	OccurredAt     time.Time `json:"occurred_at"`
}

// BookingAcceptedEvent is published when a runner accepts a booking.
type BookingAcceptedEvent struct {
	BookingID     uuid.UUID `json:"booking_id"`
	BookingNumber string    `json:"booking_number"`
	RunnerID      uuid.UUID `json:"runner_id"`
	OwnerID       uuid.UUID `json:"owner_id"`
	OccurredAt    time.Time `json:"occurred_at"`
}

// PetPickedUpEvent is published when the runner picks up the pet.
type PetPickedUpEvent struct {
	BookingID     uuid.UUID `json:"booking_id"`
	BookingNumber string    `json:"booking_number"`
	RunnerID      uuid.UUID `json:"runner_id"`
	OwnerID       uuid.UUID `json:"owner_id"`
	PickedUpAt    time.Time `json:"picked_up_at"`
	OccurredAt    time.Time `json:"occurred_at"`
}

// DeliveryConfirmedEvent is published when the owner confirms pet delivery.
type DeliveryConfirmedEvent struct {
	BookingID     uuid.UUID `json:"booking_id"`
	BookingNumber string    `json:"booking_number"`
	RunnerID      uuid.UUID `json:"runner_id"`
	OwnerID       uuid.UUID `json:"owner_id"`
	DeliveredAt   time.Time `json:"delivered_at"`
	OccurredAt    time.Time `json:"occurred_at"`
}

// BookingCompletedEvent is published when the booking is fully completed.
type BookingCompletedEvent struct {
	BookingID     uuid.UUID `json:"booking_id"`
	BookingNumber string    `json:"booking_number"`
	RunnerID      uuid.UUID `json:"runner_id"`
	OwnerID       uuid.UUID `json:"owner_id"`
	FinalPrice    int64     `json:"final_price_cents"`
	Currency      string    `json:"currency"`
	OccurredAt    time.Time `json:"occurred_at"`
}

// BookingCancelledEvent is published when a booking is cancelled.
type BookingCancelledEvent struct {
	BookingID     uuid.UUID  `json:"booking_id"`
	BookingNumber string     `json:"booking_number"`
	CancelledBy   uuid.UUID  `json:"cancelled_by"`
	Reason        string     `json:"reason"`
	OccurredAt    time.Time  `json:"occurred_at"`
}

// BookingAcceptedByShopEvent is published when a shop owner/manager accepts
// an incoming shop-fulfilled booking. Triggers shop-side preparation flow.
type BookingAcceptedByShopEvent struct {
	BookingID     uuid.UUID `json:"booking_id"`
	BookingNumber string    `json:"booking_number"`
	ShopID        uuid.UUID `json:"shop_id"`
	AcceptedBy    uuid.UUID `json:"accepted_by_user_id"`
	OccurredAt    time.Time `json:"occurred_at"`
}

// BookingPreparingEvent is published when shop staff transition a booking
// into the preparing state (order is being assembled).
type BookingPreparingEvent struct {
	BookingID     uuid.UUID `json:"booking_id"`
	BookingNumber string    `json:"booking_number"`
	ShopID        uuid.UUID `json:"shop_id"`
	StartedBy     uuid.UUID `json:"started_by_user_id"`
	OccurredAt    time.Time `json:"occurred_at"`
}

// BookingReadyForPickupEvent is published when the shop signals the order is
// staged and ready for runner pickup. Includes the QR pickup token so that
// service-tracking / service-notification can surface it to the runner.
type BookingReadyForPickupEvent struct {
	BookingID     uuid.UUID `json:"booking_id"`
	BookingNumber string    `json:"booking_number"`
	ShopID        uuid.UUID `json:"shop_id"`
	QRPickupToken string    `json:"qr_pickup_token"`
	OccurredAt    time.Time `json:"occurred_at"`
}
