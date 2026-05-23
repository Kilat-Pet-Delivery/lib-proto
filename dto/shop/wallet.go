package shop

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// WithdrawalStatus enumerates the lifecycle states of a shop withdrawal.
const (
	WithdrawalStatusPending    = "pending"
	WithdrawalStatusProcessing = "processing"
	WithdrawalStatusPaid       = "paid"
	WithdrawalStatusFailed     = "failed"
)

// ShopWalletDTO is the response representation of a shop's running wallet
// balance.
type ShopWalletDTO struct {
	ShopID              uuid.UUID `json:"shop_id"`
	AvailableCents      int64     `json:"available_cents"`
	PendingCents        int64     `json:"pending_cents"`
	LifetimeEarnedCents int64     `json:"lifetime_earned_cents"`
	Currency            string    `json:"currency"`
	UpdatedAt           time.Time `json:"updated_at"`
}

// ShopWalletLedgerEntryDTO is a single immutable wallet ledger entry.
// ChangeAmount is signed: positive for credits, negative for debits.
type ShopWalletLedgerEntryDTO struct {
	ID           uuid.UUID  `json:"id"`
	ShopID       uuid.UUID  `json:"shop_id"`
	BookingID    *uuid.UUID `json:"booking_id,omitempty"`
	WithdrawalID *uuid.UUID `json:"withdrawal_id,omitempty"`
	EntryType    string     `json:"entry_type"`
	ChangeCents  int64      `json:"change_cents"`
	BalanceAfter int64      `json:"balance_after_cents"`
	Currency     string     `json:"currency"`
	Memo         string     `json:"memo,omitempty"`
	OccurredAt   time.Time  `json:"occurred_at"`
}

// WithdrawalDTO is the response representation of a withdrawal request.
type WithdrawalDTO struct {
	ID               uuid.UUID  `json:"id"`
	ShopID           uuid.UUID  `json:"shop_id"`
	AmountCents      int64      `json:"amount_cents"`
	Currency         string     `json:"currency"`
	Status           string     `json:"status"`
	DestinationID    string     `json:"destination_id"`
	FailureReason    string     `json:"failure_reason,omitempty"`
	ProcessorRef     string     `json:"processor_ref,omitempty"`
	RequestedAt      time.Time  `json:"requested_at"`
	ProcessingAt     *time.Time `json:"processing_at,omitempty"`
	PaidAt           *time.Time `json:"paid_at,omitempty"`
	FailedAt         *time.Time `json:"failed_at,omitempty"`
}

// WithdrawRequest is the request body sent by a shop owner to draw down
// available wallet funds to a verified bank destination.
type WithdrawRequest struct {
	AmountCents   int64  `json:"amount_cents" binding:"required"`
	DestinationID string `json:"destination_id" binding:"required"`
}

// Validate returns nil if the request supplies a positive amount and a
// non-empty destination identifier.
func (r WithdrawRequest) Validate() error {
	if r.AmountCents <= 0 {
		return fmt.Errorf("amount_cents must be greater than 0")
	}
	if r.DestinationID == "" {
		return fmt.Errorf("destination_id is required")
	}
	return nil
}

// IsValidWithdrawalStatus returns true when s is one of the four known
// withdrawal statuses.
func IsValidWithdrawalStatus(s string) bool {
	switch s {
	case WithdrawalStatusPending, WithdrawalStatusProcessing,
		WithdrawalStatusPaid, WithdrawalStatusFailed:
		return true
	default:
		return false
	}
}
