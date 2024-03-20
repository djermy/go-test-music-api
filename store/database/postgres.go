package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func InitDB() (*pgx.Conn, error) {
	conn, err := dbConnect()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = migrate(conn)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return conn, nil
}

func dbConnect() (*pgx.Conn, error) {
	url := "postgres://postgres:password@localhost:5432/music"
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return conn, nil
}

func migrate(conn *pgx.Conn) error {
	for _, query := range migrationQuery {
		_, err := conn.Exec(context.Background(), query)
		if err != nil {
			return err
		}
	}
	return nil
}
