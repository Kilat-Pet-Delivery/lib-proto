package dto

import (
	"time"

	"github.com/google/uuid"
)

// Quest describes an earnable loyalty quest.
type Quest struct {
	ID              uuid.UUID      `json:"id"`
	Code            string         `json:"code"`
	Title           string         `json:"title"`
	Description     string         `json:"description"`
	Type            string         `json:"type"`
	Rule            QuestRule      `json:"rule"`
	RewardAmountMYR string         `json:"reward_amount_myr"`
	Active          bool           `json:"active"`
	Progress        *QuestProgress `json:"progress,omitempty"`
	CreatedAt       time.Time      `json:"created_at"`
}

// QuestRule is the JSON rule shape used by the quest engine.
type QuestRule struct {
	SchemaVersion int            `json:"schema_version"`
	Kind          string         `json:"kind"`
	EventMatch    string         `json:"event_match"`
	Filter        map[string]any `json:"filter,omitempty"`
	Threshold     int            `json:"threshold"`
	Window        string         `json:"window"`
	Deadline      string         `json:"deadline,omitempty"`
}

// QuestProgress captures a user's progress toward a quest.
type QuestProgress struct {
	ID            uuid.UUID  `json:"id"`
	QuestID       uuid.UUID  `json:"quest_id"`
	UserID        uuid.UUID  `json:"user_id"`
	Progress      int        `json:"progress"`
	Target        int        `json:"target"`
	WindowStartAt time.Time  `json:"window_start_at"`
	WindowEndAt   time.Time  `json:"window_end_at"`
	CompletedAt   *time.Time `json:"completed_at,omitempty"`
	RedeemedAt    *time.Time `json:"redeemed_at,omitempty"`
}

// TierSnapshot captures the current loyalty tier and underlying stats.
type TierSnapshot struct {
	UserID         uuid.UUID `json:"user_id"`
	Tier           string    `json:"tier"`
	RatingAverage  float64   `json:"rating_avg"`
	AcceptanceRate float64   `json:"acceptance_rate"`
	OnTimeRate     float64   `json:"on_time_rate"`
	Deliveries30D  int       `json:"deliveries_30d"`
	ComputedAt     time.Time `json:"computed_at"`
}

// Referral represents a runner invite.
type Referral struct {
	ID                         uuid.UUID  `json:"id"`
	ReferrerUserID             uuid.UUID  `json:"referrer_user_id"`
	RefereeUserID              *uuid.UUID `json:"referee_user_id,omitempty"`
	Code                       string     `json:"code"`
	RefereeCompletedDeliveries int        `json:"referee_completed_deliveries"`
	PayoutEligibleAt           *time.Time `json:"payout_eligible_at,omitempty"`
	PayoutAmountMYR            string     `json:"payout_amount_myr"`
	PayoutDisbursedAt          *time.Time `json:"payout_disbursed_at,omitempty"`
}

// Redemption is a pending or disbursed reward payout.
type Redemption struct {
	ID          uuid.UUID  `json:"id"`
	UserID      uuid.UUID  `json:"user_id"`
	QuestID     *uuid.UUID `json:"quest_id,omitempty"`
	ReferralID  *uuid.UUID `json:"referral_id,omitempty"`
	AmountMYR   string     `json:"amount_myr"`
	Status      string     `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	DisbursedAt *time.Time `json:"disbursed_at,omitempty"`
}

// CreateReferralCodeResponse returns an idempotent referral code.
type CreateReferralCodeResponse struct {
	Code string `json:"code"`
}

// RegisterRefereeSignupRequest links a new signup to a referral code.
type RegisterRefereeSignupRequest struct {
	Code          string    `json:"code" binding:"required"`
	RefereeUserID uuid.UUID `json:"referee_user_id" binding:"required"`
}

// RedeemQuestRequest redeems a completed quest.
type RedeemQuestRequest struct {
	QuestID uuid.UUID `json:"quest_id" binding:"required"`
}
