package contracts

import (
	"bytes"
	"testing"
	"time"

	"github.com/google/jsonapi"
	"github.com/stretchr/testify/assert"
)

func TestPipeline(t *testing.T) {

	t.Run("JSONAPIMarshalPayloadSinglePipeline", func(t *testing.T) {

		pipeline := Pipeline{
			ID:           "5",
			RepoSource:   "github.com",
			RepoOwner:    "estafette",
			RepoName:     "estafette-ci-api",
			RepoBranch:   "master",
			RepoRevision: "as23456",
			BuildVersion: "1.0.0",
			BuildStatus:  "succeeded",
			Labels:       "",
			Manifest:     "",
			InsertedAt:   time.Date(2018, 4, 17, 8, 3, 0, 0, time.UTC),
			UpdatedAt:    time.Date(2018, 4, 17, 8, 15, 0, 0, time.UTC),
		}

		b := new(bytes.Buffer)

		// act
		err := jsonapi.MarshalPayload(b, &pipeline)

		assert.Nil(t, err)
		assert.Equal(t, "{\"data\":{\"type\":\"pipelines\",\"id\":\"5\",\"attributes\":{\"build-status\":\"succeeded\",\"build-version\":\"1.0.0\",\"inserted-at\":1523952180,\"labels\":\"\",\"manifest\":\"\",\"repo-branch\":\"master\",\"repo-name\":\"estafette-ci-api\",\"repo-owner\":\"estafette\",\"repo-revision\":\"as23456\",\"repo-source\":\"github.com\",\"updated-at\":1523952900}}}\n", b.String())
	})

	t.Run("JSONAPIMarshalPayloadArrayOfPipelines", func(t *testing.T) {

		pipelines := make([]*Pipeline, 0)

		pipelines = append(pipelines, &Pipeline{
			ID:           "5",
			RepoSource:   "github.com",
			RepoOwner:    "estafette",
			RepoName:     "estafette-ci-api",
			RepoBranch:   "master",
			RepoRevision: "as23456",
			BuildVersion: "1.0.0",
			BuildStatus:  "succeeded",
			Labels:       "",
			Manifest:     "",
			InsertedAt:   time.Date(2018, 4, 17, 8, 3, 0, 0, time.UTC),
			UpdatedAt:    time.Date(2018, 4, 17, 8, 15, 0, 0, time.UTC),
		})
		pipelines = append(pipelines, &Pipeline{
			ID:           "6",
			RepoSource:   "github.com",
			RepoOwner:    "estafette",
			RepoName:     "estafette-ci-api",
			RepoBranch:   "master",
			RepoRevision: "as23456",
			BuildVersion: "1.0.0",
			BuildStatus:  "succeeded",
			Labels:       "",
			Manifest:     "",
			InsertedAt:   time.Date(2018, 4, 17, 8, 3, 0, 0, time.UTC),
			UpdatedAt:    time.Date(2018, 4, 17, 8, 15, 0, 0, time.UTC),
		})

		b := new(bytes.Buffer)

		// act
		err := jsonapi.MarshalPayload(b, pipelines)

		assert.Nil(t, err)
		assert.Equal(t, "{\"data\":[{\"type\":\"pipelines\",\"id\":\"5\",\"attributes\":{\"build-status\":\"succeeded\",\"build-version\":\"1.0.0\",\"inserted-at\":1523952180,\"labels\":\"\",\"manifest\":\"\",\"repo-branch\":\"master\",\"repo-name\":\"estafette-ci-api\",\"repo-owner\":\"estafette\",\"repo-revision\":\"as23456\",\"repo-source\":\"github.com\",\"updated-at\":1523952900}},{\"type\":\"pipelines\",\"id\":\"6\",\"attributes\":{\"build-status\":\"succeeded\",\"build-version\":\"1.0.0\",\"inserted-at\":1523952180,\"labels\":\"\",\"manifest\":\"\",\"repo-branch\":\"master\",\"repo-name\":\"estafette-ci-api\",\"repo-owner\":\"estafette\",\"repo-revision\":\"as23456\",\"repo-source\":\"github.com\",\"updated-at\":1523952900}}]}\n", b.String())
	})
}
