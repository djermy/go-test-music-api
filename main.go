package main

import (
	"log"
	"net/http"

	"example.com/music-api/database"
	"github.com/gorilla/mux"
)

var router = mux.NewRouter()

func main() {
	initHandlersMusic()
	database.InitDB()
	http.Handle("/", router)
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
