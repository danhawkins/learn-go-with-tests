package players

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type League []Player

func NewLeague(rdr *os.File) ([]Player, error) {
	var league League
	err := json.NewDecoder(rdr).Decode(&league)

	if err != nil {
		err = fmt.Errorf("problem parsing league, %v", err)
	}

	log.Printf("Got league %v", league)

	return league, err
}

func (l League) Find(name string) *Player {
	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}
	}
	return nil
}
