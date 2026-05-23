package events

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/google/uuid"
)

func Test_BookingShopSubstates_Constants(t *testing.T) {
	cases := map[string]string{
		"BookingAcceptedByShop": BookingAcceptedByShop,
		"BookingPreparing":      BookingPreparing,
		"BookingReadyForPickup": BookingReadyForPickup,
	}
	want := map[string]string{
		"BookingAcceptedByShop": "booking.accepted_by_shop",
		"BookingPreparing":      "booking.preparing",
		"BookingReadyForPickup": "booking.ready_for_pickup",
	}
	for name, got := range cases {
		if got != want[name] {
			t.Errorf("%s = %q, want %q", name, got, want[name])
		}
	}
}

func Test_BookingAcceptedByShopEvent_JSONRoundTrip(t *testing.T) {
	orig := BookingAcceptedByShopEvent{
		BookingID:     uuid.New(),
		BookingNumber: "BK-001",
		ShopID:        uuid.New(),
		AcceptedBy:    uuid.New(),
		OccurredAt:    time.Now().UTC().Truncate(time.Second),
	}
	data, err := json.Marshal(orig)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}
	var decoded BookingAcceptedByShopEvent
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if decoded != orig {
		t.Errorf("round-trip mismatch: got %+v, want %+v", decoded, orig)
	}
}

func Test_BookingReadyForPickupEvent_CarriesToken(t *testing.T) {
	orig := BookingReadyForPickupEvent{
		BookingID:     uuid.New(),
		BookingNumber: "BK-001",
		ShopID:        uuid.New(),
		QRPickupToken: "qr-xyz",
		OccurredAt:    time.Now().UTC().Truncate(time.Second),
	}
	data, err := json.Marshal(orig)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}
	var decoded BookingReadyForPickupEvent
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if decoded.QRPickupToken != "qr-xyz" {
		t.Errorf("qr_pickup_token lost: got %q", decoded.QRPickupToken)
	}
}
