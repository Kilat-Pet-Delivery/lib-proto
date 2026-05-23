package events

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/google/uuid"
)

func Test_TopicShopEvents_Constant(t *testing.T) {
	if TopicShopEvents != "shop.events" {
		t.Errorf("TopicShopEvents got %q, want %q", TopicShopEvents, "shop.events")
	}
}

func Test_ShopEventTypes_AreNamespaced(t *testing.T) {
	cases := []string{
		ShopCreated,
		ShopStatusChanged,
		ShopStaffInvited,
		ShopStaffAccepted,
		ShopStaffRemoved,
		SalesAggregateRefreshed,
	}
	for _, ev := range cases {
		t.Run(ev, func(t *testing.T) {
			if ev == "" {
				t.Fatal("event type should not be empty")
			}
			if got := ev[:5]; got != "shop." {
				t.Errorf("event %q should be namespaced under 'shop.', got prefix %q", ev, got)
			}
		})
	}
}

func Test_ShopCreatedEvent_JSONRoundTrip(t *testing.T) {
	now := time.Now().UTC().Truncate(time.Second)
	orig := ShopCreatedEvent{
		ShopID:      uuid.New(),
		OwnerUserID: uuid.New(),
		Name:        "Pet Place",
		Slug:        "pet-place",
		Category:    "vet",
		OccurredAt:  now,
	}

	data, err := json.Marshal(orig)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}
	var decoded ShopCreatedEvent
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if decoded != orig {
		t.Errorf("round-trip mismatch: got %+v, want %+v", decoded, orig)
	}
}

func Test_ShopStatusChangedEvent_AutoCloseAtOmitsWhenNil(t *testing.T) {
	ev := ShopStatusChangedEvent{
		ShopID:      uuid.New(),
		OwnerUserID: uuid.New(),
		OldStatus:   "open",
		NewStatus:   "closed",
		OccurredAt:  time.Now().UTC(),
	}
	data, err := json.Marshal(ev)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}
	if got := string(data); contains(got, "auto_close_at") {
		t.Errorf("auto_close_at should be omitted when nil, got: %s", got)
	}
}

func Test_SalesAggregateRefreshedEvent_JSONRoundTrip(t *testing.T) {
	orig := SalesAggregateRefreshedEvent{
		ShopID:          uuid.New(),
		Period:          "day",
		PeriodStart:     time.Date(2026, 5, 23, 0, 0, 0, 0, time.UTC),
		PeriodEnd:       time.Date(2026, 5, 24, 0, 0, 0, 0, time.UTC),
		OrderCount:      10,
		GrossSalesCents: 100000,
		NetSalesCents:   90000,
		Currency:        "MYR",
		OccurredAt:      time.Now().UTC().Truncate(time.Second),
	}
	data, err := json.Marshal(orig)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}
	var decoded SalesAggregateRefreshedEvent
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if decoded != orig {
		t.Errorf("round-trip mismatch: got %+v, want %+v", decoded, orig)
	}
}

func contains(s, sub string) bool {
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}
