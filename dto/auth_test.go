package dto

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestForgotPasswordRequest_JSONRoundTrip(t *testing.T) {
	original := ForgotPasswordRequest{
		Email: "user@example.com",
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}

	var decoded ForgotPasswordRequest
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}

	if !reflect.DeepEqual(original, decoded) {
		t.Errorf("round-trip mismatch: got %+v, want %+v", decoded, original)
	}
}

func TestResetPasswordRequest_JSONRoundTrip(t *testing.T) {
	original := ResetPasswordRequest{
		Token:       "some-reset-token-abc123",
		NewPassword: "s3cur3P@ssw0rd",
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}

	var decoded ResetPasswordRequest
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}

	if !reflect.DeepEqual(original, decoded) {
		t.Errorf("round-trip mismatch: got %+v, want %+v", decoded, original)
	}
}
