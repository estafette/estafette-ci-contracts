package contracts

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

type Notification struct {
	Type    NotificationType  `json:"type"`
	Level   NotificationLevel `json:"level"`
	Message string            `json:"message"`
}
