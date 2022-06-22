package main

import (
	"log"
	"net/http"
	"os"
)

//type InMemoryPlayerStore struct {
//	store map[string]int
//}
//
//func (i *InMemoryPlayerStore) GetLeague() []Player {
//	var league []Player
//	for name, wins := range i.store {
//		league = append(league, Player{name, wins})
//	}
//	return league
//}
//
//func NewInMemoryPlayerStore() *InMemoryPlayerStore {
//	return &InMemoryPlayerStore{map[string]int{}}
//}
//
//func (i *InMemoryPlayerStore) RecordWin(name string) {
//	i.store[name]++
//}
//
//func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
//	return i.store[name]
//}

const dbFileName = "game.db.json"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}
	store, err := NewFileSystemPlayerStore(db)
	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}
	server := NewPlayerServer(store)
	if err := http.ListenAndServe(":5001", server); err != nil {
		log.Fatalf("could not listen on port 5001 %v", err)
	}
}
