package players

import (
	"encoding/json"
	"io"
	"log"
)

type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
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

	log.Printf("Got league %v\n", f.league)

	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{name, 1})
	}

	f.database.Seek(0, 0)
	json.NewEncoder(f.database).Encode(f.league)
}

func NewFileSystemPlayerStore(db io.ReadWriteSeeker) *FileSystemPlayerStore {
	db.Seek(0, 0)
	league, _ := NewLeague(db)
	return &FileSystemPlayerStore{
		database: db,
		league:   league,
	}
}
