package contracts

import "errors"

type EstafetteCiBuilderEvent struct {
	JobType JobType `json:"jobType,omitempty"`
	JobName string  `json:"job_name"`
	PodName string  `json:"pod_name,omitempty"`

	Build   *Build     `json:"build,omitempty"`
	Release *Release   `json:"release,omitempty"`
	Bot     *Bot       `json:"bot,omitempty"`
	Git     *GitConfig `json:"git,omitempty"`

	// deprecated
	RepoSource   string `json:"repo_source,omitempty"`
	RepoOwner    string `json:"repo_owner,omitempty"`
	RepoName     string `json:"repo_name,omitempty"`
	RepoBranch   string `json:"repo_branch,omitempty"`
	RepoRevision string `json:"repo_revision,omitempty"`
	ReleaseID    string `json:"release_id,omitempty"`
	BuildID      string `json:"build_id,omitempty"`
	BuildStatus  Status `json:"build_status,omitempty"`
}

func (bc *EstafetteCiBuilderEvent) Validate() (err error) {

	if bc.Git == nil {
		return errors.New("git needs to be set")
	}

	switch bc.JobType {
	case JobTypeBuild:
		if bc.Build == nil {
			return errors.New("build needs to be set for jobType build")
		}
	case JobTypeRelease:
		if bc.Release == nil {
			return errors.New("release needs to be set for jobType release")
		}
	case JobTypeBot:
		if bc.Bot == nil {
			return errors.New("bot needs to be set for jobType bot")
		}
	}

	return nil
}
