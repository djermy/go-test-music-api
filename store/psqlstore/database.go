package psqlstore

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func newConn() (*pgx.Conn, error) {
	conn, err := connect()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = migrateUp(conn)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return conn, nil
}

func connect() (*pgx.Conn, error) {
	url := "postgres://postgres:password@localhost:5432/music"
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return conn, nil
}

func migrateUp(conn *pgx.Conn) error {
	for _, m := range migrations {
		_, err := conn.Exec(context.Background(), m.Up)
		if err != nil {
			return err
		}
	}
	return nil
}
