package events

import (
	"time"

	"github.com/google/uuid"
)

// Kafka topic for tracking events.
const TopicTrackingEvents = "tracking.events"

// Tracking event types.
const (
	TrackingStarted   = "tracking.started"
	TrackingUpdated   = "tracking.updated"
	TrackingCompleted = "tracking.completed"
)

// TrackingStartedEvent is published when a trip tracking session begins.
type TrackingStartedEvent struct {
	TrackID    uuid.UUID `json:"track_id"`
	BookingID  uuid.UUID `json:"booking_id"`
	RunnerID   uuid.UUID `json:"runner_id"`
	StartedAt  time.Time `json:"started_at"`
	OccurredAt time.Time `json:"occurred_at"`
}

// TrackingUpdatedEvent is published when a new waypoint is added.
type TrackingUpdatedEvent struct {
	TrackID    uuid.UUID `json:"track_id"`
	BookingID  uuid.UUID `json:"booking_id"`
	RunnerID   uuid.UUID `json:"runner_id"`
	Latitude   float64   `json:"latitude"`
	Longitude  float64   `json:"longitude"`
	Speed      float64   `json:"speed_kmh"`
	OccurredAt time.Time `json:"occurred_at"`
}

// TrackingCompletedEvent is published when a trip tracking session ends.
type TrackingCompletedEvent struct {
	TrackID       uuid.UUID `json:"track_id"`
	BookingID     uuid.UUID `json:"booking_id"`
	RunnerID      uuid.UUID `json:"runner_id"`
	TotalDistance float64   `json:"total_distance_km"`
	CompletedAt   time.Time `json:"completed_at"`
	OccurredAt    time.Time `json:"occurred_at"`
}
