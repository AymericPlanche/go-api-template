package api_test

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"myapp/internal/app/handlers/api"
	"myapp/internal/infrastructure/integration_tests"
	"myapp/internal/things"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

// This is an example of unit test
// it tests the handler with an in memory repo
// so it does not hit the DB. This is enough for simple tests and more efficient
// Note that we can only test the handler directly this way. The request does not go through the entire middleware pipeline
func TestThingsHandler(t *testing.T) {
	assert := assert.New(t)

	mustCreateThing := func(repo things.Repository, id, name string) things.Thing {
		thing := things.Thing{
			ID:   things.ID(id),
			Name: name,
		}

		err := repo.Create(thing)
		if err != nil {
			panic(err)
		}
		return thing
	}

	t.Run("list things", func(t *testing.T) {
		buf := &bytes.Buffer{}
		logger := zerolog.New(buf)

		repo := things.NewInMemoryRepository(logger)

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

// This is an example of integration test
// it tests the handler against a real DB
// Also the request is processed by the main (application level) handler, so it goes through the entire middleware pipeline
func TestThingsIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	assert := assert.New(t)

	t.Run("GET thing/{id}", func(t *testing.T) {
		t.Run("It returns a 404 when id does not exist", func(t *testing.T) {
			integration_tests.ResetDatabase()

			app := integration_tests.InitAppContext()

			req := httptest.NewRequest("GET", "http://example.com/things/blah", nil)
			w := httptest.NewRecorder()
			app.GetHTTPHandler().ServeHTTP(w, req)

			resp := w.Result()
			assert.Equal(http.StatusNotFound, resp.StatusCode)

			assert.Contains(app.GetLogsAsString(), "Response: 404")
		})

	})

}
