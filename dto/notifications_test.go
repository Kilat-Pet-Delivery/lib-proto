package dto

import (
	"encoding/json"
	"testing"
	"time"
)

func TestNotificationListResponse_JSONRoundTrip(t *testing.T) {
	readAt := time.Date(2026, 5, 16, 10, 0, 0, 0, time.UTC)
	createdAt1 := time.Date(2026, 5, 15, 8, 30, 0, 0, time.UTC)
	createdAt2 := time.Date(2026, 5, 16, 9, 0, 0, 0, time.UTC)

	original := NotificationListResponse{
		Items: []NotificationItem{
			{
				ID:        "notif-001",
				Type:      "booking_confirmed",
				Title:     "Booking Confirmed",
				Body:      "Your booking has been confirmed.",
				CreatedAt: createdAt1,
				ReadAt:    nil,
			},
			{
				ID:        "notif-002",
				Type:      "payment_received",
				Title:     "Payment Received",
				Body:      "You received a payment of MYR 50.00.",
				CreatedAt: createdAt2,
				ReadAt:    &readAt,
			},
		},
		NextCursor: "cursor-abc",
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}

	var decoded NotificationListResponse
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}

	// Validate NextCursor
	if decoded.NextCursor != "cursor-abc" {
		t.Errorf("NextCursor: got %q, want %q", decoded.NextCursor, "cursor-abc")
	}

	// Validate item count
	if len(decoded.Items) != 2 {
		t.Fatalf("Items length: got %d, want 2", len(decoded.Items))
	}

	// Item 0 — unread (ReadAt nil)
	item0 := decoded.Items[0]
	if item0.ID != "notif-001" {
		t.Errorf("Items[0].ID: got %q, want %q", item0.ID, "notif-001")
	}
	if item0.Type != "booking_confirmed" {
		t.Errorf("Items[0].Type: got %q, want %q", item0.Type, "booking_confirmed")
	}
	if item0.Title != "Booking Confirmed" {
		t.Errorf("Items[0].Title: got %q, want %q", item0.Title, "Booking Confirmed")
	}
	if item0.Body != "Your booking has been confirmed." {
		t.Errorf("Items[0].Body: got %q, want %q", item0.Body, "Your booking has been confirmed.")
	}
	if !item0.CreatedAt.Equal(createdAt1) {
		t.Errorf("Items[0].CreatedAt: got %v, want %v", item0.CreatedAt, createdAt1)
	}
	if item0.ReadAt != nil {
		t.Errorf("Items[0].ReadAt: got %v, want nil (unread)", item0.ReadAt)
	}

	// Item 1 — read (ReadAt populated)
	item1 := decoded.Items[1]
	if item1.ID != "notif-002" {
		t.Errorf("Items[1].ID: got %q, want %q", item1.ID, "notif-002")
	}
	if item1.Type != "payment_received" {
		t.Errorf("Items[1].Type: got %q, want %q", item1.Type, "payment_received")
	}
	if item1.Title != "Payment Received" {
		t.Errorf("Items[1].Title: got %q, want %q", item1.Title, "Payment Received")
	}
	if item1.Body != "You received a payment of MYR 50.00." {
		t.Errorf("Items[1].Body: got %q, want %q", item1.Body, "You received a payment of MYR 50.00.")
	}
	if !item1.CreatedAt.Equal(createdAt2) {
		t.Errorf("Items[1].CreatedAt: got %v, want %v", item1.CreatedAt, createdAt2)
	}
	if item1.ReadAt == nil {
		t.Fatal("Items[1].ReadAt: got nil, want non-nil (read)")
	}
	if !item1.ReadAt.Equal(readAt) {
		t.Errorf("Items[1].ReadAt: got %v, want %v", item1.ReadAt, readAt)
	}

	// Confirm unread item's JSON does not contain the readAt key
	var raw map[string]interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		t.Fatalf("re-unmarshal to raw map failed: %v", err)
	}
	items, _ := raw["items"].([]interface{})
	if len(items) < 1 {
		t.Fatal("raw items array empty")
	}
	item0Raw, _ := items[0].(map[string]interface{})
	if _, hasReadAt := item0Raw["readAt"]; hasReadAt {
		t.Error("Items[0] JSON should not contain 'readAt' key for unread item (omitempty)")
	}
}

func TestNotificationListResponse_EmptyItems(t *testing.T) {
	original := NotificationListResponse{
		Items:      []NotificationItem{},
		NextCursor: "",
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}

	var decoded NotificationListResponse
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}

	if len(decoded.Items) != 0 {
		t.Errorf("Items length: got %d, want 0", len(decoded.Items))
	}
	if decoded.NextCursor != "" {
		t.Errorf("NextCursor: got %q, want empty", decoded.NextCursor)
	}
}
