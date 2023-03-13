package players

import (
	"encoding/json"
	"os"
)

type FileSystemPlayerStore struct {
	database *json.Encoder
	league   League
}

func (f *FileSystemPlayerStore) GetLeague() League {
	return f.league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {

	player := f.GetLeague().Find(name)

	if player != nil {
		return player.Wins
	}

	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	player := f.league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{name, 1})
	}

	f.database.Encode(f.league)
}

func NewFileSystemPlayerStore(db *os.File) *FileSystemPlayerStore {
	league, _ := NewLeague(db)

	return &FileSystemPlayerStore{
		database: json.NewEncoder(&tape{db}),
		league:   league,
	}
}
