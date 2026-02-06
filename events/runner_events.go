package events

import (
	"time"

	"github.com/google/uuid"
)

// Kafka topic for runner events.
const TopicRunnerEvents = "runner.events"

// Runner event types.
const (
	RunnerOnline         = "runner.online"
	RunnerOffline        = "runner.offline"
	RunnerLocationUpdate = "runner.location_update"
	RunnerAssigned       = "runner.assigned"
)

// RunnerOnlineEvent is published when a runner goes active.
type RunnerOnlineEvent struct {
	RunnerID   uuid.UUID `json:"runner_id"`
	Latitude   float64   `json:"latitude"`
	Longitude  float64   `json:"longitude"`
	OccurredAt time.Time `json:"occurred_at"`
}

// RunnerOfflineEvent is published when a runner goes inactive.
type RunnerOfflineEvent struct {
	RunnerID   uuid.UUID `json:"runner_id"`
	OccurredAt time.Time `json:"occurred_at"`
}

// RunnerLocationUpdateEvent is published on each GPS ping from the runner.
type RunnerLocationUpdateEvent struct {
	RunnerID   uuid.UUID `json:"runner_id"`
	BookingID  *uuid.UUID `json:"booking_id,omitempty"`
	Latitude   float64   `json:"latitude"`
	Longitude  float64   `json:"longitude"`
	Speed      float64   `json:"speed_kmh"`
	Heading    float64   `json:"heading_degrees"`
	Timestamp  time.Time `json:"timestamp"`
	OccurredAt time.Time `json:"occurred_at"`
}

// RunnerAssignedEvent is published when a runner is assigned to a booking.
type RunnerAssignedEvent struct {
	RunnerID   uuid.UUID `json:"runner_id"`
	BookingID  uuid.UUID `json:"booking_id"`
	OccurredAt time.Time `json:"occurred_at"`
}
