package dto

import (
	"time"

	"github.com/google/uuid"
)

// UserDTO represents a user in API responses.
type UserDTO struct {
	ID         uuid.UUID `json:"id"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	FullName   string    `json:"full_name"`
	Role       string    `json:"role"`
	IsVerified bool      `json:"is_verified"`
	AvatarURL  string    `json:"avatar_url,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
}

// RunnerDTO represents a runner in API responses.
type RunnerDTO struct {
	ID             uuid.UUID      `json:"id"`
	UserID         uuid.UUID      `json:"user_id"`
	FullName       string         `json:"full_name"`
	Phone          string         `json:"phone"`
	VehicleType    string         `json:"vehicle_type"`
	VehiclePlate   string         `json:"vehicle_plate"`
	VehicleModel   string         `json:"vehicle_model"`
	AirConditioned bool           `json:"air_conditioned"`
	SessionStatus  string         `json:"session_status"`
	Rating         float64        `json:"rating"`
	TotalTrips     int            `json:"total_trips"`
	CrateSpecs     []CrateSpecDTO `json:"crate_specs,omitempty"`
	DistanceKm     *float64       `json:"distance_km,omitempty"`
	CreatedAt      time.Time      `json:"created_at"`
}

// CrateSpecDTO represents a crate specification.
type CrateSpecDTO struct {
	ID                    uuid.UUID `json:"id"`
	Size                  string    `json:"size"`
	PetTypes              []string  `json:"pet_types"`
	MaxWeightKg           float64   `json:"max_weight_kg"`
	WidthCm               float64   `json:"width_cm"`
	HeightCm              float64   `json:"height_cm"`
	DepthCm               float64   `json:"depth_cm"`
	Ventilated            bool      `json:"ventilated"`
	TemperatureControlled bool      `json:"temperature_controlled"`
}

// AddressDTO represents an address in requests/responses.
type AddressDTO struct {
	Line1      string  `json:"line1" binding:"required"`
	Line2      string  `json:"line2"`
	City       string  `json:"city" binding:"required"`
	State      string  `json:"state" binding:"required"`
	PostalCode string  `json:"postal_code"`
	Country    string  `json:"country" binding:"required"`
	Latitude   float64 `json:"latitude" binding:"required"`
	Longitude  float64 `json:"longitude" binding:"required"`
}

// PetSpecDTO represents a pet specification in requests/responses.
type PetSpecDTO struct {
	PetType      string              `json:"pet_type" binding:"required"`
	Breed        string              `json:"breed"`
	Name         string              `json:"name" binding:"required"`
	WeightKg     float64             `json:"weight_kg" binding:"required"`
	Age          int                 `json:"age_months"`
	Vaccinations []VaccinationDTO    `json:"vaccinations"`
	SpecialNeeds string              `json:"special_needs"`
	PhotoURL     string              `json:"photo_url"`
}

// VaccinationDTO represents a vaccination record.
type VaccinationDTO struct {
	VaccineName string     `json:"vaccine_name" binding:"required"`
	DateGiven   time.Time  `json:"date_given" binding:"required"`
	ExpiresAt   *time.Time `json:"expires_at"`
	VetName     string     `json:"vet_name"`
	Verified    bool       `json:"verified"`
}

// PetShopDTO represents a pet shop in API responses.
type PetShopDTO struct {
	ID           uuid.UUID  `json:"id"`
	OwnerID      *uuid.UUID `json:"owner_id,omitempty"`
	Name         string     `json:"name"`
	Address      string     `json:"address"`
	Latitude     float64    `json:"latitude"`
	Longitude    float64    `json:"longitude"`
	Phone        string     `json:"phone"`
	Email        string     `json:"email"`
	Category     string     `json:"category"`
	Services     []string   `json:"services"`
	Rating       float64    `json:"rating"`
	ImageURL     string     `json:"image_url,omitempty"`
	OpeningHours string     `json:"opening_hours"`
	Description  string     `json:"description"`
	CreatedAt    time.Time  `json:"created_at"`
}

// CoordinateDTO represents a geographic coordinate.
type CoordinateDTO struct {
	Latitude  float64 `json:"latitude" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
}
