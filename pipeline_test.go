package contracts

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPipeline(t *testing.T) {

	t.Run("JSONMarshalPayloadSinglePipeline", func(t *testing.T) {

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

		// act
		bytes, err := json.Marshal(&pipeline)

		assert.Nil(t, err)
		assert.Equal(t, "{\"id\":\"5\",\"repoSource\":\"github.com\",\"repoOwner\":\"estafette\",\"repoName\":\"estafette-ci-api\",\"repoBranch\":\"master\",\"repoRevision\":\"as23456\",\"buildVersion\":\"1.0.0\",\"buildStatus\":\"succeeded\",\"labels\":\"\",\"manifest\":\"\",\"insertedAt\":\"2018-04-17T08:03:00Z\",\"updatedAt\":\"2018-04-17T08:15:00Z\"}", string(bytes))
	})

	t.Run("JSONMarshalPayloadArrayOfPipelines", func(t *testing.T) {

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

		// act
		bytes, err := json.Marshal(&pipelines)

		assert.Nil(t, err)
		assert.Equal(t, "[{\"id\":\"5\",\"repoSource\":\"github.com\",\"repoOwner\":\"estafette\",\"repoName\":\"estafette-ci-api\",\"repoBranch\":\"master\",\"repoRevision\":\"as23456\",\"buildVersion\":\"1.0.0\",\"buildStatus\":\"succeeded\",\"labels\":\"\",\"manifest\":\"\",\"insertedAt\":\"2018-04-17T08:03:00Z\",\"updatedAt\":\"2018-04-17T08:15:00Z\"},{\"id\":\"6\",\"repoSource\":\"github.com\",\"repoOwner\":\"estafette\",\"repoName\":\"estafette-ci-api\",\"repoBranch\":\"master\",\"repoRevision\":\"as23456\",\"buildVersion\":\"1.0.0\",\"buildStatus\":\"succeeded\",\"labels\":\"\",\"manifest\":\"\",\"insertedAt\":\"2018-04-17T08:03:00Z\",\"updatedAt\":\"2018-04-17T08:15:00Z\"}]", string(bytes))
	})
}
