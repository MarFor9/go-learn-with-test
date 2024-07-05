package poker

import (
	"encoding/json"
	"io"
)

type League []Player

func NewLeague(rdr io.Reader) ([]Player, error) {
	var league []Player
	if err := json.NewDecoder(rdr).Decode(&league); err != nil {
		return nil, err
	}
	return league, nil
}

func (l League) Find(name string) *Player {
	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}
	}
	return nil
}
