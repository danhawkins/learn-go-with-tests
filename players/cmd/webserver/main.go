package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/danhawkins/learn-go-with-tests/players/poker"
)

const dbFileName = "game.db.json"
const port = 5001

func main() {
	log.SetOutput(os.Stdout)

	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	store, err := poker.NewFileSystemPlayerStore(db)

	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}

	server := poker.NewPlayerServer(store)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), server); err != nil {
		log.Fatalf("could not listen on port %d %v", port, err)
	}
}
