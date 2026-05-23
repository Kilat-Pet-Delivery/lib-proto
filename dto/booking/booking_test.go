package booking

import (
	"encoding/json"
	"strings"
	"testing"
)

func Test_BookingStatus_NewValuesRoundTrip(t *testing.T) {
	// Specifically validates the three Plan C-introduced substates.
	newStatuses := []BookingStatus{
		BookingStatusAcceptedByShop,
		BookingStatusPreparing,
		BookingStatusReadyForPickup,
	}

	for _, status := range newStatuses {
		t.Run(string(status), func(t *testing.T) {
			data, err := json.Marshal(status)
			if err != nil {
				t.Fatalf("marshal failed: %v", err)
			}
			var decoded BookingStatus
			if err := json.Unmarshal(data, &decoded); err != nil {
				t.Fatalf("unmarshal failed: %v", err)
			}
			if decoded != status {
				t.Errorf("round-trip mismatch: got %q, want %q", decoded, status)
			}
			if !decoded.IsValid() {
				t.Errorf("decoded value %q failed IsValid", decoded)
			}
			parsed, err := ParseBookingStatus(string(decoded))
			if err != nil {
				t.Errorf("ParseBookingStatus rejected its own round-tripped value: %v", err)
			}
			if parsed != status {
				t.Errorf("ParseBookingStatus mismatch: got %q, want %q", parsed, status)
			}
		})
	}
}

func Test_BookingStatus_LegacyValuesStillValid(t *testing.T) {
	legacy := []BookingStatus{
		BookingStatusRequested,
		BookingStatusAccepted,
		BookingStatusInProgress,
		BookingStatusDelivered,
		BookingStatusCompleted,
		BookingStatusCancelled,
	}
	for _, status := range legacy {
		if !status.IsValid() {
			t.Errorf("legacy status %q should remain valid", status)
		}
	}
}

func Test_BookingStatus_RejectsUnknown(t *testing.T) {
	if BookingStatus("on_hold").IsValid() {
		t.Error("unknown status 'on_hold' should not be valid")
	}
	if _, err := ParseBookingStatus("on_hold"); err == nil {
		t.Error("ParseBookingStatus should reject 'on_hold'")
	}
}

func Test_BookingDTO_QRTokenOmitsWhenNil(t *testing.T) {
	dto := BookingDTO{
		ID:            "b-1",
		BookingNumber: "BK-001",
		OwnerID:       "u-1",
		Status:        BookingStatusRequested,
		// QRPickupToken intentionally left nil
	}

	data, err := json.Marshal(dto)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}
	if strings.Contains(string(data), "qr_pickup_token") {
		t.Errorf("qr_pickup_token should be omitted when nil, got: %s", string(data))
	}
}

func Test_BookingDTO_QRTokenIncludedWhenSet(t *testing.T) {
	token := "qr-xyz-123"
	dto := BookingDTO{
		ID:            "b-1",
		BookingNumber: "BK-001",
		OwnerID:       "u-1",
		Status:        BookingStatusReadyForPickup,
		QRPickupToken: &token,
	}

	data, err := json.Marshal(dto)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}
	if !strings.Contains(string(data), `"qr_pickup_token":"qr-xyz-123"`) {
		t.Errorf("qr_pickup_token should be included, got: %s", string(data))
	}

	var decoded BookingDTO
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if decoded.QRPickupToken == nil || *decoded.QRPickupToken != token {
		t.Errorf("qr_pickup_token round-trip lost: %v", decoded.QRPickupToken)
	}
}

func Test_BookingDTO_JSONRoundTrip(t *testing.T) {
	runnerID := "r-1"
	shopID := "s-1"
	token := "qr-1"
	orig := BookingDTO{
		ID:            "b-1",
		BookingNumber: "BK-001",
		OwnerID:       "u-1",
		RunnerID:      &runnerID,
		ShopID:        &shopID,
		Status:        BookingStatusAcceptedByShop,
		QRPickupToken: &token,
		Notes:         "ring doorbell twice",
	}

	data, err := json.Marshal(orig)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}
	var decoded BookingDTO
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if decoded.ID != orig.ID || decoded.Status != orig.Status {
		t.Errorf("round-trip mismatch: %+v vs %+v", decoded, orig)
	}
	if decoded.RunnerID == nil || *decoded.RunnerID != runnerID {
		t.Errorf("runner_id round-trip lost: %v", decoded.RunnerID)
	}
}
