package main

import (
	"log"
	"net/http"

	"github.com/danhawkins/learn-go-with-tests/players/players"
)

func main() {
	server := &players.PlayerServer{&players.InMemoryPlayerStore{}}
	log.Fatal(http.ListenAndServe(":5000", server))
}
