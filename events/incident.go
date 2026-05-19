package events

import (
	"time"

	"github.com/google/uuid"
)

// Kafka topic for incident events.
const TopicIncidentEvents = "incident.events"

// Incident event types.
const (
	IncidentCreated   = "incident.created"
	IncidentAssigned  = "incident.assigned"
	IncidentEscalated = "incident.escalated"
	IncidentResolved  = "incident.resolved"
)

// IncidentCreatedEvent is published when a runner/customer incident is created.
type IncidentCreatedEvent struct {
	IncidentID     uuid.UUID  `json:"incident_id"`
	Type           string     `json:"type"`
	Severity       string     `json:"severity"`
	BookingID      *uuid.UUID `json:"booking_id,omitempty"`
	ReporterUserID uuid.UUID  `json:"reporter_user_id"`
	AssigneeUserID *uuid.UUID `json:"assignee_user_id,omitempty"`
	OccurredAt     time.Time  `json:"occurred_at"`
}

// IncidentAssignedEvent is published when an incident is assigned to an agent.
type IncidentAssignedEvent struct {
	IncidentID       uuid.UUID  `json:"incident_id"`
	AssigneeUserID   uuid.UUID  `json:"assignee_user_id"`
	AssignedByUserID *uuid.UUID `json:"assigned_by_user_id,omitempty"`
	OccurredAt       time.Time  `json:"occurred_at"`
}

// IncidentEscalatedEvent is published when an SLA breach escalates an incident.
type IncidentEscalatedEvent struct {
	IncidentID         uuid.UUID  `json:"incident_id"`
	PreviousSeverity   string     `json:"previous_severity"`
	NewSeverity        string     `json:"new_severity"`
	PreviousAssigneeID *uuid.UUID `json:"previous_assignee_user_id,omitempty"`
	NewAssigneeID      *uuid.UUID `json:"new_assignee_user_id,omitempty"`
	Reason             string     `json:"reason"`
	OccurredAt         time.Time  `json:"occurred_at"`
}

// IncidentResolvedEvent is published when an incident is resolved.
type IncidentResolvedEvent struct {
	IncidentID       uuid.UUID `json:"incident_id"`
	ResolvedByUserID uuid.UUID `json:"resolved_by_user_id"`
	ResolutionNotes  string    `json:"resolution_notes"`
	OccurredAt       time.Time `json:"occurred_at"`
}
