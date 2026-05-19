package dto

import (
	"time"

	"github.com/google/uuid"
)

// Incident represents a support incident.
type Incident struct {
	ID             uuid.UUID  `json:"id"`
	Type           string     `json:"type"`
	Severity       string     `json:"severity"`
	BookingID      *uuid.UUID `json:"booking_id,omitempty"`
	ReporterUserID uuid.UUID  `json:"reporter_user_id"`
	AssigneeUserID *uuid.UUID `json:"assignee_user_id,omitempty"`
	Status         string     `json:"status"`
	Notes          string     `json:"notes,omitempty"`
	BreachedAt     *time.Time `json:"breached_at,omitempty"`
	CreatedAt      time.Time  `json:"created_at"`
	AssignedAt     *time.Time `json:"assigned_at,omitempty"`
	ResolvedAt     *time.Time `json:"resolved_at,omitempty"`
}

// IncidentEvent is an audit-log entry for an incident mutation.
type IncidentEvent struct {
	ID          uuid.UUID `json:"id"`
	IncidentID  uuid.UUID `json:"incident_id"`
	EventType   string    `json:"event_type"`
	ActorUserID uuid.UUID `json:"actor_user_id"`
	Payload     any       `json:"payload,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}

// CreateIncidentRequest creates a new incident.
type CreateIncidentRequest struct {
	Type      string     `json:"type" binding:"required"`
	Severity  string     `json:"severity" binding:"required"`
	BookingID *uuid.UUID `json:"booking_id,omitempty"`
	Notes     string     `json:"notes,omitempty"`
}

// AssignIncidentRequest assigns an incident to a support agent.
type AssignIncidentRequest struct {
	AssigneeUserID uuid.UUID `json:"assignee_user_id" binding:"required"`
}

// TransitionIncidentRequest moves an incident through the lifecycle.
type TransitionIncidentRequest struct {
	ToStatus string `json:"to_status" binding:"required"`
	Notes    string `json:"notes,omitempty"`
}

// ResolveIncidentRequest resolves an incident.
type ResolveIncidentRequest struct {
	ResolutionNotes string `json:"resolution_notes" binding:"required"`
}

// ListIncidentsRequest contains incident filters.
type ListIncidentsRequest struct {
	Status   string `json:"status,omitempty"`
	Type     string `json:"type,omitempty"`
	Assignee string `json:"assignee,omitempty"`
	Breached *bool  `json:"breached,omitempty"`
	Limit    int    `json:"limit"`
	Offset   int    `json:"offset"`
}

// ListIncidentsResponse returns paginated incidents.
type ListIncidentsResponse struct {
	Incidents []Incident `json:"incidents"`
	Total     int        `json:"total"`
}
