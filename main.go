package main

import (
	"log"

	"example.com/music-api/handler"
	"example.com/music-api/store/psqlstore"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	s, err := psqlstore.New()
	if err != nil {
		panic(err)
	}

	h := handler.New(s)
	log.Fatal(h.Run())
}
