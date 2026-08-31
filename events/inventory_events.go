package events

import (
	"time"

	"github.com/google/uuid"
)

// Kafka topic for inventory events. Partition key is product_id so per-product
// ordering is preserved (important for the optimistic-lock invariant).
const TopicInventoryEvents = "inventory.events"

// Inventory event types.
const (
	InventoryAdjusted       = "inventory.adjusted"
	InventoryBelowThreshold = "inventory.below_threshold"
)

// InventoryAdjustedEvent is published after a successful stock mutation
// (sale, restock, manual adjustment, return, damage).
type InventoryAdjustedEvent struct {
	MovementID   uuid.UUID `json:"movement_id"`
	ShopID       uuid.UUID `json:"shop_id"`
	ProductID    uuid.UUID `json:"product_id"`
	SKU          string    `json:"sku"`
	Reason       string    `json:"reason"`
	ChangeAmount int64     `json:"change_amount"`
	BalanceAfter int64     `json:"balance_after"`
	ActorUserID  uuid.UUID `json:"actor_user_id"`
	NewVersion   int64     `json:"new_version"`
	OccurredAt   time.Time `json:"occurred_at"`
}

// InventoryBelowThresholdEvent is published when post-adjustment stock falls
// at or below the configured low-stock threshold. Consumed by
// service-notification for the merchant's low-stock alert push.
type InventoryBelowThresholdEvent struct {
	ShopID      uuid.UUID `json:"shop_id"`
	ProductID   uuid.UUID `json:"product_id"`
	SKU         string    `json:"sku"`
	Name        string    `json:"name"`
	StockOnHand int64     `json:"stock_on_hand"`
	Threshold   int64     `json:"threshold"`
	OccurredAt  time.Time `json:"occurred_at"`
}
