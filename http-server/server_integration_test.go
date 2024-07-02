package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	type testConfig struct {
		name  string
		store PlayerStore
	}
	postgresPlayerStore := NewPostgresPlayerStore()
	t.Cleanup(postgresPlayerStore.Close)
	for _, tc := range []testConfig{
		{
			name:  "InMemoryPlayerStore",
			store: NewInMemoryPlayerStore(),
		},
		{
			name:  "PostgresPlayerStore",
			store: postgresPlayerStore,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			server := PlayerServer{tc.store}
			player := "Peper"

			server.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", player), nil))
			server.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", player), nil))
			server.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", player), nil))

			response := httptest.NewRecorder()
			request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", player), nil)
			server.ServeHTTP(response, request)

			assertEqual(t, response.Code, http.StatusOK)
			assertEqual(t, response.Body.String(), "3")
		})
	}
}
