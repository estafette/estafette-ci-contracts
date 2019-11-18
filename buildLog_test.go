package contracts

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBuildLog(t *testing.T) {
	t.Run("JSONMarshalSingleBuildLog", func(t *testing.T) {
		buildLog := BuildLog{
			ID:           "5",
			RepoSource:   "github.com",
			RepoOwner:    "estafette",
			RepoName:     "estafette-ci-api",
			RepoBranch:   "master",
			RepoRevision: "as23456",
			Steps: []*BuildLogStep{
				&BuildLogStep{
					Step: "init",
					Image: &BuildLogStepDockerImage{
						Name:         "golang",
						Tag:          "1.10.2-alpine3.7",
						IsPulled:     false,
						ImageSize:    135000,
						PullDuration: 2 * time.Second,
						Error:        "",
					},
					Duration: 91 * time.Second,
					LogLines: []BuildLogLine{
						BuildLogLine{
							Timestamp:  time.Date(2018, 4, 17, 8, 3, 0, 0, time.UTC),
							StreamType: "stdout",
							Text: "ok  	github.com/estafette/estafette-ci-contracts	0.017s",
						},
					},
					ExitCode: 0,
					Status:   "SUCCEEDED",
				},
			},
			InsertedAt: time.Date(2018, 4, 17, 8, 3, 0, 0, time.UTC),
		}
		// act
		bytes, err := json.Marshal(&buildLog)
		assert.Nil(t, err)
		assert.Equal(t, "{\"id\":\"5\",\"repoSource\":\"github.com\",\"repoOwner\":\"estafette\",\"repoName\":\"estafette-ci-api\",\"repoBranch\":\"master\",\"repoRevision\":\"as23456\",\"buildID\":\"\",\"steps\":[{\"step\":\"init\",\"image\":{\"name\":\"golang\",\"tag\":\"1.10.2-alpine3.7\",\"isPulled\":false,\"imageSize\":135000,\"pullDuration\":2000000000},\"duration\":91000000000,\"logLines\":[{\"timestamp\":\"2018-04-17T08:03:00Z\",\"streamType\":\"stdout\",\"text\":\"ok  \\tgithub.com/estafette/estafette-ci-contracts\\t0.017s\"}],\"exitCode\":0,\"status\":\"SUCCEEDED\"}],\"insertedAt\":\"2018-04-17T08:03:00Z\"}", string(bytes))
	})
}

func TestGetAggregatedStatus(t *testing.T) {
	t.Run("ReturnsSucceededIfNoSteps", func(t *testing.T) {

		steps := []*BuildLogStep{}

		// act
		status := GetAggregatedStatus(steps)

		assert.Equal(t, "SUCCEEDED", status)
	})

	t.Run("ReturnsFailedIfAllStepsFailed", func(t *testing.T) {

		steps := []*BuildLogStep{
			&BuildLogStep{
				Step:   "stage-a",
				Status: "FAILED",
			},
			&BuildLogStep{
				Step:   "stage-b",
				Status: "FAILED",
			},
		}

		// act
		status := GetAggregatedStatus(steps)

		assert.Equal(t, "FAILED", status)
	})

	t.Run("ReturnsFailedIfAnyStepsFailed", func(t *testing.T) {

		steps := []*BuildLogStep{
			&BuildLogStep{
				Step:   "stage-a",
				Status: "SUCCEEDED",
			},
			&BuildLogStep{
				Step:   "stage-b",
				Status: "FAILED",
			},
			&BuildLogStep{
				Step:   "stage-c",
				Status: "SUCCEEDED",
			},
		}

		// act
		status := GetAggregatedStatus(steps)

		assert.Equal(t, "FAILED", status)
	})

	t.Run("ReturnsSucceededIfAStepFailedButSucceededInRetry", func(t *testing.T) {

		steps := []*BuildLogStep{
			&BuildLogStep{
				Step:   "stage-a",
				Status: "SUCCEEDED",
			},
			&BuildLogStep{
				Step:   "stage-b",
				Status: "FAILED",
			},
			&BuildLogStep{
				Step:     "stage-b",
				RunIndex: 1,
				Status:   "SUCCEEDED",
			},
			&BuildLogStep{
				Step:   "stage-c",
				Status: "SUCCEEDED",
			},
		}

		// act
		status := GetAggregatedStatus(steps)

		assert.Equal(t, "SUCCEEDED", status)
	})

	t.Run("ReturnsFailedIfAStepFailedButSucceededInRetryButAnotherStepFailed", func(t *testing.T) {

		steps := []*BuildLogStep{
			&BuildLogStep{
				Step:   "stage-a",
				Status: "SUCCEEDED",
			},
			&BuildLogStep{
				Step:   "stage-b",
				Status: "FAILED",
			},
			&BuildLogStep{
				Step:     "stage-b",
				RunIndex: 1,
				Status:   "SUCCEEDED",
			},
			&BuildLogStep{
				Step:   "stage-c",
				Status: "FAILED",
			},
		}

		// act
		status := GetAggregatedStatus(steps)

		assert.Equal(t, "FAILED", status)
	})
}
