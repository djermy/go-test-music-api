package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"example.com/music-api/model"
	"example.com/music-api/store"
	"github.com/gorilla/mux"
)

func getSongs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	songs, _ := store.GetSongs()
	err := writeJSON(w, songs)
	if err != nil {
		log.Println(err)
		return
	}
}

func getSongByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, ok := mux.Vars(r)["songid"]
	if !ok {
		err := errors.New("id is missing in parameters")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	targetSong, err := store.GetSong(id)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = writeJSON(w, targetSong)
	if err != nil {
		log.Println(err)
		return
	}
}

func createSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newSong model.Song

	err := readJSON(w, r, &newSong)
	if err != nil {
		log.Println(err)
		return
	}

	err = store.CreateSong(&newSong)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = writeJSON(w, &newSong)
	if err != nil {
		log.Println(err)
		return
	}
}

func updateSongByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, ok := mux.Vars(r)["songid"]
	if !ok {
		err := errors.New("id is missing in parameters")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var updatedSong model.Song

	err := readJSON(w, r, &updatedSong)
	if err != nil {
		log.Println(err)
		return
	}

	err = store.UpdateSong(id, &updatedSong)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func deleteSongByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, ok := mux.Vars(r)["songid"]
	if !ok {
		err := errors.New("id is missing in parameters")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := store.DeleteSong(id)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
