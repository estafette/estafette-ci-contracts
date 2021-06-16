package contracts

import (
	"fmt"
	"time"
)

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
	ID       string               `json:"id,omitempty"`
	LinkType NotificationLinkType `json:"linkType,omitempty"`
	LinkID   string               `json:"linkID,omitempty"`

	// fields mapped to link_detail column
	PipelineDetail  *PipelineLinkDetail  `json:"pipelineDetail,omitempty"`
	ContainerDetail *ContainerLinkDetail `json:"containerDetail,omitempty"`

	Source        string          `json:"source,omitempty"`
	Notifications []Notification  `json:"notifications,omitempty"`
	InsertedAt    *time.Time      `json:"insertedAt,omitempty"`
	Groups        []*Group        `json:"groups,omitempty"`
	Organizations []*Organization `json:"organizations,omitempty"`
}

func (nr *NotificationRecord) GetLinkDetail() interface{} {
	switch nr.LinkType {
	case NotificationLinkTypePipeline:
		return nr.PipelineDetail
	case NotificationLinkTypeContainer:
		return nr.ContainerDetail
	}

	return nil
}

func (nr *NotificationRecord) SetLinkDetail(linkDetail interface{}) (err error) {
	if linkDetail == nil {
		return
	}

	var ok bool

	switch nr.LinkType {
	case NotificationLinkTypePipeline:
		if nr.PipelineDetail, ok = linkDetail.(*PipelineLinkDetail); !ok {
			return fmt.Errorf("LinkDetail for NotificationRecord %v of type %v is not of type PipelineLinkDetail", nr.LinkID, nr.LinkType)
		}
	case NotificationLinkTypeContainer:
		if nr.ContainerDetail, ok = linkDetail.(*ContainerLinkDetail); !ok {
			return fmt.Errorf("LinkDetail for NotificationRecord %v of type %v is not of type ContainerLinkDetail", nr.LinkID, nr.LinkType)
		}
	}

	return nil
}

type PipelineLinkDetail struct {
	Branch   string `json:"branch,omitempty"`
	Revision string `json:"revision"`
	Version  string `json:"version,omitempty"`
	Status   Status `json:"status,omitempty"`
}

type ContainerLinkDetail struct {
	Tag         string `json:"tag,omitempty"`
	PublicImage bool   `json:"publicImage,omitempty"`
}
