package dto

import "fmt"

// CashOutRequest is the request body for a runner cash-out operation.
// AmountMyrCents is expressed in minor units (cents) to avoid floating-point
// representation issues. The server enforces the upper bound against the
// runner's actual balance; this DTO only rejects non-positive amounts.
type CashOutRequest struct {
	AmountMyrCents int64  `json:"amountMyrCents"`
	DestinationID  string `json:"destinationId"`
}

// Validate returns nil if the request is valid, or a descriptive error if any
// required field is missing or invalid.
func (r CashOutRequest) Validate() error {
	if r.AmountMyrCents <= 0 {
		return fmt.Errorf("amountMyrCents must be greater than 0")
	}
	if r.DestinationID == "" {
		return fmt.Errorf("destinationId is required")
	}
	return nil
}

// CashOutResponse is the response body for a successful cash-out request.
// CashOutID is the unique identifier for the cash-out transaction, and
// EtaMinutes is the estimated processing time in minutes shown to the runner.
type CashOutResponse struct {
	CashOutID  string `json:"cashOutId"`
	EtaMinutes int    `json:"etaMinutes"`
}
