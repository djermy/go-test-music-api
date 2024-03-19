package main

import (
	"context"
	"log"
	"net/http"

	"example.com/music-api/database"
	"example.com/music-api/handler"
	"github.com/jackc/pgx/v5"
)

var Conn *pgx.Conn

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	var err error
	Conn, err = database.InitDB()
	if err != nil {
		panic(err)
	}
	defer Conn.Close(context.Background())

	handler.InitHandlersMusic()
	http.Handle("/", handler.Router)
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
