package store

import "example.com/music-api/model"

type Store interface {
	GetSongs() ([]model.Song, error)
	GetSong(string) (model.Song, error)
	CreateSong(*model.Song) error
	UpdateSong(string, *model.Song) error
	DeleteSong(string) error
}
