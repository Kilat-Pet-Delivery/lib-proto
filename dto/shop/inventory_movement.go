package shop

import (
	"time"

	"github.com/google/uuid"
)

// InventoryMovementReason enumerates the recognized reasons for an inventory
// adjustment. The set is intentionally closed; service-shop rejects unknown
// values at the API boundary.
const (
	InventoryMovementReasonSale       = "sale"
	InventoryMovementReasonRestock    = "restock"
	InventoryMovementReasonAdjustment = "adjustment"
	InventoryMovementReasonReturn     = "return"
	InventoryMovementReasonDamage     = "damage"
)

// InventoryMovementDTO is the response representation of a single inventory
// ledger entry. ChangeAmount is signed: positive for restocks/returns,
// negative for sales/damage.
type InventoryMovementDTO struct {
	ID            uuid.UUID `json:"id"`
	ProductID     uuid.UUID `json:"product_id"`
	ShopID        uuid.UUID `json:"shop_id"`
	Reason        string    `json:"reason"`
	ChangeAmount  int64     `json:"change_amount"`
	BalanceAfter  int64     `json:"balance_after"`
	Note          string    `json:"note,omitempty"`
	ActorUserID   uuid.UUID `json:"actor_user_id"`
	OccurredAt    time.Time `json:"occurred_at"`
}

// IsValidInventoryMovementReason returns true when r is one of the five
// known inventory movement reasons.
func IsValidInventoryMovementReason(r string) bool {
	switch r {
	case InventoryMovementReasonSale, InventoryMovementReasonRestock,
		InventoryMovementReasonAdjustment, InventoryMovementReasonReturn,
		InventoryMovementReasonDamage:
		return true
	default:
		return false
	}
}
