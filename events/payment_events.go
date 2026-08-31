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

	// Shop wallet withdrawal lifecycle (Plan C). Routed on the existing
	// TopicPaymentEvents topic and consumed by service-notification and
	// service-shop.
	WithdrawalRequested = "payment.withdrawal_requested"
	WithdrawalPaid      = "payment.withdrawal_paid"
	WithdrawalFailed    = "payment.withdrawal_failed"
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
	PaymentID    uuid.UUID `json:"payment_id"`
	BookingID    uuid.UUID `json:"booking_id"`
	RunnerID     uuid.UUID `json:"runner_id"`
	RunnerPayout int64     `json:"runner_payout_cents"`
	PlatformFee  int64     `json:"platform_fee_cents"`
	Currency     string    `json:"currency"`
	OccurredAt   time.Time `json:"occurred_at"`
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

// WithdrawalRequestedEvent is published when a shop owner submits a
// withdrawal against their available wallet balance.
type WithdrawalRequestedEvent struct {
	WithdrawalID  uuid.UUID `json:"withdrawal_id"`
	ShopID        uuid.UUID `json:"shop_id"`
	RequestedBy   uuid.UUID `json:"requested_by_user_id"`
	AmountCents   int64     `json:"amount_cents"`
	Currency      string    `json:"currency"`
	DestinationID string    `json:"destination_id"`
	OccurredAt    time.Time `json:"occurred_at"`
}

// WithdrawalPaidEvent is published when the payout processor confirms the
// withdrawal landed in the destination bank account.
type WithdrawalPaidEvent struct {
	WithdrawalID uuid.UUID `json:"withdrawal_id"`
	ShopID       uuid.UUID `json:"shop_id"`
	AmountCents  int64     `json:"amount_cents"`
	Currency     string    `json:"currency"`
	ProcessorRef string    `json:"processor_ref,omitempty"`
	OccurredAt   time.Time `json:"occurred_at"`
}

// WithdrawalFailedEvent is published when the payout processor reports a
// terminal failure; service-payment reverses the held funds back into the
// shop wallet's available balance.
type WithdrawalFailedEvent struct {
	WithdrawalID uuid.UUID `json:"withdrawal_id"`
	ShopID       uuid.UUID `json:"shop_id"`
	AmountCents  int64     `json:"amount_cents"`
	Currency     string    `json:"currency"`
	Reason       string    `json:"reason"`
	OccurredAt   time.Time `json:"occurred_at"`
}
