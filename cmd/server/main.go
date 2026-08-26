package main

import (
	"log"
	"net/http"
	"retirement53/api"
	"retirement53/config"
	"retirement53/retirement"
	"retirement53/storage"
)

func main() {
	c := config.Load()
	s, e := storage.Open(c.DBPath)
	if e != nil {
		log.Fatal(e)
	}
	defer s.Close()
	log.Fatal(http.ListenAndServe(c.Addr, api.New(retirement.NewService(s)).Handler()))
}
