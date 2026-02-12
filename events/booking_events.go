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
