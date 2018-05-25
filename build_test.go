package contracts

import (
	"bytes"
	"testing"
	"time"

	"github.com/google/jsonapi"
	"github.com/stretchr/testify/assert"
)

func TestBuild(t *testing.T) {

	t.Run("JSONAPIMarshalPayloadSingleBuild", func(t *testing.T) {

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

		b := new(bytes.Buffer)

		// act
		err := jsonapi.MarshalPayload(b, &build)

		assert.Nil(t, err)
		assert.Equal(t, "{\"data\":{\"type\":\"builds\",\"id\":\"3\",\"attributes\":{\"build-status\":\"succeeded\",\"build-version\":\"1.0.0\",\"inserted-at\":1523952180,\"labels\":\"\",\"manifest\":\"\",\"repo-branch\":\"master\",\"repo-name\":\"estafette-ci-api\",\"repo-owner\":\"estafette\",\"repo-revision\":\"as23456\",\"repo-source\":\"github.com\",\"updated-at\":1523952900}}}\n", b.String())
	})

	t.Run("JSONAPIMarshalPayloadArrayOfBuilds", func(t *testing.T) {

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

		b := new(bytes.Buffer)

		// act
		err := jsonapi.MarshalPayload(b, builds)

		assert.Nil(t, err)
		assert.Equal(t, "{\"data\":[{\"type\":\"builds\",\"id\":\"3\",\"attributes\":{\"build-status\":\"succeeded\",\"build-version\":\"1.0.0\",\"inserted-at\":1523952180,\"labels\":\"\",\"manifest\":\"\",\"repo-branch\":\"master\",\"repo-name\":\"estafette-ci-api\",\"repo-owner\":\"estafette\",\"repo-revision\":\"as23456\",\"repo-source\":\"github.com\",\"updated-at\":1523952900}},{\"type\":\"builds\",\"id\":\"8\",\"attributes\":{\"build-status\":\"succeeded\",\"build-version\":\"1.0.0\",\"inserted-at\":1523952180,\"labels\":\"\",\"manifest\":\"\",\"repo-branch\":\"master\",\"repo-name\":\"estafette-ci-api\",\"repo-owner\":\"estafette\",\"repo-revision\":\"as23456\",\"repo-source\":\"github.com\",\"updated-at\":1523952900}}]}\n", b.String())
	})
}
