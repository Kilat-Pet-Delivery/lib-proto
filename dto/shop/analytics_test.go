package shop

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/google/uuid"
)

func Test_AnalyticsPeriod_IsValid(t *testing.T) {
	valid := []string{AnalyticsPeriodDay, AnalyticsPeriodWeek, AnalyticsPeriodMonth}
	for _, p := range valid {
		if !IsValidAnalyticsPeriod(p) {
			t.Errorf("expected %q to be valid", p)
		}
	}
	if IsValidAnalyticsPeriod("year") {
		t.Error("'year' should not yet be a recognized analytics period")
	}
}

func Test_SalesAggregateDTO_JSONRoundTrip(t *testing.T) {
	shopID := uuid.New()
	periodStart := time.Date(2026, 5, 23, 0, 0, 0, 0, time.UTC)
	orig := SalesAggregateDTO{
		ShopID:          shopID,
		Period:          AnalyticsPeriodDay,
		PeriodStart:     periodStart,
		PeriodEnd:       periodStart.Add(24 * time.Hour),
		OrderCount:      12,
		GrossSalesCents: 200000,
		NetSalesCents:   180000,
		Currency:        "MYR",
		TopProducts: []TopProductDTO{
			{ProductID: uuid.New(), SKU: "A", Name: "Food", UnitsSold: 4, RevenueCents: 80000},
		},
		RefreshedAt: time.Now().UTC().Truncate(time.Second),
	}

	data, err := json.Marshal(orig)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}
	var decoded SalesAggregateDTO
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if decoded.ShopID != orig.ShopID || decoded.Period != orig.Period || decoded.OrderCount != orig.OrderCount {
		t.Errorf("round-trip mismatch: got %+v, want %+v", decoded, orig)
	}
	if len(decoded.TopProducts) != 1 || decoded.TopProducts[0].SKU != "A" {
		t.Errorf("top_products lost in round trip: %+v", decoded.TopProducts)
	}
}
