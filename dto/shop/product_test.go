package shop

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/google/uuid"
)

func Test_UpdateInventoryRequest_RejectsNegativeQty(t *testing.T) {
	tests := []struct {
		name    string
		req     UpdateInventoryRequest
		wantErr string
	}{
		{
			name:    "negative quantity rejected",
			req:     UpdateInventoryRequest{Quantity: -1, Reason: InventoryMovementReasonAdjustment, ExpectedVersion: 0},
			wantErr: "quantity",
		},
		{
			name:    "negative expected_version rejected",
			req:     UpdateInventoryRequest{Quantity: 5, Reason: InventoryMovementReasonRestock, ExpectedVersion: -1},
			wantErr: "expected_version",
		},
		{
			name:    "unknown reason rejected",
			req:     UpdateInventoryRequest{Quantity: 5, Reason: "shrinkage", ExpectedVersion: 0},
			wantErr: "reason",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.req.Validate()
			if err == nil {
				t.Fatalf("expected error mentioning %q, got nil", tt.wantErr)
			}
			if !strings.Contains(err.Error(), tt.wantErr) {
				t.Errorf("error should mention %q, got: %v", tt.wantErr, err)
			}
		})
	}
}

func Test_UpdateInventoryRequest_AcceptsZeroQuantity(t *testing.T) {
	// Quantity is the absolute target stock level; zero is valid (stock-out).
	req := UpdateInventoryRequest{Quantity: 0, Reason: InventoryMovementReasonSale, ExpectedVersion: 3}
	if err := req.Validate(); err != nil {
		t.Errorf("expected nil for zero target quantity, got: %v", err)
	}
}

func Test_UpdateInventoryRequest_JSONRoundTrip(t *testing.T) {
	orig := UpdateInventoryRequest{
		Quantity:        12,
		Reason:          InventoryMovementReasonRestock,
		ExpectedVersion: 7,
		Note:            "Monday delivery",
	}
	data, err := json.Marshal(orig)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}
	var decoded UpdateInventoryRequest
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if decoded != orig {
		t.Errorf("round-trip mismatch: got %+v, want %+v", decoded, orig)
	}
}

func Test_CreateProductRequest_Validate(t *testing.T) {
	tests := []struct {
		name    string
		req     CreateProductRequest
		wantErr string
	}{
		{
			name: "happy path",
			req: CreateProductRequest{
				SKU: "ABC-001", Name: "Dog Food", PriceCents: 2500, Currency: "MYR",
			},
		},
		{
			name:    "missing sku",
			req:     CreateProductRequest{Name: "x", PriceCents: 100, Currency: "MYR"},
			wantErr: "sku",
		},
		{
			name:    "missing name",
			req:     CreateProductRequest{SKU: "x", PriceCents: 100, Currency: "MYR"},
			wantErr: "name",
		},
		{
			name:    "zero price rejected",
			req:     CreateProductRequest{SKU: "x", Name: "x", PriceCents: 0, Currency: "MYR"},
			wantErr: "price_cents",
		},
		{
			name:    "missing currency",
			req:     CreateProductRequest{SKU: "x", Name: "x", PriceCents: 100},
			wantErr: "currency",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.req.Validate()
			if tt.wantErr != "" {
				if err == nil {
					t.Fatalf("expected error mentioning %q, got nil", tt.wantErr)
				}
				if !strings.Contains(err.Error(), tt.wantErr) {
					t.Errorf("error should mention %q, got: %v", tt.wantErr, err)
				}
				return
			}
			if err != nil {
				t.Errorf("expected nil for valid request, got: %v", err)
			}
		})
	}
}

func Test_UpdateProductRequest_PartialUpdate(t *testing.T) {
	// Unset fields validate as no-op; only supplied fields are checked.
	emptyReq := UpdateProductRequest{}
	if err := emptyReq.Validate(); err != nil {
		t.Errorf("empty update should validate, got: %v", err)
	}

	zeroPrice := int64(0)
	badReq := UpdateProductRequest{PriceCents: &zeroPrice}
	if err := badReq.Validate(); err == nil {
		t.Error("zero price_cents should be rejected")
	}

	goodPrice := int64(500)
	goodReq := UpdateProductRequest{PriceCents: &goodPrice}
	if err := goodReq.Validate(); err != nil {
		t.Errorf("positive price_cents should validate, got: %v", err)
	}
}

func Test_ProductDTO_JSONRoundTrip(t *testing.T) {
	orig := ProductDTO{
		ID:            uuid.New(),
		ShopID:        uuid.New(),
		SKU:           "ABC-001",
		Name:          "Dog Food",
		Description:   "Premium kibble",
		PriceCents:    2500,
		Currency:      "MYR",
		Category:      "food",
		Active:        true,
		StockQuantity: 20,
		Version:       3,
	}
	data, err := json.Marshal(orig)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}
	var decoded ProductDTO
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if decoded.ID != orig.ID || decoded.SKU != orig.SKU || decoded.Version != orig.Version {
		t.Errorf("round-trip mismatch: got %+v, want %+v", decoded, orig)
	}
}

func Test_InventoryMovementReason_IsValid(t *testing.T) {
	valid := []string{
		InventoryMovementReasonSale,
		InventoryMovementReasonRestock,
		InventoryMovementReasonAdjustment,
		InventoryMovementReasonReturn,
		InventoryMovementReasonDamage,
	}
	for _, r := range valid {
		if !IsValidInventoryMovementReason(r) {
			t.Errorf("expected %q to be valid", r)
		}
	}
	if IsValidInventoryMovementReason("shrinkage") {
		t.Error("expected 'shrinkage' to be invalid")
	}
}
