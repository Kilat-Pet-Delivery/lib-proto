package events

import (
	"time"

	"github.com/google/uuid"
)

// Kafka topic for chat events.
const TopicChatEvents = "chat.events"

// Chat event types.
const (
	ChatThreadCreated = "chat.thread_created"
	ChatMessageSent   = "chat.message_sent"
	ChatMessageRead   = "chat.message_read"
	ChatTyping        = "chat.typing"
)

// ChatThreadCreatedEvent is published when a booking chat thread is created.
type ChatThreadCreatedEvent struct {
	ThreadID       uuid.UUID `json:"thread_id"`
	BookingID      uuid.UUID `json:"booking_id"`
	CustomerUserID uuid.UUID `json:"customer_user_id"`
	RunnerUserID   uuid.UUID `json:"runner_user_id"`
	OccurredAt     time.Time `json:"occurred_at"`
}

// ChatMessageSentEvent is published when a message is persisted.
type ChatMessageSentEvent struct {
	MessageID                 uuid.UUID `json:"message_id"`
	ThreadID                  uuid.UUID `json:"thread_id"`
	SenderUserID              uuid.UUID `json:"sender_user_id"`
	RecipientUserID           uuid.UUID `json:"recipient_user_id"`
	Body                      string    `json:"body,omitempty"`
	AttachmentURL             string    `json:"attachment_url,omitempty"`
	AttachmentMIME            string    `json:"attachment_mime,omitempty"`
	RecipientOnlineAtSendTime bool      `json:"recipient_online_at_send_time"`
	OccurredAt                time.Time `json:"occurred_at"`
}

// ChatMessageReadEvent is published when a participant marks messages read.
type ChatMessageReadEvent struct {
	ThreadID          uuid.UUID `json:"thread_id"`
	UserID            uuid.UUID `json:"user_id"`
	LastReadMessageID uuid.UUID `json:"last_read_message_id"`
	OccurredAt        time.Time `json:"occurred_at"`
}

// ChatTypingEvent is published for ephemeral typing indicators.
type ChatTypingEvent struct {
	ThreadID   uuid.UUID `json:"thread_id"`
	UserID     uuid.UUID `json:"user_id"`
	ExpiresAt  time.Time `json:"expires_at"`
	OccurredAt time.Time `json:"occurred_at"`
}
