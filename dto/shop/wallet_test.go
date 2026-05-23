package shop

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/google/uuid"
)

func Test_WithdrawalStatus_AllRecognized(t *testing.T) {
	statuses := []string{
		WithdrawalStatusPending,
		WithdrawalStatusProcessing,
		WithdrawalStatusPaid,
		WithdrawalStatusFailed,
	}
	for _, s := range statuses {
		if !IsValidWithdrawalStatus(s) {
			t.Errorf("expected %q to be a recognized withdrawal status", s)
		}
	}
	if IsValidWithdrawalStatus("cancelled") {
		t.Error("expected 'cancelled' to be unrecognized")
	}
}

func Test_WithdrawRequest_Validate(t *testing.T) {
	tests := []struct {
		name    string
		req     WithdrawRequest
		wantErr string
	}{
		{
			name: "happy path",
			req:  WithdrawRequest{AmountCents: 10000, DestinationID: "dest-1"},
		},
		{
			name:    "zero amount rejected",
			req:     WithdrawRequest{AmountCents: 0, DestinationID: "dest-1"},
			wantErr: "amount_cents",
		},
		{
			name:    "negative amount rejected",
			req:     WithdrawRequest{AmountCents: -100, DestinationID: "dest-1"},
			wantErr: "amount_cents",
		},
		{
			name:    "missing destination",
			req:     WithdrawRequest{AmountCents: 10000},
			wantErr: "destination_id",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.req.Validate()
			if tt.wantErr != "" {
				if err == nil {
					t.Fatalf("expected error mentioning %q, got nil", tt.wantErr)
				}
				if !strings.Contains(err.Error(), tt.wantErr) {
					t.Errorf("error should mention %q, got: %v", tt.wantErr, err)
				}
				return
			}
			if err != nil {
				t.Errorf("expected nil for valid request, got: %v", err)
			}
		})
	}
}

func Test_ShopWalletLedgerEntryDTO_OmitsNilBookingAndWithdrawal(t *testing.T) {
	entry := ShopWalletLedgerEntryDTO{
		ID:           uuid.New(),
		ShopID:       uuid.New(),
		EntryType:    "credit",
		ChangeCents:  1000,
		BalanceAfter: 5000,
		Currency:     "MYR",
	}
	data, err := json.Marshal(entry)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}
	out := string(data)
	if strings.Contains(out, "booking_id") {
		t.Errorf("booking_id should be omitted when nil, got: %s", out)
	}
	if strings.Contains(out, "withdrawal_id") {
		t.Errorf("withdrawal_id should be omitted when nil, got: %s", out)
	}
}
