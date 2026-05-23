package events

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/google/uuid"
)

func Test_WithdrawalEventTypes_Constants(t *testing.T) {
	cases := map[string]string{
		"WithdrawalRequested": WithdrawalRequested,
		"WithdrawalPaid":      WithdrawalPaid,
		"WithdrawalFailed":    WithdrawalFailed,
	}
	want := map[string]string{
		"WithdrawalRequested": "payment.withdrawal_requested",
		"WithdrawalPaid":      "payment.withdrawal_paid",
		"WithdrawalFailed":    "payment.withdrawal_failed",
	}
	for name, got := range cases {
		if got != want[name] {
			t.Errorf("%s = %q, want %q", name, got, want[name])
		}
	}
}

func Test_WithdrawalRequestedEvent_JSONRoundTrip(t *testing.T) {
	orig := WithdrawalRequestedEvent{
		WithdrawalID:  uuid.New(),
		ShopID:        uuid.New(),
		RequestedBy:   uuid.New(),
		AmountCents:   10000,
		Currency:      "MYR",
		DestinationID: "dest-1",
		OccurredAt:    time.Now().UTC().Truncate(time.Second),
	}
	data, err := json.Marshal(orig)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}
	var decoded WithdrawalRequestedEvent
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if decoded != orig {
		t.Errorf("round-trip mismatch: got %+v, want %+v", decoded, orig)
	}
}

func Test_WithdrawalPaidEvent_ProcessorRefOmitsWhenEmpty(t *testing.T) {
	orig := WithdrawalPaidEvent{
		WithdrawalID: uuid.New(),
		ShopID:       uuid.New(),
		AmountCents:  10000,
		Currency:     "MYR",
		OccurredAt:   time.Now().UTC().Truncate(time.Second),
		// ProcessorRef intentionally empty
	}
	data, err := json.Marshal(orig)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}
	if containsString(string(data), "processor_ref") {
		t.Errorf("processor_ref should be omitted when empty, got: %s", string(data))
	}
}

func Test_WithdrawalFailedEvent_JSONRoundTrip(t *testing.T) {
	orig := WithdrawalFailedEvent{
		WithdrawalID: uuid.New(),
		ShopID:       uuid.New(),
		AmountCents:  10000,
		Currency:     "MYR",
		Reason:       "insufficient_funds",
		OccurredAt:   time.Now().UTC().Truncate(time.Second),
	}
	data, err := json.Marshal(orig)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}
	var decoded WithdrawalFailedEvent
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if decoded != orig {
		t.Errorf("round-trip mismatch: got %+v, want %+v", decoded, orig)
	}
}

func containsString(s, sub string) bool {
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}
