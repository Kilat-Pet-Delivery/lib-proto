package dto

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
)

func TestRunnerApplicationRequest_JSONRoundTrip(t *testing.T) {
	t.Run("with pet experience populated", func(t *testing.T) {
		original := RunnerApplicationRequest{
			Name:                    "Ahmad Luqman",
			Phone:                   "+60123456789",
			ICNumber:                "900101-14-5678",
			VehicleType:             VehicleTypeMotorbike,
			PlateNumber:             "WXY1234",
			PetExperience:           []string{PetExperienceDogs, PetExperienceCats},
			ComfortableWithLivePets: true,
			ConsentAcknowledged:     true,
		}

		data, err := json.Marshal(original)
		if err != nil {
			t.Fatalf("marshal failed: %v", err)
		}

		var decoded RunnerApplicationRequest
		if err := json.Unmarshal(data, &decoded); err != nil {
			t.Fatalf("unmarshal failed: %v", err)
		}

		if !reflect.DeepEqual(original, decoded) {
			t.Errorf("round-trip mismatch: got %+v, want %+v", decoded, original)
		}
	})

	t.Run("with empty pet experience slice", func(t *testing.T) {
		original := RunnerApplicationRequest{
			Name:                    "Siti Nurhaliza",
			Phone:                   "+60198765432",
			ICNumber:                "950202-10-1234",
			VehicleType:             VehicleTypeCar,
			PlateNumber:             "PQR5678",
			PetExperience:           []string{},
			ComfortableWithLivePets: false,
			ConsentAcknowledged:     true,
		}

		data, err := json.Marshal(original)
		if err != nil {
			t.Fatalf("marshal failed: %v", err)
		}

		var decoded RunnerApplicationRequest
		if err := json.Unmarshal(data, &decoded); err != nil {
			t.Fatalf("unmarshal failed: %v", err)
		}

		if !reflect.DeepEqual(original, decoded) {
			t.Errorf("round-trip mismatch: got %+v, want %+v", decoded, original)
		}
	})
}

func TestRunnerApplicationRequest_RequiredFields(t *testing.T) {
	valid := RunnerApplicationRequest{
		Name:                    "Ahmad Luqman",
		Phone:                   "+60123456789",
		ICNumber:                "900101-14-5678",
		VehicleType:             VehicleTypeMotorbike,
		PlateNumber:             "WXY1234",
		PetExperience:           []string{},
		ComfortableWithLivePets: false,
		ConsentAcknowledged:     true,
	}

	t.Run("valid struct passes", func(t *testing.T) {
		if err := valid.Validate(); err != nil {
			t.Errorf("expected nil error for valid struct, got: %v", err)
		}
	})

	t.Run("missing name", func(t *testing.T) {
		r := valid
		r.Name = ""
		err := r.Validate()
		if err == nil {
			t.Fatal("expected error for missing name, got nil")
		}
		if !strings.Contains(err.Error(), "name") {
			t.Errorf("expected error to mention 'name', got: %v", err)
		}
	})

	t.Run("missing phone", func(t *testing.T) {
		r := valid
		r.Phone = ""
		err := r.Validate()
		if err == nil {
			t.Fatal("expected error for missing phone, got nil")
		}
		if !strings.Contains(err.Error(), "phone") {
			t.Errorf("expected error to mention 'phone', got: %v", err)
		}
	})

	t.Run("missing IC number", func(t *testing.T) {
		r := valid
		r.ICNumber = ""
		err := r.Validate()
		if err == nil {
			t.Fatal("expected error for missing ICNumber, got nil")
		}
		if !strings.Contains(err.Error(), "icNumber") {
			t.Errorf("expected error to mention 'icNumber', got: %v", err)
		}
	})

	t.Run("missing plate number", func(t *testing.T) {
		r := valid
		r.PlateNumber = ""
		err := r.Validate()
		if err == nil {
			t.Fatal("expected error for missing plateNumber, got nil")
		}
		if !strings.Contains(err.Error(), "plateNumber") {
			t.Errorf("expected error to mention 'plateNumber', got: %v", err)
		}
	})

	t.Run("invalid vehicle type", func(t *testing.T) {
		r := valid
		r.VehicleType = "hoverboard"
		err := r.Validate()
		if err == nil {
			t.Fatal("expected error for invalid vehicleType, got nil")
		}
		if !strings.Contains(err.Error(), "vehicleType") {
			t.Errorf("expected error to mention 'vehicleType', got: %v", err)
		}
	})

	t.Run("consent not acknowledged", func(t *testing.T) {
		r := valid
		r.ConsentAcknowledged = false
		err := r.Validate()
		if err == nil {
			t.Fatal("expected error for consent not acknowledged, got nil")
		}
		if !strings.Contains(err.Error(), "consent") {
			t.Errorf("expected error to mention 'consent', got: %v", err)
		}
	})

	t.Run("invalid_pet_experience", func(t *testing.T) {
		r := valid
		r.PetExperience = []string{"dogs", "unicorns"}
		err := r.Validate()
		if err == nil {
			t.Fatal("expected error for invalid petExperience entry, got nil")
		}
		if !strings.Contains(err.Error(), "petExperience") {
			t.Errorf("expected error to mention 'petExperience', got: %v", err)
		}
	})
}
