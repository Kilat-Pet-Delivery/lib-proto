package dto

import (
	"time"

	"github.com/google/uuid"
)

// ZonePolygon is a GeoJSON-style polygon.
type ZonePolygon struct {
	Type        string        `json:"type"`
	Coordinates [][][]float64 `json:"coordinates"`
}

// SurgeMultiplier captures the latest demand snapshot for a zone.
type SurgeMultiplier struct {
	ZoneID        uuid.UUID `json:"zone_id"`
	OpenJobs      int       `json:"open_jobs"`
	ActiveRunners int       `json:"active_runners"`
	Multiplier    float64   `json:"multiplier"`
	ComputedAt    time.Time `json:"computed_at"`
}

// Zone represents an active delivery zone.
type Zone struct {
	ID              uuid.UUID        `json:"id"`
	Code            string           `json:"code"`
	Label           string           `json:"label"`
	Polygon         ZonePolygon      `json:"polygon"`
	Active          bool             `json:"active"`
	SurgeMultiplier *SurgeMultiplier `json:"surge_multiplier,omitempty"`
	CreatedAt       time.Time        `json:"created_at"`
}

// ListZonesResponse returns active zones and their demand snapshots.
type ListZonesResponse struct {
	Zones []Zone `json:"zones"`
}

// ZoneAtPointRequest looks up a zone by coordinate.
type ZoneAtPointRequest struct {
	Latitude  float64 `json:"lat" binding:"required"`
	Longitude float64 `json:"lon" binding:"required"`
}

// ZoneAtPointResponse returns the containing zone, if any.
type ZoneAtPointResponse struct {
	Zone *Zone `json:"zone,omitempty"`
}
