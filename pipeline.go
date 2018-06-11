package contracts

import "time"

// Pipeline represents a pipeline with the latest build info, including version number, repo, branch, revision, labels and manifest
type Pipeline struct {
	ID           string    `jsonapi:"primary,pipelines" json:"id"`
	RepoSource   string    `jsonapi:"attr,repo-source" json:"repoSource"`
	RepoOwner    string    `jsonapi:"attr,repo-owner" json:"repoOwner"`
	RepoName     string    `jsonapi:"attr,repo-name" json:"repoName"`
	RepoBranch   string    `jsonapi:"attr,repo-branch" json:"repoBranch"`
	RepoRevision string    `jsonapi:"attr,repo-revision" json:"repoRevision"`
	BuildVersion string    `jsonapi:"attr,build-version" json:"buildVersion"`
	BuildStatus  string    `jsonapi:"attr,build-status" json:"buildStatus"`
	Labels       string    `jsonapi:"attr,labels" json:"labels"`
	Manifest     string    `jsonapi:"attr,manifest" json:"manifest"`
	InsertedAt   time.Time `jsonapi:"attr,inserted-at" json:"insertedAt"`
	UpdatedAt    time.Time `jsonapi:"attr,updated-at" json:"updatedAt"`
}
