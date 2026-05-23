package events

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/google/uuid"
)

func Test_TopicInventoryEvents_Constant(t *testing.T) {
	if TopicInventoryEvents != "inventory.events" {
		t.Errorf("TopicInventoryEvents got %q, want %q", TopicInventoryEvents, "inventory.events")
	}
}

func Test_InventoryAdjustedEvent_JSONRoundTrip(t *testing.T) {
	orig := InventoryAdjustedEvent{
		MovementID:   uuid.New(),
		ShopID:       uuid.New(),
		ProductID:    uuid.New(),
		SKU:          "ABC-001",
		Reason:       "sale",
		ChangeAmount: -2,
		BalanceAfter: 8,
		ActorUserID:  uuid.New(),
		NewVersion:   5,
		OccurredAt:   time.Now().UTC().Truncate(time.Second),
	}
	data, err := json.Marshal(orig)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}
	var decoded InventoryAdjustedEvent
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if decoded != orig {
		t.Errorf("round-trip mismatch: got %+v, want %+v", decoded, orig)
	}
}

func Test_InventoryBelowThresholdEvent_JSONRoundTrip(t *testing.T) {
	orig := InventoryBelowThresholdEvent{
		ShopID:      uuid.New(),
		ProductID:   uuid.New(),
		SKU:         "ABC-001",
		Name:        "Dog Food",
		StockOnHand: 2,
		Threshold:   5,
		OccurredAt:  time.Now().UTC().Truncate(time.Second),
	}
	data, err := json.Marshal(orig)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}
	var decoded InventoryBelowThresholdEvent
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if decoded != orig {
		t.Errorf("round-trip mismatch: got %+v, want %+v", decoded, orig)
	}
}
