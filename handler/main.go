package handler

import "github.com/gorilla/mux"

var Router = mux.NewRouter()

func InitHandlersMusic() {
	Router.HandleFunc("/song", getSongs).Methods("GET")
	Router.HandleFunc("/song", createSong).Methods("POST")
	Router.HandleFunc("/song/{songid}", getSongByID).Methods("GET")
	Router.HandleFunc("/song/{songid}", updateSongByID).Methods("PUT")
	Router.HandleFunc("/song/{songid}", deleteSongByID).Methods("DELETE")
}
