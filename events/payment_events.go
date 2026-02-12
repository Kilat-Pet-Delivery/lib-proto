package events

import (
	"time"

	"github.com/google/uuid"
)

// Kafka topic for payment events.
const TopicPaymentEvents = "payment.events"

// Payment event types.
const (
	PaymentEscrowCreated  = "payment.escrow_created"
	PaymentEscrowHeld     = "payment.escrow_held"
	PaymentEscrowReleased = "payment.escrow_released"
	PaymentEscrowRefunded = "payment.escrow_refunded"
	PaymentFailed         = "payment.failed"
)

// EscrowCreatedEvent is published when a payment escrow is initiated.
type EscrowCreatedEvent struct {
	PaymentID   uuid.UUID `json:"payment_id"`
	BookingID   uuid.UUID `json:"booking_id"`
	OwnerID     uuid.UUID `json:"owner_id"`
	AmountCents int64     `json:"amount_cents"`
	Currency    string    `json:"currency"`
	OccurredAt  time.Time `json:"occurred_at"`
}

// EscrowHeldEvent is published when funds are successfully held (Stripe authorize).
type EscrowHeldEvent struct {
	PaymentID       uuid.UUID `json:"payment_id"`
	BookingID       uuid.UUID `json:"booking_id"`
	StripePaymentID string    `json:"stripe_payment_id"`
	AmountCents     int64     `json:"amount_cents"`
	Currency        string    `json:"currency"`
	OccurredAt      time.Time `json:"occurred_at"`
}

// EscrowReleasedEvent is published when funds are released to the runner.
type EscrowReleasedEvent struct {
	PaymentID       uuid.UUID `json:"payment_id"`
	BookingID       uuid.UUID `json:"booking_id"`
	RunnerID        uuid.UUID `json:"runner_id"`
	RunnerPayout    int64     `json:"runner_payout_cents"`
	PlatformFee     int64     `json:"platform_fee_cents"`
	Currency        string    `json:"currency"`
	OccurredAt      time.Time `json:"occurred_at"`
}

// EscrowRefundedEvent is published when funds are refunded to the owner.
type EscrowRefundedEvent struct {
	PaymentID    uuid.UUID `json:"payment_id"`
	BookingID    uuid.UUID `json:"booking_id"`
	OwnerID      uuid.UUID `json:"owner_id"`
	AmountCents  int64     `json:"amount_cents"`
	Currency     string    `json:"currency"`
	RefundReason string    `json:"refund_reason"`
	OccurredAt   time.Time `json:"occurred_at"`
}

// PaymentFailedEvent is published when a payment operation fails.
type PaymentFailedEvent struct {
	PaymentID  uuid.UUID `json:"payment_id"`
	BookingID  uuid.UUID `json:"booking_id"`
	Reason     string    `json:"reason"`
	OccurredAt time.Time `json:"occurred_at"`
}
