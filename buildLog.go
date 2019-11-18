package contracts

import "time"

// BuildLog represents a build log for a specific revision
type BuildLog struct {
	ID           string          `json:"id,omitempty"`
	RepoSource   string          `json:"repoSource"`
	RepoOwner    string          `json:"repoOwner"`
	RepoName     string          `json:"repoName"`
	RepoBranch   string          `json:"repoBranch"`
	RepoRevision string          `json:"repoRevision"`
	BuildID      string          `json:"buildID"`
	Steps        []*BuildLogStep `json:"steps"`
	InsertedAt   time.Time       `json:"insertedAt"`
}

// BuildLogStep represents the logs for a single step of a pipeline
type BuildLogStep struct {
	Step         string                   `json:"step"`
	Depth        int                      `json:"depth,omitempty"`
	Image        *BuildLogStepDockerImage `json:"image"`
	RunIndex     int                      `json:"runIndex,omitempty"`
	Duration     time.Duration            `json:"duration"`
	LogLines     []BuildLogLine           `json:"logLines"`
	ExitCode     int64                    `json:"exitCode"`
	Status       string                   `json:"status"`
	AutoInjected bool                     `json:"autoInjected,omitempty"`
	NestedSteps  []*BuildLogStep          `json:"nestedSteps,omitempty"`
	Services     []*BuildLogStep          `json:"services,omitempty"`
}

// BuildLogStepDockerImage represents info about the docker image used for a step
type BuildLogStepDockerImage struct {
	Name         string        `json:"name"`
	Tag          string        `json:"tag"`
	IsPulled     bool          `json:"isPulled"`
	ImageSize    int64         `json:"imageSize"`
	PullDuration time.Duration `json:"pullDuration"`
	Error        string        `json:"error,omitempty"`
	IsTrusted    bool          `json:"isTrusted,omitempty"`
}

// BuildLogLine has low level log information
type BuildLogLine struct {
	LineNumber int       `json:"line,omitempty"`
	Timestamp  time.Time `json:"timestamp"`
	StreamType string    `json:"streamType"`
	Text       string    `json:"text"`
}

// TailLogLine returns a log line for streaming logs to gui during a build
type TailLogLine struct {
	Step         string                   `json:"step"`
	ParentStage  string                   `json:"parentStage,omitempty"`
	Type         string                   `json:"type"`
	Depth        int                      `json:"depth,omitempty"`
	RunIndex     int                      `json:"runIndex,omitempty"`
	LogLine      *BuildLogLine            `json:"logLine,omitempty"`
	Image        *BuildLogStepDockerImage `json:"image,omitempty"`
	Duration     *time.Duration           `json:"duration,omitempty"`
	ExitCode     *int64                   `json:"exitCode,omitempty"`
	Status       *string                  `json:"status,omitempty"`
	AutoInjected *bool                    `json:"autoInjected,omitempty"`
}

// GetAggregatedStatus returns the status aggregated across all stages
func (buildLog *BuildLog) GetAggregatedStatus() string {
	return GetAggregatedStatus(buildLog.Steps)
}

// GetAggregatedStatus returns the status aggregated across all stages
func GetAggregatedStatus(steps []*BuildLogStep) string {

	// aggregate per stage in order to take retries into account
	statusPerStage := map[string]string{}
	for _, bls := range steps {
		// last status for a stage is leading
		statusPerStage[bls.Step] = bls.Status
	}

	// if any stage ended in failure, the aggregated status is failed as well
	aggregatedStatus := "SUCCEEDED"
	for _, status := range statusPerStage {
		if status == "FAILED" {
			aggregatedStatus = "FAILED"
		}
	}

	return aggregatedStatus
}

// HasSucceededStatus returns true if aggregated status is succeeded
func (buildLog *BuildLog) HasSucceededStatus() bool {
	return HasSucceededStatus(buildLog.Steps)
}

// HasSucceededStatus returns true if aggregated status is succeeded
func HasSucceededStatus(steps []*BuildLogStep) bool {
	status := GetAggregatedStatus(steps)

	return status == "SUCCEEDED"
}
