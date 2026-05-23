package events

import (
	"time"

	"github.com/google/uuid"
)

// Kafka topic for shop-aggregate events. Partition key is shop_id so per-shop
// ordering is preserved across consumers.
const TopicShopEvents = "shop.events"

// Shop event types.
const (
	ShopCreated             = "shop.created"
	ShopStatusChanged       = "shop.status_changed"
	ShopStaffInvited        = "shop.staff_invited"
	ShopStaffAccepted       = "shop.staff_accepted"
	ShopStaffRemoved        = "shop.staff_removed"
	SalesAggregateRefreshed = "shop.sales_aggregate_refreshed"
)

// ShopCreatedEvent is published when an owner self-registers a shop.
type ShopCreatedEvent struct {
	ShopID      uuid.UUID `json:"shop_id"`
	OwnerUserID uuid.UUID `json:"owner_user_id"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"`
	Category    string    `json:"category"`
	OccurredAt  time.Time `json:"occurred_at"`
}

// ShopStatusChangedEvent is published whenever the merchant flips the shop's
// operational status (open/busy/closed/paused).
type ShopStatusChangedEvent struct {
	ShopID      uuid.UUID  `json:"shop_id"`
	OwnerUserID uuid.UUID  `json:"owner_user_id"`
	OldStatus   string     `json:"old_status"`
	NewStatus   string     `json:"new_status"`
	AutoCloseAt *time.Time `json:"auto_close_at,omitempty"`
	OccurredAt  time.Time  `json:"occurred_at"`
}

// ShopStaffInvitedEvent is published when an owner or manager dispatches a
// staff invite. Consumed by service-notification to send the invite email/SMS.
type ShopStaffInvitedEvent struct {
	InviteID  uuid.UUID `json:"invite_id"`
	ShopID    uuid.UUID `json:"shop_id"`
	InviterID uuid.UUID `json:"inviter_user_id"`
	Email     string    `json:"email,omitempty"`
	Phone     string    `json:"phone,omitempty"`
	Role      string    `json:"role"`
	Token     string    `json:"token"`
	OccurredAt time.Time `json:"occurred_at"`
}

// ShopStaffAcceptedEvent is published when an invited staff member accepts.
// Consumed by service-identity to attach the shop-side role to the user.
type ShopStaffAcceptedEvent struct {
	InviteID    uuid.UUID `json:"invite_id"`
	ShopID      uuid.UUID `json:"shop_id"`
	StaffUserID uuid.UUID `json:"staff_user_id"`
	Role        string    `json:"role"`
	OccurredAt  time.Time `json:"occurred_at"`
}

// ShopStaffRemovedEvent is published when an owner removes a staff member or
// the staff member resigns.
type ShopStaffRemovedEvent struct {
	ShopID      uuid.UUID `json:"shop_id"`
	StaffUserID uuid.UUID `json:"staff_user_id"`
	RemovedBy   uuid.UUID `json:"removed_by_user_id"`
	OccurredAt  time.Time `json:"occurred_at"`
}

// SalesAggregateRefreshedEvent is published by service-shop after rolling up
// a (shop, period_start) sales aggregate. Consumed by service-notification
// for the daily/weekly summary push notifications.
type SalesAggregateRefreshedEvent struct {
	ShopID          uuid.UUID `json:"shop_id"`
	Period          string    `json:"period"`
	PeriodStart     time.Time `json:"period_start"`
	PeriodEnd       time.Time `json:"period_end"`
	OrderCount      int64     `json:"order_count"`
	GrossSalesCents int64     `json:"gross_sales_cents"`
	NetSalesCents   int64     `json:"net_sales_cents"`
	Currency        string    `json:"currency"`
	OccurredAt      time.Time `json:"occurred_at"`
}
