package dto

import "fmt"

// Vehicle type constants for RunnerApplicationRequest.
const (
	VehicleTypeMotorbike = "motorbike"
	VehicleTypeCar       = "car"
	VehicleTypeBicycle   = "bicycle"
)

// Pet experience constants for RunnerApplicationRequest.
const (
	PetExperienceDogs      = "dogs"
	PetExperienceCats      = "cats"
	PetExperienceBirds     = "birds"
	PetExperienceFish      = "fish"
	PetExperienceSmallPets = "small_pets"
	PetExperienceVetTrips  = "vet_trips"
)

// RunnerApplicationRequest is the request body for the runner application form.
// It captures personal details, vehicle information, pet experience, and
// consent acknowledgment collected by the iOS apply form.
type RunnerApplicationRequest struct {
	Name                    string   `json:"name"`
	Phone                   string   `json:"phone"`
	ICNumber                string   `json:"icNumber"`
	VehicleType             string   `json:"vehicleType"`
	PlateNumber             string   `json:"plateNumber"`
	PetExperience           []string `json:"petExperience"`
	ComfortableWithLivePets bool     `json:"comfortableWithLivePets"`
	ConsentAcknowledged     bool     `json:"consentAcknowledged"`
}

// Validate returns nil if the request is valid, or a descriptive error if any
// required field is missing or invalid. ConsentAcknowledged must be true.
func (r RunnerApplicationRequest) Validate() error {
	if r.Name == "" {
		return fmt.Errorf("name is required")
	}
	if r.Phone == "" {
		return fmt.Errorf("phone is required")
	}
	if r.ICNumber == "" {
		return fmt.Errorf("icNumber is required")
	}
	switch r.VehicleType {
	case VehicleTypeMotorbike, VehicleTypeCar, VehicleTypeBicycle:
		// valid
	default:
		return fmt.Errorf("vehicleType must be one of: motorbike, car, bicycle")
	}
	if r.PlateNumber == "" {
		return fmt.Errorf("plateNumber is required")
	}
	for _, exp := range r.PetExperience {
		switch exp {
		case PetExperienceDogs, PetExperienceCats, PetExperienceBirds,
			PetExperienceFish, PetExperienceSmallPets, PetExperienceVetTrips:
			// valid
		default:
			return fmt.Errorf("petExperience entries must be one of: dogs, cats, birds, fish, small_pets, vet_trips")
		}
	}
	if !r.ConsentAcknowledged {
		return fmt.Errorf("consent must be acknowledged")
	}
	return nil
}
