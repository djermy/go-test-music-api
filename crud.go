package main

import (
	"errors"
	"fmt"
	"sync"

	"example.com/music-api/model"
)

var (
	songs          = make(map[string]model.Song)
	muSongs        sync.Mutex
	idCounterSongs int = 1
)

func GetSongs() map[string]model.Song {
	muSongs.Lock()
	defer muSongs.Unlock()

	return songs
}

func GetSong(id string) (model.Song, error) {
	muSongs.Lock()
	defer muSongs.Unlock()

	song, exists := songs[id]
	if !exists {
		return song, errors.New("Song not found")
	}

	return song, nil
}

func CreateSong(song model.Song) {
	muSongs.Lock()
	defer muSongs.Unlock()

	id := fmt.Sprintf(
		"%d",
		idCounterSongs,
	)

	idCounterSongs++

	songs[id] = song

}

func UpdateSong(id string, song model.Song) error {
	muSongs.Lock()
	defer muSongs.Unlock()

	_, exists := songs[id]
	if !exists {
		return errors.New("Song not found")
	}

	songs[id] = song

	return nil
}

func DeleteSong(id string) error {
	muSongs.Lock()
	defer muSongs.Unlock()

	_, exists := songs[id]
	if !exists {
		return errors.New("Song not found")
	}

	delete(songs, id)

	return nil
}
