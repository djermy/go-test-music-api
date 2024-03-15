package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"example.com/music-api/model"
	"github.com/gorilla/mux"
)

func initHandlersMusic() {
	router.HandleFunc("/song", getSongs).Methods("GET")
	router.HandleFunc("/song", createSong).Methods("POST")
	router.HandleFunc("/song/{songid}", getSongByID).Methods("GET")
	router.HandleFunc("/song/{songid}", updateSongByID).Methods("PUT")
	router.HandleFunc("/song/{songid}", deleteSongByID).Methods("DELETE")
}

func getSongs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := writeJSON(w, GetSongs())
	if err != nil {
		fmt.Println(err)
		return
	}
}

func getSongByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, ok := mux.Vars(r)["id"]
	if !ok {
		fmt.Println("id is missing in parameters")
		return
	}

	targetSong, err := GetSong(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = writeJSON(w, targetSong)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func createSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newSong model.Song

	err := readJSON(w, r, &newSong)
	if err != nil {
		fmt.Println(err)
		return
	}

	CreateSong(newSong)
}

func updateSongByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, ok := mux.Vars(r)["id"]
	if !ok {
		fmt.Println("id is missing in parameters")
		return
	}

	var updatedSong model.Song

	err := readJSON(w, r, &updatedSong)
	if err != nil {
		fmt.Println(err)
		return
	}

	UpdateSong(id, updatedSong)
}

func deleteSongByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, ok := mux.Vars(r)["id"]
	if !ok {
		fmt.Println("id is missing in parameters")
		return
	}

	err := DeleteSong(id)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func readJSON(w http.ResponseWriter, r *http.Request, v interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	return nil
}

func writeJSON(w http.ResponseWriter, v interface{}) error {
	if err := json.NewEncoder(w).Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	return nil
}
