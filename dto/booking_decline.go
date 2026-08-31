package dto

import "fmt"

// Decline reason constants for DeclineBookingRequest.
const (
	DeclineReasonTooFar          = "too_far"
	DeclineReasonCannotTransport = "cannot_transport"
	DeclineReasonAlreadyBusy     = "already_busy"
	DeclineReasonPickupIssue     = "pickup_issue"
	DeclineReasonOther           = "other"
)

// DeclineBookingRequest is the request body sent when a runner declines a booking.
// Reason is required and must be one of the five known decline reason values.
type DeclineBookingRequest struct {
	Reason string `json:"reason"`
}

// Validate returns nil if the request is valid, or a descriptive error if the
// reason field is missing or not one of the five known values.
func (r DeclineBookingRequest) Validate() error {
	switch r.Reason {
	case DeclineReasonTooFar, DeclineReasonCannotTransport, DeclineReasonAlreadyBusy,
		DeclineReasonPickupIssue, DeclineReasonOther:
		// valid
	default:
		return fmt.Errorf("reason must be one of: too_far, cannot_transport, already_busy, pickup_issue, other")
	}
	return nil
}
