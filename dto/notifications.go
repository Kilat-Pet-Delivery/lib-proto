package dto

import "time"

// NotificationItem represents a single notification in a user's inbox.
// ReadAt is nil for unread notifications; a non-nil pointer indicates when
// the notification was read.
type NotificationItem struct {
	ID        string     `json:"id"`
	Type      string     `json:"type"`
	Title     string     `json:"title"`
	Body      string     `json:"body"`
	CreatedAt time.Time  `json:"createdAt"`
	ReadAt    *time.Time `json:"readAt,omitempty"`
}

// NotificationListResponse is the paginated response for a user's notification
// inbox. NextCursor is empty when there are no further pages.
type NotificationListResponse struct {
	Items      []NotificationItem `json:"items"`
	NextCursor string             `json:"nextCursor"`
}
