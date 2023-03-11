package main

import (
	"log"
	"net/http"

	"github.com/danhawkins/learn-go-with-tests/players/players"
)

func main() {
	server := players.NewPlayerServer(players.NewInMemoryPlayerStore())
	log.Fatal(http.ListenAndServe(":5050", server))
}
