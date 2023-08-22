package api_test

import (
	"bytes"
	"context"
	"io"
	"net/http/httptest"
	"testing"

	"myapp/internal/app/handlers/api"
	"myapp/internal/things"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestThingsHandler(t *testing.T) {
	assert := assert.New(t)

	mustCreateThing := func(repo things.Repository, id, name string) things.Thing {
		thing := things.Thing{
			ID:   things.ID(id),
			Name: name,
		}

		err := repo.Create(context.Background(), thing)
		if err != nil {
			panic(err)
		}
		return thing
	}

	t.Run("list things", func(t *testing.T) {
		buf := &bytes.Buffer{}
		logger := zerolog.New(buf)

		repo := things.NewRepository(logger)

		mustCreateThing(repo, "abc", "first thing")
		mustCreateThing(repo, "def", "second thing")

		req := httptest.NewRequest("GET", "http://example.com/things", nil)
		w := httptest.NewRecorder()
		handler := api.ThingsHandler{
			Logger:     logger,
			ThingsRepo: repo,
		}
		handler.Index().ServeHTTP(w, req)

		resp := w.Result()
		body, _ := io.ReadAll(resp.Body)
		logs, _ := io.ReadAll(buf)

		expected := `[
			{"id": "abc", "name": "first thing"},
			{"id": "def", "name": "second thing"}
		]`

		assert.JSONEq(expected, string(body))
		assert.NotContains(logs, "error")
	})
}
