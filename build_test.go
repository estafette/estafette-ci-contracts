package contracts

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBuild(t *testing.T) {

	t.Run("JSONMarshalPayloadSingleBuild", func(t *testing.T) {

		build := Build{
			ID:           "3",
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
		bytes, err := json.Marshal(&build)

		assert.Nil(t, err)
		assert.Equal(t, "{\"id\":\"3\",\"repoSource\":\"github.com\",\"repoOwner\":\"estafette\",\"repoName\":\"estafette-ci-api\",\"repoBranch\":\"master\",\"repoRevision\":\"as23456\",\"buildVersion\":\"1.0.0\",\"buildStatus\":\"succeeded\",\"labels\":\"\",\"manifest\":\"\",\"insertedAt\":\"2018-04-17T08:03:00Z\",\"updatedAt\":\"2018-04-17T08:15:00Z\"}", string(bytes))
	})

	t.Run("JSONMarshalPayloadArrayOfBuilds", func(t *testing.T) {

		builds := make([]*Build, 0)

		builds = append(builds, &Build{
			ID:           "3",
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
		builds = append(builds, &Build{
			ID:           "8",
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
		bytes, err := json.Marshal(&builds)

		assert.Nil(t, err)
		assert.Equal(t, "[{\"id\":\"3\",\"repoSource\":\"github.com\",\"repoOwner\":\"estafette\",\"repoName\":\"estafette-ci-api\",\"repoBranch\":\"master\",\"repoRevision\":\"as23456\",\"buildVersion\":\"1.0.0\",\"buildStatus\":\"succeeded\",\"labels\":\"\",\"manifest\":\"\",\"insertedAt\":\"2018-04-17T08:03:00Z\",\"updatedAt\":\"2018-04-17T08:15:00Z\"},{\"id\":\"8\",\"repoSource\":\"github.com\",\"repoOwner\":\"estafette\",\"repoName\":\"estafette-ci-api\",\"repoBranch\":\"master\",\"repoRevision\":\"as23456\",\"buildVersion\":\"1.0.0\",\"buildStatus\":\"succeeded\",\"labels\":\"\",\"manifest\":\"\",\"insertedAt\":\"2018-04-17T08:03:00Z\",\"updatedAt\":\"2018-04-17T08:15:00Z\"}]", string(bytes))
	})
}
