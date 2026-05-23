package shop

import (
	"encoding/json"
	"strings"
	"testing"
)

func Test_StaffRole_RoundTripAndValidate(t *testing.T) {
	roles := []string{StaffRoleOwner, StaffRoleManager, StaffRoleStaff}
	for _, role := range roles {
		t.Run(role, func(t *testing.T) {
			req := InviteStaffRequest{Email: "x@example.com", Role: role}
			data, err := json.Marshal(req)
			if err != nil {
				t.Fatalf("marshal failed: %v", err)
			}
			var decoded InviteStaffRequest
			if err := json.Unmarshal(data, &decoded); err != nil {
				t.Fatalf("unmarshal failed: %v", err)
			}
			if decoded.Role != role {
				t.Errorf("round-trip role mismatch: got %q, want %q", decoded.Role, role)
			}
			if err := decoded.Validate(); err != nil {
				t.Errorf("Validate rejected valid round-trip: %v", err)
			}
		})
	}
}

func Test_InviteStaffRequest_RejectsMissingContact(t *testing.T) {
	req := InviteStaffRequest{Role: StaffRoleStaff}
	err := req.Validate()
	if err == nil {
		t.Fatal("expected error when both email and phone are empty")
	}
	if !strings.Contains(err.Error(), "email") || !strings.Contains(err.Error(), "phone") {
		t.Errorf("error should mention both email and phone, got: %v", err)
	}
}

func Test_InviteStaffRequest_RejectsUnknownRole(t *testing.T) {
	req := InviteStaffRequest{Email: "x@example.com", Role: "vendor"}
	err := req.Validate()
	if err == nil {
		t.Fatal("expected error for unknown role")
	}
	if !strings.Contains(err.Error(), "role") {
		t.Errorf("error should mention role, got: %v", err)
	}
}

func Test_AcceptStaffInviteRequest_RejectsEmptyToken(t *testing.T) {
	req := AcceptStaffInviteRequest{}
	if err := req.Validate(); err == nil {
		t.Fatal("expected error for empty token")
	}
}
