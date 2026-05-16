package dto

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestCashOutRequest_JSONRoundTrip(t *testing.T) {
	original := CashOutRequest{
		AmountMyrCents: 64250,
		DestinationID:  "dest-abc",
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}

	var decoded CashOutRequest
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}

	if decoded != original {
		t.Errorf("round-trip mismatch: got %+v, want %+v", decoded, original)
	}
}

func TestCashOutRequest_RejectsNegativeAmount(t *testing.T) {
	t.Run("zero amount rejected", func(t *testing.T) {
		req := CashOutRequest{AmountMyrCents: 0, DestinationID: "dest-abc"}
		err := req.Validate()
		if err == nil {
			t.Fatal("expected error for zero amountMyrCents, got nil")
		}
		if !strings.Contains(err.Error(), "amountMyrCents") {
			t.Errorf("error should mention 'amountMyrCents', got: %v", err)
		}
	})

	t.Run("negative amount rejected", func(t *testing.T) {
		req := CashOutRequest{AmountMyrCents: -100, DestinationID: "dest-abc"}
		err := req.Validate()
		if err == nil {
			t.Fatal("expected error for negative amountMyrCents, got nil")
		}
		if !strings.Contains(err.Error(), "amountMyrCents") {
			t.Errorf("error should mention 'amountMyrCents', got: %v", err)
		}
	})

	t.Run("empty destinationId rejected", func(t *testing.T) {
		req := CashOutRequest{AmountMyrCents: 100, DestinationID: ""}
		err := req.Validate()
		if err == nil {
			t.Fatal("expected error for empty destinationId, got nil")
		}
		if !strings.Contains(err.Error(), "destinationId") {
			t.Errorf("error should mention 'destinationId', got: %v", err)
		}
	})

	t.Run("valid request passes (boundary: 1 cent)", func(t *testing.T) {
		req := CashOutRequest{AmountMyrCents: 1, DestinationID: "dest-abc"}
		if err := req.Validate(); err != nil {
			t.Errorf("expected nil for valid request, got: %v", err)
		}
	})
}

func TestCashOutResponse_JSONRoundTrip(t *testing.T) {
	original := CashOutResponse{
		CashOutID:  "co-abc-123",
		EtaMinutes: 30,
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}

	var decoded CashOutResponse
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}

	if decoded != original {
		t.Errorf("round-trip mismatch: got %+v, want %+v", decoded, original)
	}
}
