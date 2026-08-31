package shop

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
)

func Test_ShopStatus_JSONRoundTrip(t *testing.T) {
	statuses := []string{
		ShopStatusOpen,
		ShopStatusBusy,
		ShopStatusClosed,
		ShopStatusPaused,
	}

	for _, status := range statuses {
		t.Run(status, func(t *testing.T) {
			orig := UpdateShopStatusRequest{Status: status}
			data, err := json.Marshal(orig)
			if err != nil {
				t.Fatalf("marshal failed: %v", err)
			}
			var decoded UpdateShopStatusRequest
			if err := json.Unmarshal(data, &decoded); err != nil {
				t.Fatalf("unmarshal failed: %v", err)
			}
			if decoded.Status != orig.Status {
				t.Errorf("round-trip mismatch: got %q, want %q", decoded.Status, orig.Status)
			}
			if !IsValidShopStatus(decoded.Status) {
				t.Errorf("IsValidShopStatus rejected its own round-tripped value: %q", decoded.Status)
			}
		})
	}
}

func Test_ShopCategory_JSONRoundTrip(t *testing.T) {
	categories := []string{
		ShopCategoryVet,
		ShopCategoryGrooming,
		ShopCategorySupplies,
	}

	for _, cat := range categories {
		t.Run(cat, func(t *testing.T) {
			orig := CreateShopRequest{
				Name:      "Shop",
				Slug:      "shop",
				Address:   "1 Test St",
				Latitude:  1.0,
				Longitude: 1.0,
				Phone:     "+60123456789",
				Category:  cat,
			}
			data, err := json.Marshal(orig)
			if err != nil {
				t.Fatalf("marshal failed: %v", err)
			}
			var decoded CreateShopRequest
			if err := json.Unmarshal(data, &decoded); err != nil {
				t.Fatalf("unmarshal failed: %v", err)
			}
			if decoded.Category != orig.Category {
				t.Errorf("round-trip mismatch: got %q, want %q", decoded.Category, orig.Category)
			}
			if err := decoded.Validate(); err != nil {
				t.Errorf("Validate rejected a valid round-tripped request: %v", err)
			}
		})
	}
}

func Test_CreateShopRequest_Validate(t *testing.T) {
	tests := []struct {
		name     string
		req      CreateShopRequest
		wantErr  bool
		errField string
	}{
		{
			name: "happy path",
			req: CreateShopRequest{
				Name: "Pet Place", Slug: "pet-place", Address: "1 Test St",
				Latitude: 3.1, Longitude: 101.6, Phone: "+60123456789",
				Category: ShopCategoryVet,
			},
		},
		{
			name:     "missing name",
			req:      CreateShopRequest{Slug: "x", Address: "x", Phone: "x", Category: ShopCategoryVet},
			wantErr:  true,
			errField: "name",
		},
		{
			name:     "missing slug",
			req:      CreateShopRequest{Name: "x", Address: "x", Phone: "x", Category: ShopCategoryVet},
			wantErr:  true,
			errField: "slug",
		},
		{
			name:     "missing address",
			req:      CreateShopRequest{Name: "x", Slug: "x", Phone: "x", Category: ShopCategoryVet},
			wantErr:  true,
			errField: "address",
		},
		{
			name:     "missing phone",
			req:      CreateShopRequest{Name: "x", Slug: "x", Address: "x", Category: ShopCategoryVet},
			wantErr:  true,
			errField: "phone",
		},
		{
			name:     "unknown category",
			req:      CreateShopRequest{Name: "x", Slug: "x", Address: "x", Phone: "x", Category: "boarding"},
			wantErr:  true,
			errField: "category",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.req.Validate()
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error for %s, got nil", tt.errField)
				}
				if !strings.Contains(err.Error(), tt.errField) {
					t.Errorf("error should mention %q, got: %v", tt.errField, err)
				}
				return
			}
			if err != nil {
				t.Errorf("expected nil for valid request, got: %v", err)
			}
		})
	}
}

func Test_UpdateShopStatusRequest_RejectsUnknown(t *testing.T) {
	req := UpdateShopStatusRequest{Status: "vacationing"}
	if err := req.Validate(); err == nil {
		t.Fatal("expected error for unknown status, got nil")
	}
}

func Test_ShopDTO_JSONRoundTrip(t *testing.T) {
	autoClose := time.Date(2026, 6, 1, 16, 0, 0, 0, time.UTC)
	orig := ShopDTO{
		ID:          uuid.New(),
		OwnerUserID: uuid.New(),
		Name:        "Pet Place",
		Slug:        "pet-place",
		Address:     "1 Test St",
		Latitude:    3.1,
		Longitude:   101.6,
		Phone:       "+60123456789",
		Category:    ShopCategoryVet,
		Status:      ShopStatusOpen,
		AutoCloseAt: &autoClose,
		CreatedAt:   time.Now().UTC().Truncate(time.Second),
		UpdatedAt:   time.Now().UTC().Truncate(time.Second),
	}

	data, err := json.Marshal(orig)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}
	var decoded ShopDTO
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if decoded.ID != orig.ID || decoded.Status != orig.Status || decoded.Category != orig.Category {
		t.Errorf("round-trip mismatch: got %+v, want %+v", decoded, orig)
	}
	if decoded.AutoCloseAt == nil || !decoded.AutoCloseAt.Equal(autoClose) {
		t.Errorf("auto_close_at round-trip lost the timestamp: got %v, want %v", decoded.AutoCloseAt, autoClose)
	}
}

func Test_ShopDTO_AutoCloseAtOmitsWhenNil(t *testing.T) {
	orig := ShopDTO{
		ID:          uuid.New(),
		OwnerUserID: uuid.New(),
		Name:        "Pet Place",
		Slug:        "pet-place",
		Status:      ShopStatusOpen,
		Category:    ShopCategoryVet,
	}

	data, err := json.Marshal(orig)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}
	if strings.Contains(string(data), "auto_close_at") {
		t.Errorf("auto_close_at should be omitted when nil, got: %s", string(data))
	}
}
