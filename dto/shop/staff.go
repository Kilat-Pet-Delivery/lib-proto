package shop

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// StaffRole enumerates the permission tiers for shop staff members.
const (
	StaffRoleOwner   = "owner"
	StaffRoleManager = "manager"
	StaffRoleStaff   = "staff"
)

// ShopStaffDTO is the response representation of a shop staff membership.
type ShopStaffDTO struct {
	ID         uuid.UUID  `json:"id"`
	ShopID     uuid.UUID  `json:"shop_id"`
	UserID     *uuid.UUID `json:"user_id,omitempty"`
	Email      string     `json:"email,omitempty"`
	Phone      string     `json:"phone,omitempty"`
	Role       string     `json:"role"`
	Status     string     `json:"status"`
	InvitedAt  time.Time  `json:"invited_at"`
	AcceptedAt *time.Time `json:"accepted_at,omitempty"`
	RemovedAt  *time.Time `json:"removed_at,omitempty"`
}

// InviteStaffRequest is the request body sent by an owner or manager to
// invite a new staff member by phone or email.
type InviteStaffRequest struct {
	Email string `json:"email,omitempty"`
	Phone string `json:"phone,omitempty"`
	Role  string `json:"role" binding:"required"`
}

// Validate returns nil if the request supplies at least one contact field
// (email or phone) and a known role value.
func (r InviteStaffRequest) Validate() error {
	if r.Email == "" && r.Phone == "" {
		return fmt.Errorf("email or phone is required")
	}
	if !IsValidStaffRole(r.Role) {
		return fmt.Errorf("role must be one of: owner, manager, staff")
	}
	return nil
}

// AcceptStaffInviteRequest is the request body for the
// POST /shops/staff/invites/{token}/accept endpoint. The caller's identity
// is taken from the auth context; the token is the single-use invite handle.
type AcceptStaffInviteRequest struct {
	Token string `json:"token" binding:"required"`
}

// Validate returns nil if the token is non-empty.
func (r AcceptStaffInviteRequest) Validate() error {
	if r.Token == "" {
		return fmt.Errorf("token is required")
	}
	return nil
}

// IsValidStaffRole returns true when r is one of the three staff role values.
func IsValidStaffRole(r string) bool {
	switch r {
	case StaffRoleOwner, StaffRoleManager, StaffRoleStaff:
		return true
	default:
		return false
	}
}
