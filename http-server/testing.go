package poker

import (
	"net/http/httptest"
	"reflect"
	"testing"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   []Player
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetLeague() League {
	return s.league
}

func AssertContentType(t *testing.T, response *httptest.ResponseRecorder) {
	if response.Result().Header.Get("content-type") != jsonContentType {
		t.Errorf("response did not have content-type of application/json, got %v", response.Result().Header)
	}
}

func AssertLeague(t testing.TB, got []Player, wantedLeague []Player) {
	t.Helper()
	if !reflect.DeepEqual(got, wantedLeague) {
		t.Errorf("got %v want %v", got, wantedLeague)
	}
}

func GetLeagueFromResponse(t testing.TB, response *httptest.ResponseRecorder) []Player {
	t.Helper()
	league, err := NewLeague(response.Body)
	if err != nil {
		t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", response.Body, err)
	}
	return league
}

func AssertEqual[T comparable](t testing.TB, got T, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %+v but want %+v", got, want)
	}
}

func AssertPlayerWin(t testing.TB, store *StubPlayerStore, winner string) {
	t.Helper()

	if len(store.winCalls) != 1 {
		t.Fatalf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
	}

	if store.winCalls[0] != winner {
		t.Errorf("did not store correct winner got %q want %q", store.winCalls[0], winner)
	}
}
