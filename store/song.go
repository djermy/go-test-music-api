package store

import (
	"context"
	"log"

	"example.com/music-api/model"
	"github.com/jackc/pgx/v5"
)

var Conn *pgx.Conn

func GetSongs() ([]model.Song, error) {
	var songs []model.Song
	// append to slice - songs = append(songs, song)

	rows, err := Conn.Query(
		context.Background(),
		`
		SELECT
			id,
			title,
			author,
			album,
			genre
		FROM
			song
		;`,
	)

	if err != nil {
		log.Println(err)
		return []model.Song{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var song model.Song
		err := rows.Scan(
			&song.ID,
			&song.Title,
			&song.Author,
			&song.Album,
			&song.Genre,
		)
		if err != nil {
			log.Println(err)
			return []model.Song{}, err
		}

		songs = append(songs, song)
	}
	return songs, err
}

func GetSong(id string) (model.Song, error) {
	var song model.Song
	rows, err := Conn.Query(
		context.Background(),
		`
		SELECT
			id,
			title,
			author,
			album,
			genre
		FROM
			song
		WHERE
			id = $1
		LIMIT 1
		;`,
		id,
	)

	if err != nil {
		log.Println(err)
		return model.Song{}, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&song.ID,
			&song.Title,
			&song.Author,
			&song.Album,
			&song.Genre,
		)
		if err != nil {
			log.Println(err)
			return model.Song{}, err
		}
	}

	return song, nil
}

func CreateSong(song *model.Song) error {
	err := Conn.QueryRow(
		context.Background(),
		`
		INSERT INTO song (
			title,
			author,
			album,
			genre
		) VALUES (
			$1,
			$2,
			$3,
			$4
		)
		RETURNING id
		;`,
		song.Title,
		song.Author,
		song.Album,
		song.Genre,
	).Scan(&song.ID)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func UpdateSong(id string, song *model.Song) error {
	rows, err := Conn.Query(
		context.Background(),
		`
		UPDATE song
		SET 
			title = $1,
			author = $2,
			album = $3,
			genre = $4
		WHERE
			id = $5;
			`,
		song.Title,
		song.Author,
		song.Album,
		song.Genre,
		id,
	)
	if err != nil {
		log.Println(err)
		return err
	}
	defer rows.Close()

	return nil
}

func DeleteSong(id string) error {
	rows, err := Conn.Query(
		context.Background(),
		`
		DELETE FROM song
		WHERE
			id = $1;
			`,
		id,
	)
	if err != nil {
		log.Println(err)
		return err
	}
	defer rows.Close()

	return nil
}
