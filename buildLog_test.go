package contracts

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"

	"github.com/google/jsonapi"
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
			Steps: []BuildLogStep{
				BuildLogStep{
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
		assert.Equal(t, "{\"id\":\"5\",\"repoSource\":\"github.com\",\"repoOwner\":\"estafette\",\"repoName\":\"estafette-ci-api\",\"repoBranch\":\"master\",\"repoRevision\":\"as23456\",\"steps\":[{\"step\":\"init\",\"image\":{\"name\":\"golang\",\"tag\":\"1.10.2-alpine3.7\",\"isPulled\":false,\"imageSize\":135000,\"pullDuration\":2000000000},\"duration\":91000000000,\"logLines\":[{\"timestamp\":\"2018-04-17T08:03:00Z\",\"streamType\":\"stdout\",\"text\":\"ok  \\tgithub.com/estafette/estafette-ci-contracts\\t0.017s\"}],\"exitCode\":0,\"status\":\"SUCCEEDED\"}],\"insertedAt\":\"2018-04-17T08:03:00Z\"}", string(bytes))
	})

	t.Run("JSONAPIMarshalPayloadSingleBuildLog", func(t *testing.T) {

		buildLog := BuildLog{
			ID:           "5",
			RepoSource:   "github.com",
			RepoOwner:    "estafette",
			RepoName:     "estafette-ci-api",
			RepoBranch:   "master",
			RepoRevision: "as23456",
			Steps: []BuildLogStep{
				BuildLogStep{
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

		b := new(bytes.Buffer)

		// act
		err := jsonapi.MarshalPayload(b, &buildLog)

		assert.Nil(t, err)
		assert.Equal(t, "{\"data\":{\"type\":\"build-logs\",\"id\":\"5\",\"attributes\":{\"inserted-at\":1523952180,\"repo-branch\":\"master\",\"repo-name\":\"estafette-ci-api\",\"repo-owner\":\"estafette\",\"repo-revision\":\"as23456\",\"repo-source\":\"github.com\",\"steps\":[{\"step\":\"init\",\"image\":{\"name\":\"golang\",\"tag\":\"1.10.2-alpine3.7\",\"isPulled\":false,\"imageSize\":135000,\"pullDuration\":2000000000},\"duration\":91000000000,\"logLines\":[{\"timestamp\":\"2018-04-17T08:03:00Z\",\"streamType\":\"stdout\",\"text\":\"ok  \\tgithub.com/estafette/estafette-ci-contracts\\t0.017s\"}],\"exitCode\":0,\"status\":\"SUCCEEDED\"}]}}}\n", b.String())
	})

}
