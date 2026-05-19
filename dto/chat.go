package dto

import (
	"time"

	"github.com/google/uuid"
)

// ChatThread represents a customer-runner chat thread bound to a booking.
type ChatThread struct {
	ID             uuid.UUID         `json:"id"`
	BookingID      uuid.UUID         `json:"booking_id"`
	CustomerUserID uuid.UUID         `json:"customer_user_id"`
	RunnerUserID   uuid.UUID         `json:"runner_user_id"`
	Participants   []ChatParticipant `json:"participants,omitempty"`
	LastMessage    *ChatMessage      `json:"last_message,omitempty"`
	UnreadCount    int               `json:"unread_count"`
	CreatedAt      time.Time         `json:"created_at"`
	ArchivedAt     *time.Time        `json:"archived_at,omitempty"`
}

// ChatParticipant describes one participant in a chat thread.
type ChatParticipant struct {
	UserID      uuid.UUID  `json:"user_id"`
	DisplayName string     `json:"display_name"`
	Role        string     `json:"role"`
	AvatarURL   string     `json:"avatar_url,omitempty"`
	Presence    string     `json:"presence,omitempty"`
	LastSeenAt  *time.Time `json:"last_seen_at,omitempty"`
}

// ChatMessage is a persisted message.
type ChatMessage struct {
	ID             uuid.UUID `json:"id"`
	ThreadID       uuid.UUID `json:"thread_id"`
	SenderUserID   uuid.UUID `json:"sender_user_id"`
	Body           string    `json:"body,omitempty"`
	AttachmentURL  string    `json:"attachment_url,omitempty"`
	AttachmentMIME string    `json:"attachment_mime,omitempty"`
	DeliveryState  string    `json:"delivery_state"`
	CreatedAt      time.Time `json:"created_at"`
}

// ChatReadReceipt captures the read state for a thread participant.
type ChatReadReceipt struct {
	ThreadID          uuid.UUID `json:"thread_id"`
	UserID            uuid.UUID `json:"user_id"`
	LastReadMessageID uuid.UUID `json:"last_read_message_id"`
	LastReadAt        time.Time `json:"last_read_at"`
}

// ListChatThreadsRequest contains thread list filters.
type ListChatThreadsRequest struct {
	Archived bool `json:"archived"`
	Limit    int  `json:"limit"`
	Offset   int  `json:"offset"`
}

// ListChatThreadsResponse returns paginated chat threads.
type ListChatThreadsResponse struct {
	Threads []ChatThread `json:"threads"`
	Total   int          `json:"total"`
}

// FetchChatMessagesRequest contains cursor pagination options.
type FetchChatMessagesRequest struct {
	Cursor string `json:"cursor,omitempty"`
	Limit  int    `json:"limit"`
}

// FetchChatMessagesResponse returns cursor-paginated messages.
type FetchChatMessagesResponse struct {
	Messages   []ChatMessage `json:"messages"`
	NextCursor string        `json:"next_cursor,omitempty"`
}

// SendChatMessageRequest creates a text and/or attachment message.
type SendChatMessageRequest struct {
	Body           string `json:"body,omitempty"`
	AttachmentURL  string `json:"attachment_url,omitempty"`
	AttachmentMIME string `json:"attachment_mime,omitempty"`
}

// MarkChatReadRequest marks a thread read through a message id.
type MarkChatReadRequest struct {
	MessageID uuid.UUID `json:"message_id" binding:"required"`
}

// QuickReply is a reusable runner chat response.
type QuickReply struct {
	ID        uuid.UUID `json:"id"`
	Label     string    `json:"label"`
	Body      string    `json:"body"`
	SortOrder int       `json:"sort_order"`
}
