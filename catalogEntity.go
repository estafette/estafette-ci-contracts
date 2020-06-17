package contracts

import "time"

// CatalogEntity represents any entity stored in the catalog tree
type CatalogEntity struct {
	ID             string                 `json:"id,omitempty"`
	ParentKey      string                 `json:"parent_key,omitempty"`
	ParentValue    string                 `json:"parent_value,omitempty"`
	Key            string                 `json:"entity_key,omitempty"`
	Value          string                 `json:"entity_value,omitempty"`
	LinkedPipeline string                 `json:"linked_pipeline,omitempty"`
	Labels         []Label                `json:"labels,omitempty"`
	Metadata       map[string]interface{} `json:"entity_metadata,omitempty"`
	InsertedAt     *time.Time             `json:"insertedAt,omitempty"`
	UpdatedAt      *time.Time             `json:"updatedAt,omitempty"`
}
