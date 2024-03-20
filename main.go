package main

import (
	"context"
	"log"
	"net/http"

	"example.com/music-api/handler"
	"example.com/music-api/store"
	"example.com/music-api/store/database"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	var err error
	Conn, err = database.InitDB()
	if err != nil {
		panic(err)
	}
	defer store.Conn.Close(context.Background())

	handler.InitHandlersMusic()
	http.Handle("/", handler.Router)
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
