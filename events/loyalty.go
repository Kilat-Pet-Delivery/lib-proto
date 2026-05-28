package events

import (
	"time"

	"github.com/google/uuid"
)

// Kafka topic for loyalty events.
const TopicLoyaltyEvents = "loyalty.events"

// Loyalty event types.
const (
	QuestCompleted    = "loyalty.quest_completed"
	TierPromoted      = "loyalty.tier_promoted"
	RedemptionCreated = "loyalty.redemption_created"
	ReferralPayoutDue = "loyalty.referral_payout_due"
)

// QuestCompletedEvent is published when a runner completes a quest threshold.
type QuestCompletedEvent struct {
	QuestID      uuid.UUID `json:"quest_id"`
	QuestCode    string    `json:"quest_code"`
	UserID       uuid.UUID `json:"user_id"`
	RewardAmount int64     `json:"reward_amount_cents"`
	Currency     string    `json:"currency"`
	OccurredAt   time.Time `json:"occurred_at"`
}

// TierPromotedEvent is published when a runner's loyalty tier increases.
type TierPromotedEvent struct {
	UserID       uuid.UUID `json:"user_id"`
	PreviousTier string    `json:"previous_tier"`
	NewTier      string    `json:"new_tier"`
	OccurredAt   time.Time `json:"occurred_at"`
}

// RedemptionCreatedEvent is published when a quest or referral reward is redeemed.
type RedemptionCreatedEvent struct {
	RedemptionID uuid.UUID  `json:"redemption_id"`
	UserID       uuid.UUID  `json:"user_id"`
	QuestID      *uuid.UUID `json:"quest_id,omitempty"`
	ReferralID   *uuid.UUID `json:"referral_id,omitempty"`
	Amount       int64      `json:"amount_cents"`
	Currency     string     `json:"currency"`
	OccurredAt   time.Time  `json:"occurred_at"`
}

// ReferralPayoutDueEvent is published when a referral crosses the payout threshold.
type ReferralPayoutDueEvent struct {
	ReferralID     uuid.UUID `json:"referral_id"`
	ReferrerUserID uuid.UUID `json:"referrer_user_id"`
	RefereeUserID  uuid.UUID `json:"referee_user_id"`
	PayoutAmount   int64     `json:"payout_amount_cents"`
	Currency       string    `json:"currency"`
	EligibleAt     time.Time `json:"eligible_at"`
	OccurredAt     time.Time `json:"occurred_at"`
}
