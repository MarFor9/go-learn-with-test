package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   []Player
}

func (s *StubPlayerStore) GetLeague() League {
	return s.league
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	result, ok := s.scores[name]
	if !ok {
		return 0
	}
	return result
}
func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		scores: map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
	}
	server := NewPlayerServer(&store)

	t.Run("returns Pepper's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertEqual(t, response.Body.String(), "20")
		assertEqual(t, response.Code, http.StatusOK)
	})
	t.Run("returns Floyd's socre", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Floyd", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertEqual(t, response.Body.String(), "10")
		assertEqual(t, response.Code, http.StatusOK)
	})
	t.Run("returns 404 on missing players", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/MissingPlayer", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertEqual(t, response.Code, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		scores: map[string]int{},
	}
	server := NewPlayerServer(&store)
	t.Run("it records wins when POST", func(t *testing.T) {
		player := "Pepper"
		request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", player), nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertEqual(t, response.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Errorf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
		}
		if store.winCalls[0] != player {
			t.Errorf("did not storage correct winner got %q want %q", store.winCalls[0], player)
		}
	})
}

func TestLeague(t *testing.T) {
	wantedLeague := []Player{
		{"Cleo", 32},
		{"Chris", 20},
		{"Tiest", 14},
	}
	store := StubPlayerStore{league: wantedLeague}
	server := NewPlayerServer(&store)
	t.Run("it returns 200 on /league", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/league", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := getLeagueFromResponse(t, response)
		assertEqual(t, response.Code, http.StatusOK)
		assertLeague(t, got, wantedLeague)
		assertContentType(t, response)
	})
}

func assertContentType(t *testing.T, response *httptest.ResponseRecorder) {
	if response.Result().Header.Get("content-type") != jsonContentType {
		t.Errorf("response did not have content-type of application/json, got %v", response.Result().Header)
	}
}

func assertLeague(t testing.TB, got []Player, wantedLeague []Player) {
	t.Helper()
	if !reflect.DeepEqual(got, wantedLeague) {
		t.Errorf("got %v want %v", got, wantedLeague)
	}
}

func getLeagueFromResponse(t testing.TB, response *httptest.ResponseRecorder) []Player {
	t.Helper()
	league, err := NewLeague(response.Body)
	if err != nil {
		t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", response.Body, err)
	}
	return league
}

func assertEqual[T comparable](t testing.TB, got T, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %+v but want %+v", got, want)
	}
}
