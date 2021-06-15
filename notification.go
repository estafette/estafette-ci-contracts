package contracts

import "time"

type NotificationType string

const (
	NotificationTypeUnknown       NotificationType = ""
	NotificationTypeVulnerability NotificationType = "vulnerability"
	NotificationTypeWarning       NotificationType = "warning"
)

type NotificationLevel string

const (
	NotificationLevelUnknown  NotificationLevel = ""
	NotificationLevelCritical NotificationLevel = "critical"
	NotificationLevelHigh     NotificationLevel = "high"
	NotificationLevelMedium   NotificationLevel = "medium"
	NotificationLevelLow      NotificationLevel = "low"
)

type NotificationLinkType string

const (
	NotificationLinkTypeUnknown   NotificationLinkType = ""
	NotificationLinkTypePipeline  NotificationLinkType = "pipeline"
	NotificationLinkTypeContainer NotificationLinkType = "container"
)

type Notification struct {
	Type    NotificationType  `json:"type,omitempty"`
	Level   NotificationLevel `json:"level,omitempty"`
	Message string            `json:"message,omitempty"`
}

type NotificationRecord struct {
	ID            string               `json:"id,omitempty"`
	LinkType      NotificationLinkType `json:"linkType,omitempty"`
	LinkEntity    string               `json:"linkEntity,omitempty"`
	Source        string               `json:"source,omitempty"`
	Notifications []Notification       `json:"notifications,omitempty"`
	InsertedAt    *time.Time           `json:"insertedAt,omitempty"`
	Groups        []*Group             `json:"groups,omitempty"`
	Organizations []*Organization      `json:"organizations,omitempty"`
}
