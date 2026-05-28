package events

import (
	"time"

	"github.com/google/uuid"
)

// Kafka topic for zone events.
const TopicZonesEvents = "zones.events"

// Zone event types.
const (
	ZoneSurgeChanged = "zones.surge_changed"
)

// ZoneSurgeChangedEvent is published when a zone multiplier transitions.
type ZoneSurgeChangedEvent struct {
	ZoneID             uuid.UUID `json:"zone_id"`
	ZoneCode           string    `json:"zone_code"`
	PreviousMultiplier float64   `json:"previous_multiplier"`
	NewMultiplier      float64   `json:"new_multiplier"`
	OpenJobs           int       `json:"open_jobs"`
	ActiveRunners      int       `json:"active_runners"`
	OccurredAt         time.Time `json:"occurred_at"`
}
