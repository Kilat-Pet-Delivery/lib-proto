package dto

import (
	"strings"
	"testing"
)

func TestDeclineBookingRequest_AcceptsKnownReasons(t *testing.T) {
	for _, reason := range []string{
		DeclineReasonTooFar,
		DeclineReasonCannotTransport,
		DeclineReasonAlreadyBusy,
		DeclineReasonPickupIssue,
		DeclineReasonOther,
	} {
		t.Run(reason, func(t *testing.T) {
			r := DeclineBookingRequest{Reason: reason}
			if err := r.Validate(); err != nil {
				t.Errorf("expected nil error for reason %q, got: %v", reason, err)
			}
		})
	}
}

func TestDeclineBookingRequest_RejectsUnknownReason(t *testing.T) {
	t.Run("empty reason", func(t *testing.T) {
		r := DeclineBookingRequest{Reason: ""}
		err := r.Validate()
		if err == nil {
			t.Fatal("expected error for empty reason, got nil")
		}
		if !strings.Contains(err.Error(), "reason") {
			t.Errorf("expected error to mention 'reason', got: %v", err)
		}
	})

	t.Run("made_up_reason", func(t *testing.T) {
		r := DeclineBookingRequest{Reason: "made_up_reason"}
		err := r.Validate()
		if err == nil {
			t.Fatal("expected error for unknown reason, got nil")
		}
		if !strings.Contains(err.Error(), "reason") {
			t.Errorf("expected error to mention 'reason', got: %v", err)
		}
	})
}
