package contracts

import (
	"fmt"
	"time"

	manifest "github.com/estafette/estafette-ci-manifest"
)

// Release represents a release of a pipeline
type Release struct {
	Name            string                    `json:"name"`
	Action          string                    `json:"action,omitempty"`
	ID              string                    `json:"id,omitempty"`
	RepoSource      string                    `json:"repoSource,omitempty"`
	RepoOwner       string                    `json:"repoOwner,omitempty"`
	RepoName        string                    `json:"repoName,omitempty"`
	ReleaseVersion  string                    `json:"releaseVersion,omitempty"`
	ReleaseStatus   string                    `json:"releaseStatus,omitempty"`
	Events          []manifest.EstafetteEvent `json:"triggerEvents,omitempty"`
	InsertedAt      *time.Time                `json:"insertedAt,omitempty"`
	StartedAt       *time.Time                `json:"startedAt,omitempty"`
	UpdatedAt       *time.Time                `json:"updatedAt,omitempty"`
	Duration        *time.Duration            `json:"duration,omitempty"`
	PendingDuration *time.Duration            `json:"pendingDuration,omitempty"`
}

// GetFullRepoPath returns the full path of the release repository with source, owner and name
func (release *Release) GetFullRepoPath() string {
	return fmt.Sprintf("%v/%v/%v", release.RepoSource, release.RepoOwner, release.RepoName)
}
